package service

import (
	"context"
	"fmt"

	"github.com/DmitriiS-dev/basic-backend/internal/store"
	pb "github.com/DmitriiS-dev/basic-backend/proto"
	//"google.golang.org/grpc/codes"
	// "google.golang.org/grpc/status"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	Store store.DatabaseType
}

// Controller Level functions
func (s *TaskServer) ListTasks(ctx context.Context, req *pb.Empty) (*pb.TaskList, error) {
	tasks, err := s.Store.ListTasks()
	if err != nil {
		return nil, fmt.Errorf("Failed to list tasks %w", err)
	}
	return &pb.TaskList{
		Tasks: tasks,
	}, nil
}

func (s *TaskServer) AddTask(
	ctx context.Context, req *pb.AddTaskRequest,
) (*pb.Task, error) {
	err := s.Store.AddTask(req.Task)
	if err != nil {
		return nil, err
	}

	return req.Task, nil
}

func (s *TaskServer) GetTask(
	ctx context.Context,
	req *pb.GetTaskRequest,
) (*pb.Task, error) {
	task, err := s.Store.GetTask(req.Id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskServer) DeleteTask(
	ctx context.Context,
	req *pb.DeleteTaskRequest,
) (*pb.Task, error) {
	task, err := s.Store.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// Can Add this later:
// if err != nil {
//     // Turns into an official gRPC NOT_FOUND status code
//     return nil, status.Errorf(codes.NotFound, "task with ID %d not found", req.Id)
// }
