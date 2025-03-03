package main

import (
	"context"
	"log"

	"github.com/shch989/my-grpc-go-client/internal/adapter/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(log.Writer())

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:9090", opts...)

	if err != nil {
		log.Fatalln("Can not connect to gRPC server :", err)
	}

	defer conn.Close()

	helloAdapter, err := hello.NewHelloAdapter(conn)

	if err != nil {
		log.Fatalln("Can not create HelloAdapter :", err)
	}

	// runSayHello(helloAdapter, "Bruce Wayne")
	// runSayManyHello(helloAdapter, "Diana Prince")
	// runSayHelloToEveryone(helloAdapter, []string{"Andy", "Bill", "Christian", "Donny", "Ellen"})
	runSayHelloContinuous(helloAdapter, []string{"Anna", "Bella", "Carol", "Diana", "Emma"})
}

func runSayHello(adapter *hello.HelloAdapter, name string) {
	greet, err := adapter.SayHello(context.Background(), name)

	if err != nil {
		log.Fatalln("Can not call SayHello :", err)
	}

	log.Println(greet.Greet)
}

func runSayManyHello(adapter *hello.HelloAdapter, name string) {
	adapter.SayManyHellos(context.Background(), name)
}

func runSayHelloToEveryone(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloToEveryone(context.Background(), names)
}

func runSayHelloContinuous(adapter *hello.HelloAdapter, names []string) {
	adapter.SayHelloContinuous(context.Background(), names)
}
