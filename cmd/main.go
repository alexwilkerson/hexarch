package main

import (
	"log"
	"os"

	gRPC "github.com/alexwilkerson/hexarch/internal/adapters/left/grpc"
	"github.com/alexwilkerson/hexarch/internal/adapters/right/db"
	"github.com/alexwilkerson/hexarch/internal/application/api"
	"github.com/alexwilkerson/hexarch/internal/application/arithmetic"
)

func main() {
	dbaseDriver := os.Getenv("DB_DRIVER")
	dbaseName := os.Getenv("DS_NAME")

	dbAdapter, err := db.NewAdapter(dbaseDriver, dbaseName)
	if err != nil {
		log.Fatalf("failed to initialize db connection: %v", err)
	}
	defer dbAdapter.CloseDbConnection()

	core := arithmetic.New()

	applicationAPI := api.NewApplication(dbAdapter, core)

	gRPCAdapter := gRPC.NewAdapter(applicationAPI)
	gRPCAdapter.Run()
}
