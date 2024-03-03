package worker

type LogicTask interface {
	GetLogicID() LogicTaskID
	GetData() []byte
}

type defaultTask struct {
	logicID LogicTaskID
	data    []byte
}

func NewLogicTask(logicID LogicTaskID, data []byte) LogicTask {
	return &defaultTask{logicID: logicID, data: data}
}

func (d *defaultTask) GetLogicID() LogicTaskID {
	return d.logicID
}

func (d *defaultTask) GetData() []byte {
	return d.data
}
