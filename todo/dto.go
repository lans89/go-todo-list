package todoheader

import "time"

type ToDoHeaderDto struct {
	UUID         string          `json:"uuid"`
	Owner        string          `json:"owner"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	Active       bool            `json:"active"`
	CreationTime time.Time       `json:"creationTime"`
	FinishTime   *time.Time      `json:"finishTime"`
	UpdateTime   *time.Time      `json:"updateTime"`
	Details      []ToDoDetailDto `json:"details"`
}

type ToDoDetailDto struct {
	UUID            string     `json:"uuid"`
	StepDescription string     `json:"stepDescription"`
	IsFinished      bool       `json:"isFinished"`
	CreationTime    time.Time  `json:"creationTime"`
	FinishTime      *time.Time `json:"finishTime"`
}
