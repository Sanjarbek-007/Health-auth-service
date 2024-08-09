package main

import (
	"auth-service/api"
	"auth-service/config"
	pb "auth-service/genproto/user"
	"auth-service/service"
	"auth-service/storage"
	"auth-service/storage/postgres"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	cfg := config.Load().Server

	var wg sync.WaitGroup
	wg.Add(2)

	go RunService(&wg, db, &cfg)
	go RunRouter(&wg, db, &cfg)

	wg.Wait()
}

func RunService(wg *sync.WaitGroup, db storage.IStorage, cfg *config.ServerConfig) {
	defer wg.Done()

	lis, err := net.Listen("tcp", cfg.AUTH_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	u := service.NewUserService(db)
	server := grpc.NewServer()
	pb.RegisterUsersServer(server, u)

	log.Printf("Service is listening on port %s...\n", cfg.AUTH_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving auth service: %s", err)
	}
}

func RunRouter(wg *sync.WaitGroup, db storage.IStorage, cfg *config.ServerConfig) {
	defer wg.Done()

	router := api.NewRouter(db)
	log.Printf("Router is running on port %s...\n", cfg.AUTH_ROUTER_PORT)
	router.Run(cfg.AUTH_ROUTER_PORT)
}
