package main

import "fmt"

type Payable interface {
	salary() float64
}

type Persona struct {
	Name string
	Age  int
}

func (p *Persona) Greeting() string {
	return fmt.Sprintf("Hello, I'm both, %s and %d years old", p.Name, p.Age)
}

type Ingeniero struct {
	Persona
	Startup string
}

func (i *Ingeniero) salary() float64 {
	return float64(i.Age*10 + 100)
}

type Futbolista struct {
	Persona
	Team string
}

func (f *Futbolista) salary() float64 {
	return float64(f.Age * 100)
}

func main() {
	var messi Futbolista = Futbolista{Team: "Barcelona"}
	messi.Name = "Lionel Messi"
	messi.Age = 33

	var peter Ingeniero = Ingeniero{
		Startup: "DigitalOcean",
		Persona: Persona{Name: "Peter Norvig", Age: 26},
	}

	fmt.Println(messi.Greeting())
	fmt.Println(peter.Greeting())
	fmt.Printf("%s earns %.2f\n", messi.Name, messi.salary())
	fmt.Printf("%s earns %.2f\n", peter.Name, peter.salary())
}
