package leaderworker

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"leaderworker/worker"
)

type Interface interface {
	ForLeaderUseInterface
	ForWorkerUseInterface
	//// Keep the leader status
	//Keep() <-chan error
	//// Race start the leader election
	//Race() (<-chan bool, <-chan error)
	//// List the all nodes
	//List() ([]Node, error)
	//Remove(id ...string) (int64, error)
	//OnLeader()
}

type ForLeaderUseInterface interface {
	OnLeader(func(ctx context.Context))
	//LeaderHookOnWorkerAdd(WorkerAddHandler)
	//LeaderHookOnWorkerDelete(WorkerDeleteHandler)
	AssignLogicTaskToWorker(ctx context.Context, workerID worker.ID, logicTask worker.LogicTask) error
	IsTaskBeginProcessed(ctx context.Context, logicTaskID worker.LogicTaskID) (bool, worker.ID)
	RegisterLeaderListener(l Listener)
	LoadCancelingTasks(ctx context.Context)
}

type ForWorkerUseInterface interface {
	RegisterCandidateWorker(ctx context.Context, w worker.Worker) error
	WorkerHookOnWorkerDelete(WorkerDeleteHandler)
}

type GeneralInterface interface {
	ListWorkers(ctx context.Context, workerTypes ...worker.Type) ([]worker.Worker, error)
	ListenPrefix(ctx context.Context, prefix string, putHandler, deleteHandler func(context.Context, *clientv3.Event))
	Start()
	CancelLogicTask(ctx context.Context, logicTaskID worker.LogicTaskID) error
}
