package start_auction_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs/prune"
	"github.com/cloudfoundry-incubator/runtime-schema/bbs/shared"
	"github.com/cloudfoundry-incubator/runtime-schema/metric"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/cloudfoundry/storeadapter"
	"github.com/pivotal-golang/lager"
)

const (
	convergeLrpStartCounter     = metric.Counter("converge-lrp-start-auction")
	pruneInvalidLrpStartCounter = metric.Counter("prune-invalid-lrp-start-auction")
	pruneClaimedLrpStartCounter = metric.Counter("prune-claimed-lrp-start-auction")
	pruneStartCounter           = metric.Counter("prune-start-auction-failed")
	casLrpStartCounter          = metric.Counter("compare-and-swap-lrp-start-auction")
)

type compareAndSwappableLRPStartAuction struct {
	OldIndex           uint64
	NewLRPStartAuction models.LRPStartAuction
}

func (bbs *StartAuctionBBS) ConvergeLRPStartAuctions(kickPendingDuration time.Duration, expireClaimedDuration time.Duration) {
	convergeLrpStartCounter.Increment()
	auctionsToCAS := []compareAndSwappableLRPStartAuction{}

	err := prune.Prune(bbs.store, shared.LRPStartAuctionSchemaRoot, func(auctionNode storeadapter.StoreNode) (shouldKeep bool) {
		auction, err := models.NewLRPStartAuctionFromJSON(auctionNode.Value)
		if err != nil {
			bbs.logger.Info("detected-invalid-start-auction-json", lager.Data{
				"error":   err.Error(),
				"payload": auctionNode.Value,
			})
			pruneInvalidLrpStartCounter.Increment()
			return false
		}

		updatedAt := time.Unix(0, auction.UpdatedAt)

		switch auction.State {
		case models.LRPStartAuctionStatePending:
			if bbs.timeProvider.Time().Sub(updatedAt) > kickPendingDuration {
				bbs.logger.Info("detected-pending-auction", lager.Data{
					"auction":       auction,
					"kick-duration": kickPendingDuration,
				})

				auctionsToCAS = append(auctionsToCAS, compareAndSwappableLRPStartAuction{
					OldIndex:           auctionNode.Index,
					NewLRPStartAuction: auction,
				})
			}

		case models.LRPStartAuctionStateClaimed:
			if bbs.timeProvider.Time().Sub(updatedAt) > expireClaimedDuration {
				bbs.logger.Info("detected-expired-claim", lager.Data{
					"auction":             auction,
					"expiration-duration": expireClaimedDuration,
				})
				pruneClaimedLrpStartCounter.Increment()
				return false
			}
		}

		return true
	})

	if err != nil {
		pruneStartCounter.Increment()
		bbs.logger.Error("failed-to-prune-start-auction", err)
		return
	}

	casLrpStartCounter.Add(uint64(len(auctionsToCAS)))
	bbs.batchCompareAndSwapLRPStartAuctions(auctionsToCAS)
}

func (bbs *StartAuctionBBS) batchCompareAndSwapLRPStartAuctions(auctionsToCAS []compareAndSwappableLRPStartAuction) {
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(len(auctionsToCAS))
	for _, auctionToCAS := range auctionsToCAS {
		auction := auctionToCAS.NewLRPStartAuction
		newStoreNode := storeadapter.StoreNode{
			Key:   shared.LRPStartAuctionSchemaPath(auction),
			Value: auction.ToJSON(),
		}

		go func(auctionToCAS compareAndSwappableLRPStartAuction, newStoreNode storeadapter.StoreNode) {
			err := bbs.store.CompareAndSwapByIndex(auctionToCAS.OldIndex, newStoreNode)
			if err != nil {
				bbs.logger.Error("failed-to-compare-and-swap", err)
			}

			waitGroup.Done()
		}(auctionToCAS, newStoreNode)
	}

	waitGroup.Wait()
}
