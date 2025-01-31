package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"

	contador "contador/pkg"

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

	// TODO: Validar que el puerto se encuentre entre 1 y 65535
	// De no estar en ese rango finalizar el programa

	puerto, _ := strconv.Atoi(puertoServidor)
	if puerto < 1 || puerto > 65535 {
		log.Fatalln("Error de puerto")
	}

	direccion := fmt.Sprintf("%s:%s", direccionServidor, puertoServidor)

	// Establece una conexión con el servidor
	conexion, _ := grpc.Dial(
		// dirección del servidor
		direccion,
		// indica que se debe conectar usando TCP sin SSL
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// bloquea el hilo hasta que la conexión se establezca
		grpc.WithBlock(),
	)
	// crea un nuevo cliente gRPC sobre la conexión
	cliente := contador.NewContadorClient(conexion)

	// PASO 1:
	// TODO: genere de forma secuencial 1000 incrementos y
	// al final verifique el valor del contador
	for i := 0; i < 1000; i++ {
		cliente.Incrementar(context.Background(), &contador.Vacio{})
	}
	resultado, _ := cliente.Obtener(context.Background(), &contador.Vacio{})
	fmt.Printf("Valor: %v\n", resultado)
	// PASO 2:
	// TODO: genere 1000 gorrutinas donde cada una produce un incremento y
	// al final verifique el valor del contador

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			cliente.Incrementar(context.Background(), &contador.Vacio{})
			wg.Done()
		}()

	}
	wg.Wait()
	resultado, _ = cliente.Obtener(context.Background(), &contador.Vacio{})
	fmt.Printf("Valor: %v\n", resultado)
}
