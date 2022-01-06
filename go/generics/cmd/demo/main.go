package main

import (
	"fmt"
	g "generics/internal/generics"
)

type person struct {
	Name string
	Age  int
}

func isAdult(p person) (bool, error) {
	return 18 < p.Age, nil
}

func getName(p person) (string, error) {
	return p.Name, nil
}

func ageAsc(p []person) func(i, j int) bool {
	return func(i, j int) bool {
		return p[i].Age < p[j].Age
	}
}

func ageDesc(p []person) func(i, j int) bool {
	return func(i, j int) bool {
		return p[j].Age < p[i].Age
	}
}

func stringAsc(s []string) func(i, j int) bool {
	return func(i, j int) bool {
		return s[i] < s[j]
	}
}

func main() {
	fmt.Println(
		g.IsEqual(1, 1),
		// true

		g.IsNotEqual(1, 1),
		// false

		g.IsEqual("1", "2"),
		// false

		g.IsNotEqual("1", "2"),
		// true

		g.IsEqual(person{"a", 1}, person{"a", 1}),
		// true

		g.IsNotEqual(person{"a", 1}, person{"a", 1}),
		// false
	)

	fmt.Println(
		g.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			g.Or(
				g.Equal(9),
				g.Between(4, 6),
				g.LessThan(3),
			),
		),
	)
	// [1 2 4 5 6 9] <nil>

	fmt.Println(
		g.Filter([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "jj", "kk"},
			g.Or(
				g.Equal("e"),
				g.LessThan("c"),
				g.And(
					g.GreaterThan("h"),
					g.StringLength(2),
					g.NotEqual("kk"),
				),
			),
		),
	)
	// [a b e jj] <nil>

	people := []person{{"b", 30}, {"a", 10}, {"c", 40}}
	fmt.Println(people)
	// [{b 30} {a 10} {c 40}]

	people, _ = g.OrderBy(people, ageAsc(people))
	fmt.Println(people)
	// [{a 10} {b 30} {c 40}]

	people, _ = g.OrderBy(people, ageDesc(people))
	fmt.Println(people)
	// [{c 40} {b 30} {a 10}]

	people, _ = g.Filter(people, isAdult)
	fmt.Println(people)
	// [{c 40} {b 30}]

	names, _ := g.Map(people, getName)
	fmt.Println(names)
	// [c b]

	names, _ = g.OrderBy(names, stringAsc(names))
	fmt.Println(names)
	// [b c]

	names, _ = g.Filter(names, g.NotEqual("b"))
	fmt.Println(names)
	// [c]
}
