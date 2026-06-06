package store

import (
	"fmt"
	"sync"

	pb "github.com/DmitriiS-dev/basic-backend/proto"
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

func GetTask(id int32) (*pb.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, task := range tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("Task with Id %d not found", id)
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
