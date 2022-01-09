package main

import (
	"fmt"
	g "generics/internal/generics"
	"generics/internal/person"
)

func main() {

	filteredInts, _ := g.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		g.Or(
			g.Equal(9),
			g.Between(4, 6),
			g.LessThan(3),
		),
	)
	fmt.Println(filteredInts)
	// [1 2 4 5 6 9]

	filteredStrings, _ := g.Filter([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "jj", "kk"},
		g.Or(
			g.Equal("e"),
			g.LessThan("c"),
			g.And(
				g.GreaterThan("h"),
				g.StringLength(2),
				g.NotEqual("kk"),
			),
		),
	)
	fmt.Println(filteredStrings)
	// [a b e jj]

	people := []person.Person{{Name: "b", Age: 30}, {Name: "a", Age: 10}, {Name: "c", Age: 40}}
	fmt.Println(people)
	// [{b 30} {a 10} {c 40}]

	people, _ = g.OrderBy(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println(people)
	// [{a 10} {b 30} {c 40}]

	people, _ = g.OrderBy(people, func(i, j int) bool { return people[j].Age < people[i].Age })
	fmt.Println(people)
	// [{c 40} {b 30} {a 10}]

	people, _ = g.Filter(people, func(p person.Person) (bool, error) { return p.IsAdult(), nil })
	fmt.Println(people)
	// [{c 40} {b 30}]

	names, _ := g.Map(people, func(p person.Person) (string, error) { return p.Name, nil })
	fmt.Println(names)
	// [c b]

	names, _ = g.OrderBy(names, func(i, j int) bool { return names[i] < names[j] })
	fmt.Println(names)
	// [b c]

	names, _ = g.Filter(names, g.NotEqual("b"))
	fmt.Println(names)
	// [c]

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(numbers)
	// [1 2 3 4 5 6 7 8 9 10]

	_, err := g.Page(numbers, -1, 0)
	fmt.Println(err)
	// invalid offset

	_, err = g.Page(numbers, 0, 0)
	fmt.Println(err)
	// invalid limit

	for offset, limit := 0, 3; ; offset += limit {
		page, err := g.Page(numbers, offset, limit)
		if err != nil {
			fmt.Println("ERROR:", err)
			break
		}
		if len(page) == 0 {
			break
		}
		fmt.Println(page)
	}
	// [1 2 3]
	// [4 5 6]
	// [7 8 9]
	// [10]
}
