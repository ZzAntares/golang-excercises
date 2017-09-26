package main

import "fmt"

type persona struct {
	name string
	age  int
}

func (p persona) saluda(nombre string) string {
	return fmt.Sprintf("Hola %s!, soy %s y tengo %d años", nombre, p.name, p.age)
}

func main() {
	var messi = persona{"Leo Messi", 30}

	fmt.Println("Nombre:", messi.name)
	fmt.Println("Edad:", messi.age)
	fmt.Println("Saludo:", messi.saluda("pedro"))
}
