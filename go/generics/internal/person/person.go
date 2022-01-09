package person

type Person struct {
	Name string
	Age  int
}

func (p *Person) IsAdult() bool {
	return 18 <= p.Age
}
