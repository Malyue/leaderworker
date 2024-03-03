package leaderworker

import (
	"context"
	"leaderworker/worker"
	"path/filepath"
	"strings"
)

// deleteWorker key is $EtcdKeyPrefixWithSlash/type/$typID/$workerID
func (lw *LeaderWorker) deleteWorker(ctx context.Context, w worker.Worker) error {
	key := lw.makeEtcdWorkerKey(w.GetID(), w.GetType())
	_, err := lw.Etcd.Delete(ctx, key)
	return err
}

// ========== Worker Key =============
func (lw *LeaderWorker) makeEtcdWorkerKey(workerID worker.ID, typ worker.Type) string {
	keyPrefix := lw.makeEtcdWorkerKeyPrefix(typ)
	key := filepath.Join(keyPrefix, workerID.String())
	return key
}

func (lw *LeaderWorker) makeEtcdWorkerKeyPrefix(typ worker.Type) string {
	return filepath.Clean(filepath.Join(lw.Cfg.Worker.EtcdKeyPrefixWithSlash, "type", typ.String())) + "/"
}

func (lw *LeaderWorker) getEtcdWorkerIDFromWorkerKey(key string, typ worker.Type) worker.ID {
	prefix := lw.makeEtcdWorkerKeyPrefix(typ)
	return worker.ID(TrimPrefixes(key, prefix))
}

// ========== Dispatch Key =============
func (lw *LeaderWorker) makeEtcdWorkerGeneralDispatchPrefix() string {
	return filepath.Join(lw.Cfg.Worker.EtcdKeyPrefixWithSlash, "dispatch/worker") + "/"
}

func (lw *LeaderWorker) makeEtcdWorkerLogicTaskListenPrefix(workerID worker.ID) string {
	prefix := lw.makeEtcdWorkerGeneralDispatchPrefix()
	return filepath.Join(prefix, workerID.String(), "task") + "/"
}

// $prefix/dispatch/worker/$workerID/task/$logicTaskID
func (lw *LeaderWorker) makeEtcdWorkerTaskDispatchKey(workerID worker.ID, logicTaskID worker.LogicTaskID) string {
	prefix := lw.makeEtcdWorkerLogicTaskListenPrefix(workerID)
	return filepath.Join(prefix, logicTaskID.String())
}

// ========== Leader Logic Task Cancel Key =============
// $prefix/dispatch/cancel-task/$logicTaskID
func (lw *LeaderWorker) makeEtcdLeaderLogicTaskCancelKey(logicTaskID worker.LogicTaskID) string {
	return filepath.Join(lw.makeEtcdLeaderLogicTaskCancelKeyPrefix(), logicTaskID.String())
}

func (lw *LeaderWorker) makeEtcdLeaderLogicTaskCancelKeyPrefix() string {
	return filepath.Join(lw.Cfg.Leader.EtcdKeyPrefixWithSlash, "dispatch/cancel-task") + "/"
}

func (lw *LeaderWorker) getLogicTaskIDFromLeaderCancelKey(key string) worker.LogicTaskID {
	return worker.LogicTaskID(TrimPrefixes(key, lw.makeEtcdLeaderLogicTaskCancelKeyPrefix()))
}

// ========== Worker Logic Task Cancel Key =============
// $prefix/dispatch/worker/$workerID/cancel-task/$logicTaskID
func (lw *LeaderWorker) makeEtcdWorkerLogicTaskCancelListenPrefix(workerID worker.ID) string {
	return filepath.Join(lw.makeEtcdWorkerGeneralDispatchPrefix(), workerID.String(), "cancel-task") + "/"
}

func (lw *LeaderWorker) makeEtcdWorkerLogicTaskCancelKey(workerID worker.ID, logicTaskID worker.LogicTaskID) string {
	return filepath.Join(lw.makeEtcdWorkerLogicTaskCancelListenPrefix(workerID), logicTaskID.String())
}

func (lw *LeaderWorker) getLogicTaskIDFromWorkerCancelKey(workerID worker.ID, key string) worker.LogicTaskID {
	prefix := lw.makeEtcdWorkerLogicTaskCancelListenPrefix(workerID)
	if !strings.HasPrefix(key, prefix) {
		return ""
	}
	return worker.LogicTaskID(TrimPrefixes(key, prefix))
}

// ========== heartbeat Key =============
func (lw *LeaderWorker) makeEtcdWorkerHeartbeatKeyPrefix() string {
	return filepath.Clean(filepath.Join(lw.Cfg.Worker.EtcdKeyPrefixWithSlash, "heartbeat")) + "/"
}

func (lw *LeaderWorker) makeEtcdWorkerHeartbeatKey(workerID worker.ID) string {
	prefix := lw.makeEtcdWorkerHeartbeatKeyPrefix()
	return filepath.Join(prefix, workerID.String())
}

func (lw *LeaderWorker) getWorkerIDFromEtcdWorkerHeartbeatKey(key string) worker.ID {
	prefix := lw.makeEtcdWorkerHeartbeatKeyPrefix()
	return worker.ID(TrimPrefixes(key, prefix))
}

func (lw *LeaderWorker) getWorkerIDFromIncomingKey(key string) worker.ID {
	prefix := lw.makeEtcdWorkerGeneralDispatchPrefix()
	if !strings.HasPrefix(key, prefix) {
		return ""
	}
	workerIDAndSuffix := TrimPrefixes(key, prefix)
	workerIDAndLogicTaskID := strings.Split(workerIDAndSuffix, "/task/")
	if len(workerIDAndLogicTaskID) != 2 {
		return ""
	}
	return worker.ID(workerIDAndLogicTaskID[0])
}

func (lw *LeaderWorker) getWorkerIDFromWorkerLogicTaskCancelKey(key string) worker.ID {
	prefix := lw.makeEtcdWorkerGeneralDispatchPrefix()
	if !strings.HasPrefix(key, prefix) {
		return ""
	}
	workerIDAndSuffix := TrimPrefixes(key, prefix)
	workerIDAndLogicTaskID := strings.Split(workerIDAndSuffix, "/cancel-task/")
	if len(workerIDAndLogicTaskID) != 2 {
		return ""
	}
	return worker.ID(workerIDAndLogicTaskID[0])
}
