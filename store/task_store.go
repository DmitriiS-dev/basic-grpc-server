package store

import (
	pb "github.com/dmitriis-dev/basic-backend/proto"
	"sync"
)

var (
	tasks = make([]*pb.Task, 0)
	mu    sync.Mutex
)

func AddTask(t *pb.Task) {
	mu.Lock()
	defer mu.Unlock()
	tasks = append(tasks, t)
}

func GetTasks() []*pb.Task {
	mu.Lock()
	defer mu.Unlock()

	out := make([]*pb.Task, len(tasks))
	copy(out, tasks)
	return out
}

func DeleteTask(id int32) {
	mu.Lock()
	defer mu.Unlock()

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

}
