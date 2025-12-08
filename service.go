package main

import (
	"context"
	"fmt"
	"net"

	"github.com/savageking-io/savagedog/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type DogService struct {
	proto.UnimplementedDogServiceServer

	discord  *Discord
	hostname string
	port     uint16
}

func (d *DogService) Notification(ctx context.Context, req *proto.NotificationMessage) (*proto.NotificationResponse, error) {
	log.Tracef("DogService::Notification %+v", req)
	if d.discord == nil {
		return &proto.NotificationResponse{Code: 1, ErrorText: "Discord is nil"}, fmt.Errorf("discord is nil")
	}

	if err := d.discord.SendMessage(req.FromService, req); err != nil {
		return &proto.NotificationResponse{Code: 2, ErrorText: err.Error()}, err
	}

	return &proto.NotificationResponse{Code: 0}, nil
}

func (d *DogService) Init(hostname string, port uint16, discord *Discord) error {
	if discord == nil {
		return fmt.Errorf("discord is nil")
	}
	d.hostname = hostname
	d.port = port
	d.discord = discord

	return nil
}

func (d *DogService) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", d.hostname, d.port))
	if err != nil {
		return fmt.Errorf("error listening: %v", err)
	}

	srv := grpc.NewServer()

	proto.RegisterDogServiceServer(srv, d)

	log.Infof("gRPC server is running on %s:%d", d.hostname, d.port)
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	log.Infoln("Service stopped")

	return nil
}
