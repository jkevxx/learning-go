package domain

import (
	"fmt"
	"learning-go/domain/entities"
)

func App() {
	persona := entities.Persona{Nombre: "Kevin", Apellido: "Paez", Edad: 25}
	elemento := entities.TipoElemento{Nombre: "Learning"}

	fmt.Println(persona.Saludar())
	persona.CumplirAnios()
	fmt.Println(persona)
	fmt.Println(elemento)
}
