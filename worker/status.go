package worker

type Status string

const LeaderStatus Status = "leader"
const WorkerStatus Status = "worker"

func (s Status) String() string {
	return string(s)
}
