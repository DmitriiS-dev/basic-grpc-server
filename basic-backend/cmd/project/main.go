package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/DmitriiS-dev/basic-backend/internal/service"
	"github.com/DmitriiS-dev/basic-backend/internal/store"
	pb "github.com/DmitriiS-dev/basic-backend/proto"
)

func main() {

	// Flag for Postgres or In-Memory store (Default: Postgres)
	dbType := flag.String("DatabaseType", "postgres", "Type of Database to use")
	flag.Parse() // read the cli args and update variable

	var db store.DatabaseType
	var err error

	if *dbType == "In-Memory" {
		fmt.Println("Setting up volatile in-memory storage...")
		db = store.NewInMemoryStore()
	} else {
		fmt.Println("Connecting to persistent Postgres instance...")

		host := os.Getenv("DB_HOST")
		port := os.Getenv("DB_PORT")
		user := os.Getenv("DB_USER")
		password := os.Getenv("DB_PASSWORD")
		dbname := os.Getenv("DB_NAME")

		db, err = store.NewPostgresStore(host, port, user, password, dbname)
		if err != nil {
			log.Fatal(err)
		}
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTaskServiceServer(
		grpcServer,
		&service.TaskServer{Store: db},
	)

	log.Println("gRPC Server Running on :50051")

	grpcServer.Serve(lis)
}
