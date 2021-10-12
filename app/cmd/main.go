package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/jmsilvadev/golangtechtask/api"
	"github.com/jmsilvadev/golangtechtask/pkg/config"
	"github.com/jmsilvadev/golangtechtask/pkg/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	config := config.GetDeaultConfig()
	StartServer(*config)

	log.Println("Shutting down service")
}

// StartServer starts a new gRPC server with TLS
func StartServer(config config.Config) context.Context {
	lis, err := net.Listen("tcp", config.ServerPort)
	if err != nil {
		config.Logger.Fatal("failed to listen: " + err.Error())
	}

	pathToCert, _ := filepath.Abs("certs")
	creds, err := credentials.NewServerTLSFromFile(pathToCert+"/server.crt", pathToCert+"/server.key")
	if err != nil {
		config.Logger.Fatal(err.Error())
	}
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	s := grpc.NewServer(opts...)

	api.RegisterVotingServiceServer(s, &server.Server{
		Db:     config.DBProvider,
		Logger: config.Logger,
	})

	config.Logger.Info("server listening at " + fmt.Sprint(lis.Addr()))

	go func() {
		if err := s.Serve(lis); err != nil {
			config.Logger.Fatal("failed to serve: " + err.Error())
		}
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, os.Interrupt, syscall.SIGTERM)
	sig := <-gracefulStop
	config.Logger.Warn("received a system call " + fmt.Sprint(sig))

	ctx, cancel := context.WithTimeout(config.Context, 5*time.Second)
	defer cancel()

	defer func() {
		s.GracefulStop()
		config.Logger.Info("clean shutdown")
	}()

	return ctx
}
