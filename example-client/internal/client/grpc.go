package client

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/DmitriiS-dev/basic-backend/proto"
)

/*
Purpose of this file:
Client
├── connection
├── grpc client
├── ListTasks()
├── AddTask()
└── GetTask()
*/

type Client struct {
	conn *grpc.ClientConn
	api  pb.TaskServiceClient
}

func NewClient(serverFlagURL string) (*Client, error) {
	conn, err := grpc.NewClient(
		serverFlagURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		conn: conn,
		api:  pb.NewTaskServiceClient(conn),
	}, nil

}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) ListTasks(ctx context.Context) (*pb.TaskList, error) {
	return c.api.ListTasks(ctx, &pb.Empty{})
}

func (c *Client) AddTask(ctx context.Context, task *pb.Task) (*pb.Task, error) {
	return c.api.AddTask(
		ctx,
		&pb.AddTaskRequest{
			Task: task,
		},
	)
}

func (c *Client) GetTask(ctx context.Context, id int32) (*pb.Task, error) {
	return c.api.GetTask(
		ctx,
		&pb.GetTaskRequest{
			Id: id,
		},
	)
}
