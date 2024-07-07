package pkg

import (
	"context"
	"fmt"
	"hash/fnv"
	"sync"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type Servidor struct {
	UnimplementedBaseServer
	store map[string][]byte
	mutex sync.RWMutex
	nodos []string
}

func NuevoServidor(nodos []string) Servidor {
	return Servidor{
		store: make(map[string][]byte),
		nodos: nodos,
	}
}

func (s *Servidor) hash(clave string) int {
	h := fnv.New32a()
	h.Write([]byte(clave))
	return int(h.Sum32()) % len(s.nodos)
}

func (s *Servidor) getNodo(clave string) string {
	return s.nodos[s.hash(clave)]
}

func (s *Servidor) Put(ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
	nodo := s.getNodo(msg.Clave)
	log.Printf("Put: clave %s -> nodo %s", msg.Clave, nodo)
	if nodo == "localhost:12345" { // Verificar si el nodo actual es el nodo correcto
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.store[msg.Clave] = msg.Valor
		log.Printf("Put: clave %s almacenada", msg.Clave)
		s.replicar(msg) // Replicar a nodos seguidores
		return &ResultadoPut{}, nil
	} else {
		if nodo == "localhost:12346" { // Verificar si el nodo actual es el nodo correcto
			s.mutex.Lock()
			defer s.mutex.Unlock()
			s.store[msg.Clave] = msg.Valor
			log.Printf("Put: clave %s almacenada", msg.Clave)
			s.replicar(msg) // Replicar a nodos seguidores
			return &ResultadoPut{}, nil
		} else {
			if nodo == "localhost:12347" { // Verificar si el nodo actual es el nodo correcto
				s.mutex.Lock()
				defer s.mutex.Unlock()
				s.store[msg.Clave] = msg.Valor
				log.Printf("Put: clave %s almacenada", msg.Clave)
				s.replicar(msg) // Replicar a nodos seguidores
				return &ResultadoPut{}, nil
			}
			return s.redirigirPut(ctx, nodo, msg) // Redirigir si no es el nodo correcto
		}
	}

}

func (s *Servidor) redirigirPut(ctx context.Context, nodo string, msg *ParametroPut) (*ResultadoPut, error) {
	conn, err := grpc.Dial(nodo, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al nodo %s: %v", nodo, err)
	}
	defer conn.Close()

	cliente := NewBaseClient(conn)
	return cliente.Put(ctx, msg)
}

func (s *Servidor) replicar(msg *ParametroPut) {
	for _, nodo := range s.nodos {
		if nodo != "localhost:12345" {
			conn, err := grpc.Dial(nodo, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Printf("No se pudo conectar al nodo %s para replicar: %v", nodo, err)
				continue
			}
			defer conn.Close()

			cliente := NewBaseClient(conn)
			_, err = cliente.Put(context.Background(), msg)
			if err != nil {
				log.Printf("Error al replicar en nodo %s: %v\n", nodo, err)
			}
		}
	}
}

func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
	nodo := s.getNodo(msg.Clave)
	log.Printf("Get: clave %s -> nodo %s", msg.Clave, nodo)
	if nodo == "localhost:12345" {
		s.mutex.RLock()
		defer s.mutex.RUnlock()
		valor, exists := s.store[msg.Clave]
		if !exists {
			return nil, status.Error(404, "clave no encontrada")
		}
		return &ResultadoGet{Valor: valor}, nil
	} else {
		if nodo == "localhost:12346" {
			s.mutex.RLock()
			defer s.mutex.RUnlock()
			valor, exists := s.store[msg.Clave]
			if !exists {
				return nil, status.Error(404, "clave no encontrada")
			}
			return &ResultadoGet{Valor: valor}, nil
		} else {
			if nodo == "localhost:12347" {
				s.mutex.RLock()
				defer s.mutex.RUnlock()
				valor, exists := s.store[msg.Clave]
				if !exists {
					return nil, status.Error(404, "clave no encontrada")
				}
				return &ResultadoGet{Valor: valor}, nil
			}
		}
	}

	return s.redirigirGet(ctx, nodo, msg)
}

func (s *Servidor) redirigirGet(ctx context.Context, nodo string, msg *ParametroGet) (*ResultadoGet, error) {
	conn, err := grpc.Dial(nodo, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("no se pudo conectar al nodo %s: %v", nodo, err)
	}
	defer conn.Close()

	cliente := NewBaseClient(conn)
	return cliente.Get(ctx, msg)
}

func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	result := &ResultadoGetAll{}
	for clave, valor := range s.store {
		result.ClaveValor = append(result.ClaveValor, &ClaveValor{Clave: clave, Valor: valor})
	}
	return result, nil
}
