package models

type Persona struct {
	nombre string
	edad   int
}

//Constructor de la struct Persona
//Encapsulacion
func (p *Persona) Constructor(nombre string, edad int) {
	p.nombre = nombre
	p.edad = edad
}

//Metodos Getter y Setter

func (p *Persona) GetNombre() string {
	return p.nombre
}
func (p *Persona) SetNombre(nuevoNombre string) {
	p.nombre = nuevoNombre
}

func (p *Persona) GetEdad() int {
	return p.edad
}
func (p *Persona) SetEdad(nuevaEdad int) {
	p.edad = nuevaEdad
}
