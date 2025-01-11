package grpcclient

import (
	"context"

	pb "github.com/dezh-tech/panda/infrastructures/grpc_client/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	StringService pb.GetStringServiceClient
	conn          *grpc.ClientConn
}

func New(endpoint string) (*Client, error) {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &Client{
		StringService: pb.NewGetStringServiceClient(conn),
		conn:          conn,
	}, nil
}

func (c *Client) GetString() (string, error) {
	resp, err := c.StringService.GetString(context.Background(), &pb.GetStringRequest{})
	if err != nil {
		return "", err
	}

	return resp.Str, nil
}
