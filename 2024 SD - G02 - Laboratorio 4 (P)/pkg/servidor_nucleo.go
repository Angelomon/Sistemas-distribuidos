package pkg

import (
    "context"
    "fmt"
    "sync"
)

// La implementación del servidor
type Servidor struct {
    UnimplementedBaseServer
    store map[string][]byte
    mutex sync.RWMutex
}

// NuevoServidor retorna una instancia del servidor
func NuevoServidor() Servidor {
    return Servidor{
        store: make(map[string][]byte),
    }
}

// Implementación de Put definido en el archivo .proto.
func (s *Servidor) Put(ctx context.Context, msg *ParametroPut) (*ResultadoPut, error) {
    s.mutex.Lock()
    defer s.mutex.Unlock()
    s.store[msg.Clave] = msg.Valor
    return &ResultadoPut{}, nil
}

// Implementación de Get definido en el archivo .proto.
func (s *Servidor) Get(ctx context.Context, msg *ParametroGet) (*ResultadoGet, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    valor, exists := s.store[msg.Clave]
    if !exists {
        return nil, fmt.Errorf("clave no encontrada")
    }
    return &ResultadoGet{Valor: valor}, nil
}

// Implementación de GetAll definido en el archivo .proto.
func (s *Servidor) GetAll(ctx context.Context, msg *ParametroGetAll) (*ResultadoGetAll, error) {
    s.mutex.RLock()
    defer s.mutex.RUnlock()
    result := &ResultadoGetAll{}
    for clave, valor := range s.store {
        result.ClaveValor = append(result.ClaveValor, &ClaveValor{Clave: clave, Valor: valor})
    }
    return result, nil
}