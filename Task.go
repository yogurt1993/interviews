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

func (t Task) Finish() {
	t.status = "Finished"
	GetInstance().taskFinished(t)
}

func (t Task) Reschedule(newTime time.Time) {
	t.status = "NotStarted"
	t.timestamp = newTime
	GetInstance().taskFinished(t)
	GetInstance().SkeduleTask(t)
}
