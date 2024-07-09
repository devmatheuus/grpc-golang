package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/devmatheuus/grpc/internal/database"
	"github.com/devmatheuus/grpc/internal/pb"
	"github.com/devmatheuus/grpc/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")

	if err != nil {
		panic(err)
	}

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		panic(err)
	}

	log.Println("gRPC server running at port :50051")

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
