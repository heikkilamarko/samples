package person

type Person struct {
	Name string
	Age  int
}

func New(name string, age int) Person {
	return Person{name, age}
}

func IsAdult(p Person) (bool, error) {
	return 18 < p.Age, nil
}

func GetName(p Person) (string, error) {
	return p.Name, nil
}

func AgeAsc(p []Person) func(i, j int) bool {
	return func(i, j int) bool {
		return p[i].Age < p[j].Age
	}
}

func AgeDesc(p []Person) func(i, j int) bool {
	return func(i, j int) bool {
		return p[j].Age < p[i].Age
	}
}

func StringAsc(s []string) func(i, j int) bool {
	return func(i, j int) bool {
		return s[i] < s[j]
	}
}
