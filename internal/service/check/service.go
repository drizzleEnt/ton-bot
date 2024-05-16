package check

import (
	"google.golang.org/grpc"
)

type CheckService struct {
	conn *grpc.ClientConn
}

func NewCheckService(conn *grpc.ClientConn) *CheckService {
	return &CheckService{
		conn: conn,
	}
}
