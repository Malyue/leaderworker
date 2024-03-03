package leaderworker

import (
	"context"
	"fmt"
	"leaderworker/worker"
	"time"
)

func (lw *LeaderWorker) OnLeader(h func(ctx context.Context)) {
	lw.mustNotStarted()
	lw.mu.Lock()
	defer lw.mu.Unlock()
	if lw.forLeaderUse.handlersOnLeader == nil {
		lw.forLeaderUse.handlersOnLeader = []func(ctx context.Context){}
	}
	lw.forLeaderUse.handlersOnLeader = append(lw.forLeaderUse.handlersOnLeader, h)
}

func (lw *LeaderWorker) LeaderHookOnWorkerAdd(h WorkerAddHandler) {
	lw.mustNotStarted()
	lw.mu.Lock()
	defer lw.mu.Unlock()
	if lw.forLeaderUse.handlersOnWorkerAdd == nil {
		lw.forLeaderUse.handlersOnWorkerAdd = []WorkerAddHandler{}
	}
	lw.forLeaderUse.handlersOnWorkerAdd = append(lw.forLeaderUse.handlersOnWorkerAdd, h)
}

func (lw *LeaderWorker) LeaderHookOnWorkerDelete(h WorkerDeleteHandler) {
	lw.mustNotStarted()
	lw.mu.Lock()
	defer lw.mu.Unlock()
	if lw.forLeaderUse.handlersOnWorkerDelete == nil {
		lw.forLeaderUse.handlersOnWorkerDelete = []WorkerDeleteHandler{}
	}
	lw.forLeaderUse.handlersOnWorkerDelete = append(lw.forLeaderUse.handlersOnWorkerDelete, h)
}

// AssignLogicTaskToWorker use rabbitmq to send msg
func (lw *LeaderWorker) AssignLogicTaskToWorker(ctx context.Context, workerID worker.ID, task worker.LogicTask) error {

}

func (lw *LeaderWorker) CancelLogicTask(ctx context.Context, logicTaskID worker.LogicTaskID) error {

}

func (lw *LeaderWorker) IsTaskBeginProcessed(ctx context.Context, logicTaskID worker.LogicTaskID) (bool, worker.ID) {
	// wait for init
	for {
		lw.mu.Lock()
		if lw.forLeaderUse.initialized {
			lw.mu.Unlock()
			break
		}
		lw.mu.Unlock()
		time.Sleep(lw.Cfg.Worker.RetryInterval)
	}

	lw.mu.Lock()
	defer lw.mu.Unlock()
	workerID, ok := lw.forLeaderUse.findWorkerByTask[logicTaskID]
	if !ok {
		return false, ""
	}
	// check valid worker
	_, workerExist := lw.forLeaderUse.Workers[workerID]
	if !workerExist {
		return false, ""
	}
	return true, workerID
}

func (lw *LeaderWorker) RegisterLeaderListener(i Listener) {
	lw.mustNotStarted()
	lw.mu.Lock()
	defer lw.mu.Unlock()
	lw.forLeaderUse.listeners = append(lw.forLeaderUse.listeners, i)
}

func (lw *LeaderWorker) mustBeLeader() {
	// check is leader
}

func (lw *LeaderWorker) mustNotStarted() {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	if lw.started {
		panic(fmt.Errorf("cannot invoke this mehot after started"))
	}
}
