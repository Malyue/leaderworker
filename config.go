package leaderworker

import "time"

type config struct {
	Leader leaderConfig `json:"leader"`
	Worker workerConfig `json:"worker"`
}

type leaderConfig struct {
	IsWorker               bool          `json:"is_worker"`
	CleanupInterval        time.Duration `json:"cleanup_interval"`
	RetryInterval          time.Duration `json:"retry_interval"`
	EtcdKeyPrefixWithSlash string        `json:"etcd_key_prefix_with_slash"`
}

type workerConfig struct {
	Candidate              candidateWorkerConfig `json:"candidate"`
	EtcdKeyPrefixWithSlash string                `json:"etcd_key_prefix_with_slash"`
	LivenessProbeInterval  time.Duration         `json:"liveness_probe_interval"`
	Task                   workerTaskConfig      `json:"task"`
	Heartbeat              workerHeartbeatConfig `json:"heartbeat"`
	RetryInterval          time.Duration         `json:"retry_interval"`
}

type workerTaskConfig struct {
	RetryDeleteTaskInterval time.Duration `json:"retry_delete_task_interval"`
}

type workerHeartbeatConfig struct {
	ReportInterval                     time.Duration `json:"report_interval"`
	AllowedMaxContinueLostContactTimes int           `json:"allowed_max_continue_lost_contact_times"`
}

type candidateWorkerConfig struct {
	ThresholdToBecomeOfficial time.Duration `json:"threshold_to_become_official"`
}
