package main

func main() {
	httpServer := NewHttpServer(":8000")
	go httpServer.Run()

	grpcSV := NewGRPCServer(":9000")
	grpcSV.Run()
}
