package main

import (
	db "db/pkg"
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	DIRECCION_SERVIDOR_PREDETERMINADA string = "localhost"
	PUERTO_SERVIDOR_PREDETERMINADO    string = "12345"
)

var (
	puertoServidor    string
	direccionServidor string
)

func main() {
	flag.StringVar(&puertoServidor, "p", PUERTO_SERVIDOR_PREDETERMINADO, "puerto a conectarse")
	flag.StringVar(&direccionServidor, "d", DIRECCION_SERVIDOR_PREDETERMINADA, "dirección del servidor")
	flag.Parse()

	// Validar que el puerto se encuentre entre 1 y 65535
	puerto, err := strconv.Atoi(puertoServidor)
	if err != nil || puerto < 1 || puerto > 65535 {
		log.Fatalf("El puerto debe estar entre 1 y 65535. Puerto proporcionado: %s", puertoServidor)
	}

	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)

	// Establece una conexión con el servidor
	conexion, err := grpc.Dial(
		direccion,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conexion.Close()

	cliente := db.NewBaseClient(conexion)

	// Interacciones con el servidor
	ctx := context.Background()

	// Ingresar clave: 1, valor: Sistemas Distribuidos
	_, err = cliente.Put(ctx, &db.ParametroPut{Clave: "1", Valor: []byte("Sistemas Distribuidos")})
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}

	// Ingresar clave: 2, valor: Sist. Operativos
	_, err = cliente.Put(ctx, &db.ParametroPut{Clave: "2", Valor: []byte("Sist. Operativos")})
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}
	// Ingresar clave: 3, valor: S.O
	_, err = cliente.Put(ctx, &db.ParametroPut{Clave: "3", Valor: []byte("S.O")})
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}
	// Obtener clave: 2 e imprimir lo devuelto
	getResp, err := cliente.Get(ctx, &db.ParametroGet{Clave: "2"})
	if err != nil {
		log.Printf("Error en Get: %v", err)
	} else {
		fmt.Printf("Get clave 2: %s\n", string(getResp.Valor))
	}
	
	// Obtener clave: 3 e imprimir lo devuelto
	getResp, err = cliente.Get(ctx, &db.ParametroGet{Clave: "3"})
	if err != nil {
		log.Printf("Error en Get: %v", err)
	} else {
		fmt.Printf("Get clave 3: %s\n", string(getResp.Valor))
	}

	// Ingresar clave: 2, valor: Sistemas Operativos
	_, err = cliente.Put(ctx, &db.ParametroPut{Clave: "2", Valor: []byte("Sistemas Operativos")})
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}

	// Obtener todos los datos almacenados e imprimir lo devuelto
	getAllResp, err := cliente.GetAll(ctx, &db.ParametroGetAll{})
	if err != nil {
		log.Fatalf("Error en GetAll: %v", err)
	}
	for _, kv := range getAllResp.ClaveValor {
		fmt.Printf("Clave: %s, Valor: %s\n", kv.Clave, string(kv.Valor))
	}
}