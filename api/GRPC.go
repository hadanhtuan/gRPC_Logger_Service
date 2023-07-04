package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"logger/model"
	pb "logger/protobuf"
	"net"
)

type LoggerServer struct {
	pb.UnimplementedLogServiceServer
	model.LogEntry
}

// TODO: tạm thời mới define 1 hàm, còn server side stream và client side stream chưa define
func (l *LoggerServer) WriteLog(ctx context.Context, req *pb.LogRequest) (*pb.LogResponse, error) {
	input := req.GetLogEntry()

	logEntry := model.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	createResult := model.LogEntryDB.Create(logEntry)

	res := &pb.LogResponse{
		Status:  createResult.Status,
		Message: createResult.Message,
	}

	return res, nil
}

func GRPCListen() {
	gRpcPort := "5151"
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", gRpcPort))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterLogServiceServer(s, &LoggerServer{})

	log.Printf("gRPC Server started on port %s", gRpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
}
