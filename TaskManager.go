package interview

import (
	"container/list"
	"database/sql"
	"sync"
	"time"
)

type TaskManager struct {
	queue    list.List // container/list
	running  map[Task]struct{}
	blocked  map[Task]struct{}
	conn     *sql.Conn
	Incoming chan Task // alternative way to schedule task
}

var instance *TaskManager = nil

func GetInstance() *TaskManager {
	new(sync.Once).Do(func() {
		instance = new(TaskManager)
		instance.queue = list.List{}
		instance.running = map[Task]struct{}{}
		instance.blocked = map[Task]struct{}{}
		instance.Incoming = make(chan Task)
		for task := range instance.Incoming {
			instance.SkeduleTask(task)
		}
		go func() {
			for {
				instance.checkRun()
			}
		}()
	})
	return instance
}

func (me TaskManager) RunTask(task Task) {
	var tx, _ = me.conn.BeginTx(nil, nil)
	if notContains(me.running, *task.dependsOn) {
		me.running[task] = struct{}{}
		task.status = "Running"
		go func() {
			task.work(&task, tx)
		}()
		saveRecord(tx, TaskHistoryRecord{
			record: "task " + task.name + " " + task.timestamp.String() + " " +
				" started " + time.Now().String(),
		})
	} else {
		me.blocked[task] = struct{}{}
		task.status = "Blocked"
		saveRecord(tx, TaskHistoryRecord{
			record: "task " + task.name + " " + task.timestamp.String() + " " +
				" blocked " + time.Now().String(),
		})
	}
	tx.Commit()
}

func (me TaskManager) taskFinished(task Task) {
	var tx, _ = me.conn.BeginTx(nil, nil)
	delete(me.running, task)
	saveRecord(tx, TaskHistoryRecord{
		record: "task " + task.name + " " + task.timestamp.String() + " " +
			" finished " + time.Now().String(),
	})
	for t := range me.blocked {
		if t.dependsOn == &task {
			delete(me.running, t)
			me.RunTask(t)
		}
	}
	tx.Commit()
}

func notContains(source map[Task]struct{}, target Task) bool {
	_, ok := source[target]
	return ok
}

func (me TaskManager) SkeduleTask(task Task) {
	for e := me.queue.Front(); e != nil; e = e.Next() {
		var t Task = e.Value.(Task)
		if task.timestamp.Before(t.timestamp) {
			me.queue.InsertBefore(task, e)
		}
	}
}

func (me TaskManager) IsRunning(task Task) bool {
	_, ok := me.running[task]
	return ok
}

func (me TaskManager) checkRun() {
	for e := me.queue.Front(); e != nil; e = e.Next() {
		var task Task = e.Value.(Task)
		if task.timestamp.After(time.Now()) {
			me.RunTask(task)
		} else {
			return
		}
	}
}
