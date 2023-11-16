package main

import (
	"context"
	"github.com/Domains18/golang-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TheInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s TheInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte(req.From),
		Docx: []byte(req.To),
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error creating listener: %s", err)
	}

	serverRegistrar := grpc.NewServer()
	service := TheInvoicerServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, &service)
	if err := serverRegistrar.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %s", err)
	}
}
