package cmd

import (
	"e-wallet-transaction/helpers"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ServeGRPC() {
	lis, err := net.Listen("tcp", ":"+helpers.GetEnv("GRPC_PORT", "5000"))
	if err != nil {
		log.Fatal("Failed to listen grpc port: ", err)
	}
	s := grpc.NewServer()

	// list methody

	logrus.Info("start grpc server on port: ", helpers.GetEnv("GRPC_PORT", "5000"), "")
	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve grpc port: ", err)
	}
}
