package main

import (
	contador "contador/pkg"
	//"context"
	"log"
	"net"
	"google.golang.org/grpc"
)

type servidor struct {
	contador.UnimplementedContadorServer
}

func main() {
	// TODO: Implemente el inicio del servidor
	// Use NuevoServidor() definido en servidor_nucleo.go
	//contador.NuevoServidor()
	
	servidorReal := grpc.NewServer()
	servidor := contador.NuevoServidor()
	contador.RegisterContadorServer(servidorReal, &servidor)
	listen, err := net.Listen("tcp","localhost:12345")
	if err != nil{
		log.Fatalf("Fallo al escuchar %v",err)
	}
	if err := servidorReal.Serve(listen);err != nil{
		log.Fatalf("Fallo al servir: %v", err)
	}
}
