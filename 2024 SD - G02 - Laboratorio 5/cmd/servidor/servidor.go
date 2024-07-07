package main

import (
	db "db/pkg"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
)

var (
	puerto     string
	nodos      string
	lider      string
	seguidores string
)

func main() {
	// Manejar argumentos para especificar el puerto
	flag.StringVar(&puerto, "p", "12345", "puerto en el cual escuchar")
	flag.StringVar(&nodos, "nodos", "", "lista de nodos separados por coma")
	flag.StringVar(&lider, "lider", "", "nodo líder")
	flag.StringVar(&seguidores, "seguidores", "", "lista de seguidores separados por coma")
	flag.Parse()

	if puerto == "" || nodos == "" || lider == "" || seguidores == "" {
		log.Fatal("El puerto, nodos, líder y seguidores son obligatorios")
	}

	listaNodos := strings.Split(nodos, ",")
	listaSeguidores := strings.Split(seguidores, ",")

	log.Printf("Nodos configurados: %v", listaNodos)
	log.Printf("Nodo líder: %s", lider)
	log.Printf("Seguidores configurados: %v", listaSeguidores)

	servidorReal := grpc.NewServer()
	servidor := db.NuevoServidor(listaNodos, lider, listaSeguidores)
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
