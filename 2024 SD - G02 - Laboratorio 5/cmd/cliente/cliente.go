package main

import (
	db "db/pkg"
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"
	"math/rand"

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

	puerto, err := strconv.Atoi(puertoServidor)
	if err != nil || puerto < 1 || puerto > 65535 {
		log.Fatalf("El puerto debe estar entre 1 y 65535. Puerto proporcionado: %s", puertoServidor)
	}

	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)
	log.Printf("Conectando a %s", direccion)

	conexion, err := grpc.Dial(direccion, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v", err)
	}
	defer conexion.Close()
	log.Println("Conexión establecida con el servidor")

	cliente := db.NewBaseClient(conexion)

	ctx := context.Background()

	// Crear generador de números aleatorios
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Insertar claves aleatorias para probar particionamiento
	log.Println("Insertando claves aleatorias...")
	for i := 0; i < 100; i++ {
		clave := strconv.Itoa(rng.Intn(10000))
		valor := []byte(fmt.Sprintf("Valor%d", i))
		_, err := cliente.Put(ctx, &db.ParametroPut{Clave: clave, Valor: valor})
		if err != nil {
			log.Fatalf("Error en Put: %v", err)
		}
		if i % 10 == 0 {
			log.Printf("%d claves insertadas...", i)
		}
	}
	log.Println("Claves insertadas")

	// Obtener todas las claves y valores almacenados
	log.Println("Obteniendo todas las claves...")
	getAllResp, err := cliente.GetAll(ctx, &db.ParametroGetAll{})
	if err != nil {
		log.Fatalf("Error en GetAll: %v", err)
	}
	for _, kv := range getAllResp.ClaveValor {
		fmt.Printf("Clave: %s, Valor: %s\n", kv.Clave, string(kv.Valor))
	}
	log.Println("Obtención de claves completada")
}