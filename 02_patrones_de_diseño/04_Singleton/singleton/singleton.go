package singleton

import (
	"fmt"
	"sync"
)

// mantengo el struct como privado. Para instanciarlo voy a tener q recurrir a un método
type single struct {
	instancia int
}

func (s single) VerIntancia() string {
	return fmt.Sprintf("La instancia es: %d", s.instancia)
}

// aca voy a guardar la instancia generada
var singleInstance *single

// el sync.Once... Do() hace que la función solo se ejecute una vez
var once sync.Once

// el id es para identificar la instancia generada y ver q siempre devuelvo la misma
func GetInstancia(id int) *single {

	// al hace q la funcion que instancea al struct ocurra solo una vez, siempre devuelvo el mismo struct
	once.Do(func() {
		fmt.Println("Generando Instancia")
		singleInstance = &single{
			instancia: id,
		}
	})

	return singleInstance
}
