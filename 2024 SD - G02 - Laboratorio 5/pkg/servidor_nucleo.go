package pkg

import (
	"context"
	"fmt"
	"log"
	"sync"

	"hash/fnv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// La implementación del servidor
type Servidor struct {
	UnimplementedBaseServer
	store      map[string][]byte
	mutex      sync.RWMutex
	nodos      []string
	lider      string
	seguidores []string
}

// NuevoServidor retorna una instancia del servidor
func NuevoServidor(nodos []string, lider string, seguidores []string) Servidor {
	return Servidor{
		store:      make(map[string][]byte),
		nodos:      nodos,
		lider:      lider,
		seguidores: seguidores,
	}
}

// Función hash para distribuir las claves entre nodos
func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
func conectarServidor(puerto string) BaseClient {
	// Crear conexion
	conn, err := grpc.Dial(puerto, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	// Crear cliente
	cliente := NewBaseClient(conn)
	return cliente
}

// Implementación de Put definido en el archivo `.proto`.
func (s *Servidor) Put(ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
	nodo := hash(msg.Clave) % len(s.nodos)

	if s.nodos[nodo] != s.lider {
		// Redirigir al nodo líder apropiado
		// Implementar la lógica para redirigir la solicitud al nodo líder correcto
		cliente := conectarServidor(s.lider)
		return cliente.Put(context.Background(), msg)
		//return nil, fmt.Errorf("redireccionar a nodo %s", s.nodos[nodo])
	}

	// Guardar en el nodo líder y replicar a los seguidores
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[msg.Clave] = msg.Valor

	// Replicar a los seguidores
	for _, seguidor := range s.seguidores {
		conn, err := grpc.Dial(seguidor, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("No se pudo conectar a seguidor %s: %v", seguidor, err)
			continue
		}
		defer conn.Close()
		client := NewBaseClient(conn)
		replicaMsg := &ReplicaPut{
			Clave: msg.Clave,
			Valor: msg.Valor,
			Lider: s.lider,
		}
		_, err = client.Replicate(ctx, replicaMsg)
		if err != nil {
			log.Printf("Error replicando a seguidor %s: %v", seguidor, err)
		}
	}

	return &ResultadoPut{}, nil
}

// Implementación de Get definido en el archivo `.proto`.
func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
	nodo := hash(msg.Clave) % len(s.nodos)
	if s.nodos[nodo] != s.lider {
		// Redirigir al nodo líder apropiado
		// Implementar la lógica para redirigir la solicitud al nodo líder correcto

		return nil, fmt.Errorf("redireccionar a nodo %s", s.nodos[nodo])
	}

	s.mutex.RLock()
	defer s.mutex.RUnlock()
	valor, exists := s.store[msg.Clave]
	if !exists {
		return nil, fmt.Errorf("clave no encontrada")
	}
	return &ResultadoGet{Valor: valor}, nil
}

// Implementación de GetAll definido en el archivo `.proto`.
func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result := &ResultadoGetAll{}
	for clave, valor := range s.store {
		result.ClaveValor = append(result.ClaveValor, &ClaveValor{Clave: clave, Valor: valor})
	}
	return result, nil
}

// Implementación de Replicate definido en el archivo `.proto`.
func (s *Servidor) Replicate(ctx context.Context, msg *ReplicaPut) (*ReplicaResponse, error) {
	if msg.Lider != s.lider {
		return nil, fmt.Errorf("no soy el líder")
	}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.store[msg.Clave] = msg.Valor
	return &ReplicaResponse{}, nil
}
