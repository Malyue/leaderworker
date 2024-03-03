package leaderworker

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	"leaderworker/worker"
)

type Event struct {
	Type         mvccpb.Event_EventType
	WorkerID     worker.ID
	LogicTaskIDs []worker.LogicTaskID
}

type workerWithCancel struct {
	Worker     worker.Worker
	Ctx        context.Context
	CancelFunc context.CancelFunc
	LogicTasks map[worker.LogicTaskID]logicTaskWithCtx
}

type logicTaskWithCtx struct {
	LogicTask worker.LogicTask
	Ctx       context.Context
}

type (
	WorkerAddHandler    func(ctx context.Context, ev Event)
	WorkerDeleteHandler func(ctx context.Context, ev Event)
)

type Listener interface {
	BeforeExecOnLeader(ctx context.Context)
	AfterExecOnLeader(ctx context.Context)
}

var _ Listener = (*DefaultListener)(nil)

type DefaultListener struct {
	BeforeExecOnLeaderFunc func(ctx context.Context)
	AfterExecOnLeaderFunc  func(ctx context.Context)
}

func (l *DefaultListener) BeforeExecOnLeader(ctx context.Context) {
	if l.BeforeExecOnLeaderFunc == nil {
		return
	}
	l.BeforeExecOnLeaderFunc(ctx)
}

func (l *DefaultListener) AfterExecOnLeader(ctx context.Context) {
	if l.AfterExecOnLeaderFunc == nil {
		return
	}
	l.AfterExecOnLeaderFunc(ctx)
}
