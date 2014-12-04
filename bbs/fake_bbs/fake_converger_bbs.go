// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/tedsuo/ifrit"
)

type FakeConvergerBBS struct {
	ConvergeLRPsStub        func()
	convergeLRPsMutex       sync.RWMutex
	convergeLRPsArgsForCall []struct{}
	ActualLRPsByProcessGuidStub        func(string) ([]models.ActualLRP, error)
	actualLRPsByProcessGuidMutex       sync.RWMutex
	actualLRPsByProcessGuidArgsForCall []struct {
		arg1 string
	}
	actualLRPsByProcessGuidReturns struct {
		result1 []models.ActualLRP
		result2 error
	}
	RequestStopLRPInstanceStub        func(lrp models.ActualLRP) error
	requestStopLRPInstanceMutex       sync.RWMutex
	requestStopLRPInstanceArgsForCall []struct {
		lrp models.ActualLRP
	}
	requestStopLRPInstanceReturns struct {
		result1 error
	}
	WatchForDesiredLRPChangesStub        func() (<-chan models.DesiredLRPChange, chan<- bool, <-chan error)
	watchForDesiredLRPChangesMutex       sync.RWMutex
	watchForDesiredLRPChangesArgsForCall []struct{}
	watchForDesiredLRPChangesReturns struct {
		result1 <-chan models.DesiredLRPChange
		result2 chan<- bool
		result3 <-chan error
	}
	CreateActualLRPStub        func(models.ActualLRP) (*models.ActualLRP, error)
	createActualLRPMutex       sync.RWMutex
	createActualLRPArgsForCall []struct {
		arg1 models.ActualLRP
	}
	createActualLRPReturns struct {
		result1 *models.ActualLRP
		result2 error
	}
	ConvergeLRPStartAuctionsStub        func(kickPendingDuration time.Duration, expireClaimedDuration time.Duration)
	convergeLRPStartAuctionsMutex       sync.RWMutex
	convergeLRPStartAuctionsArgsForCall []struct {
		kickPendingDuration   time.Duration
		expireClaimedDuration time.Duration
	}
	RequestLRPStartAuctionStub        func(models.LRPStartAuction) error
	requestLRPStartAuctionMutex       sync.RWMutex
	requestLRPStartAuctionArgsForCall []struct {
		arg1 models.LRPStartAuction
	}
	requestLRPStartAuctionReturns struct {
		result1 error
	}
	ConvergeLRPStopAuctionsStub        func(kickPendingDuration time.Duration, expireClaimedDuration time.Duration)
	convergeLRPStopAuctionsMutex       sync.RWMutex
	convergeLRPStopAuctionsArgsForCall []struct {
		kickPendingDuration   time.Duration
		expireClaimedDuration time.Duration
	}
	RequestLRPStopAuctionStub        func(models.LRPStopAuction) error
	requestLRPStopAuctionMutex       sync.RWMutex
	requestLRPStopAuctionArgsForCall []struct {
		arg1 models.LRPStopAuction
	}
	requestLRPStopAuctionReturns struct {
		result1 error
	}
	ConvergeTaskStub        func(timeToClaim, convergenceInterval, timeToResolve time.Duration)
	convergeTaskMutex       sync.RWMutex
	convergeTaskArgsForCall []struct {
		timeToClaim         time.Duration
		convergenceInterval time.Duration
		timeToResolve       time.Duration
	}
	NewConvergeLockStub        func(convergerID string, interval time.Duration) ifrit.Runner
	newConvergeLockMutex       sync.RWMutex
	newConvergeLockArgsForCall []struct {
		convergerID string
		interval    time.Duration
	}
	newConvergeLockReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeConvergerBBS) ConvergeLRPs() {
	fake.convergeLRPsMutex.Lock()
	fake.convergeLRPsArgsForCall = append(fake.convergeLRPsArgsForCall, struct{}{})
	fake.convergeLRPsMutex.Unlock()
	if fake.ConvergeLRPsStub != nil {
		fake.ConvergeLRPsStub()
	}
}

func (fake *FakeConvergerBBS) ConvergeLRPsCallCount() int {
	fake.convergeLRPsMutex.RLock()
	defer fake.convergeLRPsMutex.RUnlock()
	return len(fake.convergeLRPsArgsForCall)
}

func (fake *FakeConvergerBBS) ActualLRPsByProcessGuid(arg1 string) ([]models.ActualLRP, error) {
	fake.actualLRPsByProcessGuidMutex.Lock()
	fake.actualLRPsByProcessGuidArgsForCall = append(fake.actualLRPsByProcessGuidArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.actualLRPsByProcessGuidMutex.Unlock()
	if fake.ActualLRPsByProcessGuidStub != nil {
		return fake.ActualLRPsByProcessGuidStub(arg1)
	} else {
		return fake.actualLRPsByProcessGuidReturns.result1, fake.actualLRPsByProcessGuidReturns.result2
	}
}

func (fake *FakeConvergerBBS) ActualLRPsByProcessGuidCallCount() int {
	fake.actualLRPsByProcessGuidMutex.RLock()
	defer fake.actualLRPsByProcessGuidMutex.RUnlock()
	return len(fake.actualLRPsByProcessGuidArgsForCall)
}

func (fake *FakeConvergerBBS) ActualLRPsByProcessGuidArgsForCall(i int) string {
	fake.actualLRPsByProcessGuidMutex.RLock()
	defer fake.actualLRPsByProcessGuidMutex.RUnlock()
	return fake.actualLRPsByProcessGuidArgsForCall[i].arg1
}

func (fake *FakeConvergerBBS) ActualLRPsByProcessGuidReturns(result1 []models.ActualLRP, result2 error) {
	fake.ActualLRPsByProcessGuidStub = nil
	fake.actualLRPsByProcessGuidReturns = struct {
		result1 []models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeConvergerBBS) RequestStopLRPInstance(lrp models.ActualLRP) error {
	fake.requestStopLRPInstanceMutex.Lock()
	fake.requestStopLRPInstanceArgsForCall = append(fake.requestStopLRPInstanceArgsForCall, struct {
		lrp models.ActualLRP
	}{lrp})
	fake.requestStopLRPInstanceMutex.Unlock()
	if fake.RequestStopLRPInstanceStub != nil {
		return fake.RequestStopLRPInstanceStub(lrp)
	} else {
		return fake.requestStopLRPInstanceReturns.result1
	}
}

func (fake *FakeConvergerBBS) RequestStopLRPInstanceCallCount() int {
	fake.requestStopLRPInstanceMutex.RLock()
	defer fake.requestStopLRPInstanceMutex.RUnlock()
	return len(fake.requestStopLRPInstanceArgsForCall)
}

func (fake *FakeConvergerBBS) RequestStopLRPInstanceArgsForCall(i int) models.ActualLRP {
	fake.requestStopLRPInstanceMutex.RLock()
	defer fake.requestStopLRPInstanceMutex.RUnlock()
	return fake.requestStopLRPInstanceArgsForCall[i].lrp
}

func (fake *FakeConvergerBBS) RequestStopLRPInstanceReturns(result1 error) {
	fake.RequestStopLRPInstanceStub = nil
	fake.requestStopLRPInstanceReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConvergerBBS) WatchForDesiredLRPChanges() (<-chan models.DesiredLRPChange, chan<- bool, <-chan error) {
	fake.watchForDesiredLRPChangesMutex.Lock()
	fake.watchForDesiredLRPChangesArgsForCall = append(fake.watchForDesiredLRPChangesArgsForCall, struct{}{})
	fake.watchForDesiredLRPChangesMutex.Unlock()
	if fake.WatchForDesiredLRPChangesStub != nil {
		return fake.WatchForDesiredLRPChangesStub()
	} else {
		return fake.watchForDesiredLRPChangesReturns.result1, fake.watchForDesiredLRPChangesReturns.result2, fake.watchForDesiredLRPChangesReturns.result3
	}
}

func (fake *FakeConvergerBBS) WatchForDesiredLRPChangesCallCount() int {
	fake.watchForDesiredLRPChangesMutex.RLock()
	defer fake.watchForDesiredLRPChangesMutex.RUnlock()
	return len(fake.watchForDesiredLRPChangesArgsForCall)
}

func (fake *FakeConvergerBBS) WatchForDesiredLRPChangesReturns(result1 <-chan models.DesiredLRPChange, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForDesiredLRPChangesStub = nil
	fake.watchForDesiredLRPChangesReturns = struct {
		result1 <-chan models.DesiredLRPChange
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeConvergerBBS) CreateActualLRP(arg1 models.ActualLRP) (*models.ActualLRP, error) {
	fake.createActualLRPMutex.Lock()
	fake.createActualLRPArgsForCall = append(fake.createActualLRPArgsForCall, struct {
		arg1 models.ActualLRP
	}{arg1})
	fake.createActualLRPMutex.Unlock()
	if fake.CreateActualLRPStub != nil {
		return fake.CreateActualLRPStub(arg1)
	} else {
		return fake.createActualLRPReturns.result1, fake.createActualLRPReturns.result2
	}
}

func (fake *FakeConvergerBBS) CreateActualLRPCallCount() int {
	fake.createActualLRPMutex.RLock()
	defer fake.createActualLRPMutex.RUnlock()
	return len(fake.createActualLRPArgsForCall)
}

func (fake *FakeConvergerBBS) CreateActualLRPArgsForCall(i int) models.ActualLRP {
	fake.createActualLRPMutex.RLock()
	defer fake.createActualLRPMutex.RUnlock()
	return fake.createActualLRPArgsForCall[i].arg1
}

func (fake *FakeConvergerBBS) CreateActualLRPReturns(result1 *models.ActualLRP, result2 error) {
	fake.CreateActualLRPStub = nil
	fake.createActualLRPReturns = struct {
		result1 *models.ActualLRP
		result2 error
	}{result1, result2}
}

func (fake *FakeConvergerBBS) ConvergeLRPStartAuctions(kickPendingDuration time.Duration, expireClaimedDuration time.Duration) {
	fake.convergeLRPStartAuctionsMutex.Lock()
	fake.convergeLRPStartAuctionsArgsForCall = append(fake.convergeLRPStartAuctionsArgsForCall, struct {
		kickPendingDuration   time.Duration
		expireClaimedDuration time.Duration
	}{kickPendingDuration, expireClaimedDuration})
	fake.convergeLRPStartAuctionsMutex.Unlock()
	if fake.ConvergeLRPStartAuctionsStub != nil {
		fake.ConvergeLRPStartAuctionsStub(kickPendingDuration, expireClaimedDuration)
	}
}

func (fake *FakeConvergerBBS) ConvergeLRPStartAuctionsCallCount() int {
	fake.convergeLRPStartAuctionsMutex.RLock()
	defer fake.convergeLRPStartAuctionsMutex.RUnlock()
	return len(fake.convergeLRPStartAuctionsArgsForCall)
}

func (fake *FakeConvergerBBS) ConvergeLRPStartAuctionsArgsForCall(i int) (time.Duration, time.Duration) {
	fake.convergeLRPStartAuctionsMutex.RLock()
	defer fake.convergeLRPStartAuctionsMutex.RUnlock()
	return fake.convergeLRPStartAuctionsArgsForCall[i].kickPendingDuration, fake.convergeLRPStartAuctionsArgsForCall[i].expireClaimedDuration
}

func (fake *FakeConvergerBBS) RequestLRPStartAuction(arg1 models.LRPStartAuction) error {
	fake.requestLRPStartAuctionMutex.Lock()
	fake.requestLRPStartAuctionArgsForCall = append(fake.requestLRPStartAuctionArgsForCall, struct {
		arg1 models.LRPStartAuction
	}{arg1})
	fake.requestLRPStartAuctionMutex.Unlock()
	if fake.RequestLRPStartAuctionStub != nil {
		return fake.RequestLRPStartAuctionStub(arg1)
	} else {
		return fake.requestLRPStartAuctionReturns.result1
	}
}

func (fake *FakeConvergerBBS) RequestLRPStartAuctionCallCount() int {
	fake.requestLRPStartAuctionMutex.RLock()
	defer fake.requestLRPStartAuctionMutex.RUnlock()
	return len(fake.requestLRPStartAuctionArgsForCall)
}

func (fake *FakeConvergerBBS) RequestLRPStartAuctionArgsForCall(i int) models.LRPStartAuction {
	fake.requestLRPStartAuctionMutex.RLock()
	defer fake.requestLRPStartAuctionMutex.RUnlock()
	return fake.requestLRPStartAuctionArgsForCall[i].arg1
}

func (fake *FakeConvergerBBS) RequestLRPStartAuctionReturns(result1 error) {
	fake.RequestLRPStartAuctionStub = nil
	fake.requestLRPStartAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConvergerBBS) ConvergeLRPStopAuctions(kickPendingDuration time.Duration, expireClaimedDuration time.Duration) {
	fake.convergeLRPStopAuctionsMutex.Lock()
	fake.convergeLRPStopAuctionsArgsForCall = append(fake.convergeLRPStopAuctionsArgsForCall, struct {
		kickPendingDuration   time.Duration
		expireClaimedDuration time.Duration
	}{kickPendingDuration, expireClaimedDuration})
	fake.convergeLRPStopAuctionsMutex.Unlock()
	if fake.ConvergeLRPStopAuctionsStub != nil {
		fake.ConvergeLRPStopAuctionsStub(kickPendingDuration, expireClaimedDuration)
	}
}

func (fake *FakeConvergerBBS) ConvergeLRPStopAuctionsCallCount() int {
	fake.convergeLRPStopAuctionsMutex.RLock()
	defer fake.convergeLRPStopAuctionsMutex.RUnlock()
	return len(fake.convergeLRPStopAuctionsArgsForCall)
}

func (fake *FakeConvergerBBS) ConvergeLRPStopAuctionsArgsForCall(i int) (time.Duration, time.Duration) {
	fake.convergeLRPStopAuctionsMutex.RLock()
	defer fake.convergeLRPStopAuctionsMutex.RUnlock()
	return fake.convergeLRPStopAuctionsArgsForCall[i].kickPendingDuration, fake.convergeLRPStopAuctionsArgsForCall[i].expireClaimedDuration
}

func (fake *FakeConvergerBBS) RequestLRPStopAuction(arg1 models.LRPStopAuction) error {
	fake.requestLRPStopAuctionMutex.Lock()
	fake.requestLRPStopAuctionArgsForCall = append(fake.requestLRPStopAuctionArgsForCall, struct {
		arg1 models.LRPStopAuction
	}{arg1})
	fake.requestLRPStopAuctionMutex.Unlock()
	if fake.RequestLRPStopAuctionStub != nil {
		return fake.RequestLRPStopAuctionStub(arg1)
	} else {
		return fake.requestLRPStopAuctionReturns.result1
	}
}

func (fake *FakeConvergerBBS) RequestLRPStopAuctionCallCount() int {
	fake.requestLRPStopAuctionMutex.RLock()
	defer fake.requestLRPStopAuctionMutex.RUnlock()
	return len(fake.requestLRPStopAuctionArgsForCall)
}

func (fake *FakeConvergerBBS) RequestLRPStopAuctionArgsForCall(i int) models.LRPStopAuction {
	fake.requestLRPStopAuctionMutex.RLock()
	defer fake.requestLRPStopAuctionMutex.RUnlock()
	return fake.requestLRPStopAuctionArgsForCall[i].arg1
}

func (fake *FakeConvergerBBS) RequestLRPStopAuctionReturns(result1 error) {
	fake.RequestLRPStopAuctionStub = nil
	fake.requestLRPStopAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeConvergerBBS) ConvergeTask(timeToClaim time.Duration, convergenceInterval time.Duration, timeToResolve time.Duration) {
	fake.convergeTaskMutex.Lock()
	fake.convergeTaskArgsForCall = append(fake.convergeTaskArgsForCall, struct {
		timeToClaim         time.Duration
		convergenceInterval time.Duration
		timeToResolve       time.Duration
	}{timeToClaim, convergenceInterval, timeToResolve})
	fake.convergeTaskMutex.Unlock()
	if fake.ConvergeTaskStub != nil {
		fake.ConvergeTaskStub(timeToClaim, convergenceInterval, timeToResolve)
	}
}

func (fake *FakeConvergerBBS) ConvergeTaskCallCount() int {
	fake.convergeTaskMutex.RLock()
	defer fake.convergeTaskMutex.RUnlock()
	return len(fake.convergeTaskArgsForCall)
}

func (fake *FakeConvergerBBS) ConvergeTaskArgsForCall(i int) (time.Duration, time.Duration, time.Duration) {
	fake.convergeTaskMutex.RLock()
	defer fake.convergeTaskMutex.RUnlock()
	return fake.convergeTaskArgsForCall[i].timeToClaim, fake.convergeTaskArgsForCall[i].convergenceInterval, fake.convergeTaskArgsForCall[i].timeToResolve
}

func (fake *FakeConvergerBBS) NewConvergeLock(convergerID string, interval time.Duration) ifrit.Runner {
	fake.newConvergeLockMutex.Lock()
	fake.newConvergeLockArgsForCall = append(fake.newConvergeLockArgsForCall, struct {
		convergerID string
		interval    time.Duration
	}{convergerID, interval})
	fake.newConvergeLockMutex.Unlock()
	if fake.NewConvergeLockStub != nil {
		return fake.NewConvergeLockStub(convergerID, interval)
	} else {
		return fake.newConvergeLockReturns.result1
	}
}

func (fake *FakeConvergerBBS) NewConvergeLockCallCount() int {
	fake.newConvergeLockMutex.RLock()
	defer fake.newConvergeLockMutex.RUnlock()
	return len(fake.newConvergeLockArgsForCall)
}

func (fake *FakeConvergerBBS) NewConvergeLockArgsForCall(i int) (string, time.Duration) {
	fake.newConvergeLockMutex.RLock()
	defer fake.newConvergeLockMutex.RUnlock()
	return fake.newConvergeLockArgsForCall[i].convergerID, fake.newConvergeLockArgsForCall[i].interval
}

func (fake *FakeConvergerBBS) NewConvergeLockReturns(result1 ifrit.Runner) {
	fake.NewConvergeLockStub = nil
	fake.newConvergeLockReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.ConvergerBBS = new(FakeConvergerBBS)
