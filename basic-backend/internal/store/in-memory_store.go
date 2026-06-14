package store

import (
	"fmt"
	"sync"

	pb "github.com/DmitriiS-dev/basic-backend/proto"
)

type InMemoryStore struct {
	tasks  []*pb.Task
	nextId int32
	mu     sync.Mutex
}

// Like a constructor
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		tasks:  make([]*pb.Task, 0),
		nextId: 1,
	}
}

func (s *InMemoryStore) AddTask(t *pb.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock() // when this function exists then Unlock (works even if you return early)
	t.Id = s.nextId
	s.tasks = append(s.tasks, t)
	s.nextId += 1
	return nil
}

func (s *InMemoryStore) ListTasks() ([]*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	out := make([]*pb.Task, len(s.tasks))
	copy(out, s.tasks)
	return out, nil
}

func (s *InMemoryStore) GetTask(id int32) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		if task.Id == id {
			return task, nil
		}
	}
	return nil, fmt.Errorf("Task with Id %d not found", id)
}

func (s *InMemoryStore) DeleteTask(id int32) (*pb.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.Id == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return task, nil
		}
	}
	return nil, fmt.Errorf("Task with Id %d not found", id)

}
