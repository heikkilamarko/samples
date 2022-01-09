package person

type Person struct {
	Name string
	Age  int
}

func New(name string, age int) Person {
	return Person{name, age}
}

func (p *Person) IsAdult() bool {
	return 18 <= p.Age
}
