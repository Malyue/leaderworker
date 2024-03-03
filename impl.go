package leaderworker

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"leaderworker/worker"
	"log"
	"sync"
)

type LeaderWorker struct {
	Log log.Logger
	Cfg *config

	mu           sync.Mutex
	started      bool
	forLeaderUse forLeaderUse
	forWorkerUse forWorkerUse

	Etcd *clientv3.Client
}

type forLeaderUse struct {
	Workers          map[worker.ID]worker.Worker
	initialized      bool
	findWorkerByTask map[worker.LogicTaskID]worker.ID
	findTaskByWorker map[worker.ID]map[worker.LogicTaskID]struct{}

	listeners []Listener

	handlersOnLeader       []func(ctx context.Context)
	handlersOnWorkerAdd    []WorkerAddHandler
	handlersOnWorkerDelete []WorkerDeleteHandler
}

type forWorkerUse struct {
	Workers                map[worker.ID]workerWithCancel
	handlersOnWorkerDelete []WorkerDeleteHandler
}

func (lw *LeaderWorker) Init(ctx context.Context) error {
	lw.forLeaderUse.Workers = make(map[worker.ID]worker.Worker)
	if len(lw.Cfg.Leader.EtcdKeyPrefixWithSlash) == 0 {

	}
}
