package main

import (
	"context"
	"github.com/Domains18/golang-grpc/invoicer"
	"google.golang.org/grpc"
	"log"
	"net"
)

type TheInvoicerServer struct {
}

func (s TheInvoicerServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf: []byte("Hello World!"),
		Docx: []byte("Hello World!"),
	}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error creating listener: %s", err)
	}
	S := grpc.NewServer{}
	invoicer.RegisterInvoicerServer()
}
