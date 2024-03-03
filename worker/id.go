package worker

import "github.com/google/uuid"

type ID string

func NewID() ID {
	return ID(uuid.New().String())
}

func (id ID) String() string {
	return string(id)
}

type LogicTaskID string

func NewLogicTaskID() LogicTaskID {
	return LogicTaskID(NewID())
}

func (id LogicTaskID) String() string {
	return string(id)
}
