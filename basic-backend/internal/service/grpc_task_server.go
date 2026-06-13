package service

import (
	"context"

	"github.com/DmitriiS-dev/basic-backend/internal/store"
	pb "github.com/DmitriiS-dev/basic-backend/proto"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
}

// Controller Level functions
func (s *TaskServer) ListTasks(ctx context.Context, req *pb.Empty) (*pb.TaskList, error) {
	tasks := store.GetTasks()
	return &pb.TaskList{
		Tasks: tasks,
	}, nil
}

func (s *TaskServer) AddTask(
	ctx context.Context, req *pb.AddTaskRequest,
) (*pb.Task, error) {
	store.AddTask(req.Task)

	return req.Task, nil
}

func (s *TaskServer) GetTask(
	ctx context.Context,
	req *pb.GetTaskRequest,
) (*pb.Task, error) {
	task, err := store.GetTask(req.Id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskServer) DeleteTask(
	ctx context.Context,
	req *pb.DeleteTaskRequest,
) (*pb.Task, error) {
	task, err := store.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}
	return task, nil
}
