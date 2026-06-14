package store

import pb "github.com/DmitriiS-dev/basic-backend/proto"

// Interface to switch between Postgres & In-Memory:
type DatabaseType interface {
	AddTask(t *pb.Task) error
	ListTasks() ([]*pb.Task, error)
	GetTask(id int32) (*pb.Task, error)
	DeleteTask(id int32) (*pb.Task, error)
}
