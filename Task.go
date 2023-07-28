package interview

import (
	"database/sql"
	"time"
)

type UnitOfWork func(task *Task, tx *sql.Tx)

type Task struct {
	name      string
	timestamp time.Time
	status    string // NotStarted, Blocked, Working, Waiting, Finished, Closed
	dependsOn *Task
	work      UnitOfWork // will call Finish or Reschedule in the end/
}

func (this Task) Finish() {
	this.status = "Finished"
	GetInstance().taskFinished(this)
}

func (this Task) Reschedule(newTime time.Time) {
	this.status = "NotStarted"
	this.timestamp = newTime
	GetInstance().taskFinished(this)
	GetInstance().SkeduleTask(this)
}
