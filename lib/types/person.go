package types

type Person struct {
	id   int
	name string
}

func (p *Person) GetId() int {
	return p.id
}

func (p *Person) GetName() string {
	return p.name
}
