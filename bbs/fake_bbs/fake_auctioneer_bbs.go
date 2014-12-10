// This file was generated by counterfeiter
package fake_bbs

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/runtime-schema/bbs"
	"github.com/cloudfoundry-incubator/runtime-schema/models"
	"github.com/tedsuo/ifrit"
)

type FakeAuctioneerBBS struct {
	CellsStub        func() ([]models.CellPresence, error)
	cellsMutex       sync.RWMutex
	cellsArgsForCall []struct{}
	cellsReturns     struct {
		result1 []models.CellPresence
		result2 error
	}
	WatchForLRPStartAuctionStub        func() (<-chan models.LRPStartAuction, chan<- bool, <-chan error)
	watchForLRPStartAuctionMutex       sync.RWMutex
	watchForLRPStartAuctionArgsForCall []struct{}
	watchForLRPStartAuctionReturns     struct {
		result1 <-chan models.LRPStartAuction
		result2 chan<- bool
		result3 <-chan error
	}
	ClaimLRPStartAuctionStub        func(models.LRPStartAuction) error
	claimLRPStartAuctionMutex       sync.RWMutex
	claimLRPStartAuctionArgsForCall []struct {
		arg1 models.LRPStartAuction
	}
	claimLRPStartAuctionReturns struct {
		result1 error
	}
	ResolveLRPStartAuctionStub        func(models.LRPStartAuction) error
	resolveLRPStartAuctionMutex       sync.RWMutex
	resolveLRPStartAuctionArgsForCall []struct {
		arg1 models.LRPStartAuction
	}
	resolveLRPStartAuctionReturns struct {
		result1 error
	}
	WatchForLRPStopAuctionStub        func() (<-chan models.LRPStopAuction, chan<- bool, <-chan error)
	watchForLRPStopAuctionMutex       sync.RWMutex
	watchForLRPStopAuctionArgsForCall []struct{}
	watchForLRPStopAuctionReturns     struct {
		result1 <-chan models.LRPStopAuction
		result2 chan<- bool
		result3 <-chan error
	}
	ClaimLRPStopAuctionStub        func(models.LRPStopAuction) error
	claimLRPStopAuctionMutex       sync.RWMutex
	claimLRPStopAuctionArgsForCall []struct {
		arg1 models.LRPStopAuction
	}
	claimLRPStopAuctionReturns struct {
		result1 error
	}
	ResolveLRPStopAuctionStub        func(models.LRPStopAuction) error
	resolveLRPStopAuctionMutex       sync.RWMutex
	resolveLRPStopAuctionArgsForCall []struct {
		arg1 models.LRPStopAuction
	}
	resolveLRPStopAuctionReturns struct {
		result1 error
	}
	WatchForDesiredTaskStub        func() (<-chan models.Task, chan<- bool, <-chan error)
	watchForDesiredTaskMutex       sync.RWMutex
	watchForDesiredTaskArgsForCall []struct{}
	watchForDesiredTaskReturns     struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}
	CompleteTaskStub        func(taskGuid string, failed bool, failureReason string, result string) error
	completeTaskMutex       sync.RWMutex
	completeTaskArgsForCall []struct {
		taskGuid      string
		failed        bool
		failureReason string
		result        string
	}
	completeTaskReturns struct {
		result1 error
	}
	NewAuctioneerLockStub        func(auctioneerID string, interval time.Duration) ifrit.Runner
	newAuctioneerLockMutex       sync.RWMutex
	newAuctioneerLockArgsForCall []struct {
		auctioneerID string
		interval     time.Duration
	}
	newAuctioneerLockReturns struct {
		result1 ifrit.Runner
	}
}

func (fake *FakeAuctioneerBBS) Cells() ([]models.CellPresence, error) {
	fake.cellsMutex.Lock()
	fake.cellsArgsForCall = append(fake.cellsArgsForCall, struct{}{})
	fake.cellsMutex.Unlock()
	if fake.CellsStub != nil {
		return fake.CellsStub()
	} else {
		return fake.cellsReturns.result1, fake.cellsReturns.result2
	}
}

func (fake *FakeAuctioneerBBS) CellsCallCount() int {
	fake.cellsMutex.RLock()
	defer fake.cellsMutex.RUnlock()
	return len(fake.cellsArgsForCall)
}

func (fake *FakeAuctioneerBBS) CellsReturns(result1 []models.CellPresence, result2 error) {
	fake.CellsStub = nil
	fake.cellsReturns = struct {
		result1 []models.CellPresence
		result2 error
	}{result1, result2}
}

func (fake *FakeAuctioneerBBS) WatchForLRPStartAuction() (<-chan models.LRPStartAuction, chan<- bool, <-chan error) {
	fake.watchForLRPStartAuctionMutex.Lock()
	fake.watchForLRPStartAuctionArgsForCall = append(fake.watchForLRPStartAuctionArgsForCall, struct{}{})
	fake.watchForLRPStartAuctionMutex.Unlock()
	if fake.WatchForLRPStartAuctionStub != nil {
		return fake.WatchForLRPStartAuctionStub()
	} else {
		return fake.watchForLRPStartAuctionReturns.result1, fake.watchForLRPStartAuctionReturns.result2, fake.watchForLRPStartAuctionReturns.result3
	}
}

func (fake *FakeAuctioneerBBS) WatchForLRPStartAuctionCallCount() int {
	fake.watchForLRPStartAuctionMutex.RLock()
	defer fake.watchForLRPStartAuctionMutex.RUnlock()
	return len(fake.watchForLRPStartAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) WatchForLRPStartAuctionReturns(result1 <-chan models.LRPStartAuction, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForLRPStartAuctionStub = nil
	fake.watchForLRPStartAuctionReturns = struct {
		result1 <-chan models.LRPStartAuction
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeAuctioneerBBS) ClaimLRPStartAuction(arg1 models.LRPStartAuction) error {
	fake.claimLRPStartAuctionMutex.Lock()
	fake.claimLRPStartAuctionArgsForCall = append(fake.claimLRPStartAuctionArgsForCall, struct {
		arg1 models.LRPStartAuction
	}{arg1})
	fake.claimLRPStartAuctionMutex.Unlock()
	if fake.ClaimLRPStartAuctionStub != nil {
		return fake.ClaimLRPStartAuctionStub(arg1)
	} else {
		return fake.claimLRPStartAuctionReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) ClaimLRPStartAuctionCallCount() int {
	fake.claimLRPStartAuctionMutex.RLock()
	defer fake.claimLRPStartAuctionMutex.RUnlock()
	return len(fake.claimLRPStartAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) ClaimLRPStartAuctionArgsForCall(i int) models.LRPStartAuction {
	fake.claimLRPStartAuctionMutex.RLock()
	defer fake.claimLRPStartAuctionMutex.RUnlock()
	return fake.claimLRPStartAuctionArgsForCall[i].arg1
}

func (fake *FakeAuctioneerBBS) ClaimLRPStartAuctionReturns(result1 error) {
	fake.ClaimLRPStartAuctionStub = nil
	fake.claimLRPStartAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) ResolveLRPStartAuction(arg1 models.LRPStartAuction) error {
	fake.resolveLRPStartAuctionMutex.Lock()
	fake.resolveLRPStartAuctionArgsForCall = append(fake.resolveLRPStartAuctionArgsForCall, struct {
		arg1 models.LRPStartAuction
	}{arg1})
	fake.resolveLRPStartAuctionMutex.Unlock()
	if fake.ResolveLRPStartAuctionStub != nil {
		return fake.ResolveLRPStartAuctionStub(arg1)
	} else {
		return fake.resolveLRPStartAuctionReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) ResolveLRPStartAuctionCallCount() int {
	fake.resolveLRPStartAuctionMutex.RLock()
	defer fake.resolveLRPStartAuctionMutex.RUnlock()
	return len(fake.resolveLRPStartAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) ResolveLRPStartAuctionArgsForCall(i int) models.LRPStartAuction {
	fake.resolveLRPStartAuctionMutex.RLock()
	defer fake.resolveLRPStartAuctionMutex.RUnlock()
	return fake.resolveLRPStartAuctionArgsForCall[i].arg1
}

func (fake *FakeAuctioneerBBS) ResolveLRPStartAuctionReturns(result1 error) {
	fake.ResolveLRPStartAuctionStub = nil
	fake.resolveLRPStartAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) WatchForLRPStopAuction() (<-chan models.LRPStopAuction, chan<- bool, <-chan error) {
	fake.watchForLRPStopAuctionMutex.Lock()
	fake.watchForLRPStopAuctionArgsForCall = append(fake.watchForLRPStopAuctionArgsForCall, struct{}{})
	fake.watchForLRPStopAuctionMutex.Unlock()
	if fake.WatchForLRPStopAuctionStub != nil {
		return fake.WatchForLRPStopAuctionStub()
	} else {
		return fake.watchForLRPStopAuctionReturns.result1, fake.watchForLRPStopAuctionReturns.result2, fake.watchForLRPStopAuctionReturns.result3
	}
}

func (fake *FakeAuctioneerBBS) WatchForLRPStopAuctionCallCount() int {
	fake.watchForLRPStopAuctionMutex.RLock()
	defer fake.watchForLRPStopAuctionMutex.RUnlock()
	return len(fake.watchForLRPStopAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) WatchForLRPStopAuctionReturns(result1 <-chan models.LRPStopAuction, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForLRPStopAuctionStub = nil
	fake.watchForLRPStopAuctionReturns = struct {
		result1 <-chan models.LRPStopAuction
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeAuctioneerBBS) ClaimLRPStopAuction(arg1 models.LRPStopAuction) error {
	fake.claimLRPStopAuctionMutex.Lock()
	fake.claimLRPStopAuctionArgsForCall = append(fake.claimLRPStopAuctionArgsForCall, struct {
		arg1 models.LRPStopAuction
	}{arg1})
	fake.claimLRPStopAuctionMutex.Unlock()
	if fake.ClaimLRPStopAuctionStub != nil {
		return fake.ClaimLRPStopAuctionStub(arg1)
	} else {
		return fake.claimLRPStopAuctionReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) ClaimLRPStopAuctionCallCount() int {
	fake.claimLRPStopAuctionMutex.RLock()
	defer fake.claimLRPStopAuctionMutex.RUnlock()
	return len(fake.claimLRPStopAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) ClaimLRPStopAuctionArgsForCall(i int) models.LRPStopAuction {
	fake.claimLRPStopAuctionMutex.RLock()
	defer fake.claimLRPStopAuctionMutex.RUnlock()
	return fake.claimLRPStopAuctionArgsForCall[i].arg1
}

func (fake *FakeAuctioneerBBS) ClaimLRPStopAuctionReturns(result1 error) {
	fake.ClaimLRPStopAuctionStub = nil
	fake.claimLRPStopAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) ResolveLRPStopAuction(arg1 models.LRPStopAuction) error {
	fake.resolveLRPStopAuctionMutex.Lock()
	fake.resolveLRPStopAuctionArgsForCall = append(fake.resolveLRPStopAuctionArgsForCall, struct {
		arg1 models.LRPStopAuction
	}{arg1})
	fake.resolveLRPStopAuctionMutex.Unlock()
	if fake.ResolveLRPStopAuctionStub != nil {
		return fake.ResolveLRPStopAuctionStub(arg1)
	} else {
		return fake.resolveLRPStopAuctionReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) ResolveLRPStopAuctionCallCount() int {
	fake.resolveLRPStopAuctionMutex.RLock()
	defer fake.resolveLRPStopAuctionMutex.RUnlock()
	return len(fake.resolveLRPStopAuctionArgsForCall)
}

func (fake *FakeAuctioneerBBS) ResolveLRPStopAuctionArgsForCall(i int) models.LRPStopAuction {
	fake.resolveLRPStopAuctionMutex.RLock()
	defer fake.resolveLRPStopAuctionMutex.RUnlock()
	return fake.resolveLRPStopAuctionArgsForCall[i].arg1
}

func (fake *FakeAuctioneerBBS) ResolveLRPStopAuctionReturns(result1 error) {
	fake.ResolveLRPStopAuctionStub = nil
	fake.resolveLRPStopAuctionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) WatchForDesiredTask() (<-chan models.Task, chan<- bool, <-chan error) {
	fake.watchForDesiredTaskMutex.Lock()
	fake.watchForDesiredTaskArgsForCall = append(fake.watchForDesiredTaskArgsForCall, struct{}{})
	fake.watchForDesiredTaskMutex.Unlock()
	if fake.WatchForDesiredTaskStub != nil {
		return fake.WatchForDesiredTaskStub()
	} else {
		return fake.watchForDesiredTaskReturns.result1, fake.watchForDesiredTaskReturns.result2, fake.watchForDesiredTaskReturns.result3
	}
}

func (fake *FakeAuctioneerBBS) WatchForDesiredTaskCallCount() int {
	fake.watchForDesiredTaskMutex.RLock()
	defer fake.watchForDesiredTaskMutex.RUnlock()
	return len(fake.watchForDesiredTaskArgsForCall)
}

func (fake *FakeAuctioneerBBS) WatchForDesiredTaskReturns(result1 <-chan models.Task, result2 chan<- bool, result3 <-chan error) {
	fake.WatchForDesiredTaskStub = nil
	fake.watchForDesiredTaskReturns = struct {
		result1 <-chan models.Task
		result2 chan<- bool
		result3 <-chan error
	}{result1, result2, result3}
}

func (fake *FakeAuctioneerBBS) CompleteTask(taskGuid string, failed bool, failureReason string, result string) error {
	fake.completeTaskMutex.Lock()
	fake.completeTaskArgsForCall = append(fake.completeTaskArgsForCall, struct {
		taskGuid      string
		failed        bool
		failureReason string
		result        string
	}{taskGuid, failed, failureReason, result})
	fake.completeTaskMutex.Unlock()
	if fake.CompleteTaskStub != nil {
		return fake.CompleteTaskStub(taskGuid, failed, failureReason, result)
	} else {
		return fake.completeTaskReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) CompleteTaskCallCount() int {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return len(fake.completeTaskArgsForCall)
}

func (fake *FakeAuctioneerBBS) CompleteTaskArgsForCall(i int) (string, bool, string, string) {
	fake.completeTaskMutex.RLock()
	defer fake.completeTaskMutex.RUnlock()
	return fake.completeTaskArgsForCall[i].taskGuid, fake.completeTaskArgsForCall[i].failed, fake.completeTaskArgsForCall[i].failureReason, fake.completeTaskArgsForCall[i].result
}

func (fake *FakeAuctioneerBBS) CompleteTaskReturns(result1 error) {
	fake.CompleteTaskStub = nil
	fake.completeTaskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLock(auctioneerID string, interval time.Duration) ifrit.Runner {
	fake.newAuctioneerLockMutex.Lock()
	fake.newAuctioneerLockArgsForCall = append(fake.newAuctioneerLockArgsForCall, struct {
		auctioneerID string
		interval     time.Duration
	}{auctioneerID, interval})
	fake.newAuctioneerLockMutex.Unlock()
	if fake.NewAuctioneerLockStub != nil {
		return fake.NewAuctioneerLockStub(auctioneerID, interval)
	} else {
		return fake.newAuctioneerLockReturns.result1
	}
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockCallCount() int {
	fake.newAuctioneerLockMutex.RLock()
	defer fake.newAuctioneerLockMutex.RUnlock()
	return len(fake.newAuctioneerLockArgsForCall)
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockArgsForCall(i int) (string, time.Duration) {
	fake.newAuctioneerLockMutex.RLock()
	defer fake.newAuctioneerLockMutex.RUnlock()
	return fake.newAuctioneerLockArgsForCall[i].auctioneerID, fake.newAuctioneerLockArgsForCall[i].interval
}

func (fake *FakeAuctioneerBBS) NewAuctioneerLockReturns(result1 ifrit.Runner) {
	fake.NewAuctioneerLockStub = nil
	fake.newAuctioneerLockReturns = struct {
		result1 ifrit.Runner
	}{result1}
}

var _ bbs.AuctioneerBBS = new(FakeAuctioneerBBS)
