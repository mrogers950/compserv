package main

import (
	"flag"
	"log"
	"net"

	api "github.com/rhmdnd/compserv/pkg/api"
	config "github.com/rhmdnd/compserv/pkg/config"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configDir := flag.String("config-dir", "configs/",
		"Path to YAML configuration directory containing a config.yaml file.")
	configFile := flag.String("config-file", "config.yaml",
		"File name of the service config")
	flag.Parse()
	c := config.ParseConfig(*configDir, *configFile)
	connStr := config.GetDatabaseConnectionString(c)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}

	log.Printf("Connected to database: %v", db)

	appStr := c["app_host"] + ":" + c["app_port"]
	lis, err := net.Listen("tcp", appStr)
	if err != nil {
		log.Fatalf("Failed to listen to %s: %v", appStr, err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	api.RegisterComplianceServiceServer(grpcServer, api.NewServer(db))
	log.Printf("Server listening on %s", appStr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start grpc server %v", err)
	}
}
