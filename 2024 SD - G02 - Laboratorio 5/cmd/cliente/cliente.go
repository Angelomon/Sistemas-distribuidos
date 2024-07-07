package main

import (
	"context"
	db "db/pkg"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

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
	nodos             string
)

func main() {
	flag.StringVar(&puertoServidor, "p", PUERTO_SERVIDOR_PREDETERMINADO, "puerto a conectarse")
	flag.StringVar(&direccionServidor, "d", DIRECCION_SERVIDOR_PREDETERMINADA, "dirección del servidor")
	flag.StringVar(&nodos, "nodos", "", "lista de nodos separados por coma")
	flag.Parse()

	// Validar que el puerto se encuentre entre 1 y 65535
	puerto, err := strconv.Atoi(puertoServidor)
	if err != nil || puerto < 1 || puerto > 65535 {
		log.Fatalf("El puerto debe estar entre 1 y 65535. Puerto proporcionado: %s", puertoServidor)
	}

	listaNodos := strings.Split(nodos, ",")
	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)

	// Establece una conexión inicial con el servidor
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
	log.Printf("Ingreso 1")
	// Ingresar clave: 1, valor: Sistemas Distribuidos
	err = putWithRedirect(ctx, cliente, &db.ParametroPut{Clave: "1", Valor: []byte("Sistemas Distribuidos")}, listaNodos)
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}
	log.Printf("Ingreso 2")
	// Ingresar clave: 2, valor: Sist. Operativos
	err = putWithRedirect(ctx, cliente, &db.ParametroPut{Clave: "2", Valor: []byte("Sist. Operativos")}, listaNodos)
	if err != nil {
		log.Fatalf("Error en Put: %v", err)
	}
	log.Printf("Get all")
	// Obtener todos los datos almacenados e imprimir lo devuelto
	getAllResp, err := cliente.GetAll(ctx, &db.ParametroGetAll{})
	if err != nil {
		log.Fatalf("Error en GetAll: %v", err)
	}
	for _, kv := range getAllResp.ClaveValor {
		fmt.Printf("Clave: %s, Valor: %s\n", kv.Clave, string(kv.Valor))
	}
	log.Printf("Procesando")
}

func putWithRedirect(ctx context.Context, cliente db.BaseClient, msg *db.ParametroPut, nodos []string) error {
	for _, nodo := range nodos {
		direccion := fmt.Sprintf("%s", nodo)
		log.Printf("Intentando conectar a %s", direccion) // Registro de depuración

		conexion, err := grpc.Dial(direccion, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Printf("No se pudo conectar a %s: %v", direccion, err) // Registro de error
			continue
		}
		defer conexion.Close()

		cliente = db.NewBaseClient(conexion)
		_, err = cliente.Put(ctx, msg)
		if err != nil {
			log.Printf("Error en Put en nodo %s: %v", nodo, err) // Registro de error
			if grpc.Code(err) == grpc.Code(fmt.Errorf("redireccionar a nodo %s", nodo)) {
				log.Printf("Redirigiendo a nodo %s", nodo) // Registro de redirección
				continue
			}
			return err
		}
		log.Printf("Put exitoso en nodo %s", nodo) // Registro de éxito
		return nil
	}
	return nil
	//return fmt.Errorf("no se pudo completar la operación Put en ningún nodo")
}

func getWithRedirect(ctx context.Context, cliente db.BaseClient, msg *db.ParametroGet, nodos []string) (*db.ResultadoGet, error) {
	for _, nodo := range nodos {
		direccion := fmt.Sprintf("%s", nodo)
		log.Printf("Intentando conectar a %s", direccion) // Registro de depuración

		conexion, err := grpc.Dial(direccion, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
		if err != nil {
			log.Printf("No se pudo conectar a %s: %v", direccion, err) // Registro de error
			continue
		}
		defer conexion.Close()

		cliente = db.NewBaseClient(conexion)
		resp, err := cliente.Get(ctx, msg)
		if err != nil {
			log.Printf("Error en Get en nodo %s: %v", nodo, err) // Registro de error
			if grpc.Code(err) == grpc.Code(fmt.Errorf("redireccionar a nodo %s", nodo)) {
				log.Printf("Redirigiendo a nodo %s", nodo) // Registro de redirección
				continue
			}
			return nil, err
		}
		log.Printf("Get exitoso en nodo %s", nodo) // Registro de éxito
		return resp, nil
	}
	return nil, fmt.Errorf("no se pudo completar la operación Get en ningún nodo")
}
