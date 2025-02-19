package main

func main() {
	grpcServer := NewgRPCServer(":9000")
	go grpcServer.Run()

	httpServer := NewHttpServer(":8000")
	httpServer.Run()
}
