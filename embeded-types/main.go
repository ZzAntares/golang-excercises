package main

import "fmt"

type Persona struct {
	name string
	age  int
}

func (p *Persona) Greeting() string {
	return fmt.Sprintf("Hello! I'm %s and I'm %d years old.", p.name, p.age)
}

type Futbolista struct {
	Persona
	team string
}

func (f *Futbolista) Greeting() string {
	return fmt.Sprintf("Soy %s del %s con %d añitos.", f.name, f.team, f.age)
}

func main() {
	var messi Futbolista = Futbolista{
		team: "Barcelona",
	}

	messi.name = "Lionel Messi"
	messi.age = 33

	fmt.Println(messi.Greeting())
}
