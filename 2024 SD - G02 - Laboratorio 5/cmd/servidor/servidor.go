package main

import (
	db "db/pkg"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	puerto string
	nodos  []string
)

func main() {
	flag.StringVar(&puerto, "p", "12345", "puerto en el cual escuchar")
	flag.Parse()

	if puerto == "" {
		log.Fatal("El puerto es obligatorio")
	}

	nodos = []string{"localhost:12346", "localhost:12347"} // direcciones de nodos seguidores

	servidorReal := grpc.NewServer()
	servidor := db.NuevoServidor(nodos)
	db.RegisterBaseServer(servidorReal, &servidor)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", puerto))
	if err != nil {
		log.Fatalf("Fallo al escuchar: %v", err)
	}
	log.Printf("Servidor escuchando en el puerto %s", puerto)

	if err := servidorReal.Serve(listen); err != nil {
		log.Fatalf("Fallo al servir: %v", err)
	}
}
