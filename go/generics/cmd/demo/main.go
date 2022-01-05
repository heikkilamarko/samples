package main

import (
	"fmt"
	g "generics/internal/generics"
)

type dude struct {
	Name string
	Age  int
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

		g.IsEqual(dude{"a", 1}, dude{"a", 1}),
		// true

		g.IsNotEqual(dude{"a", 1}, dude{"a", 1}),
		// false

		g.Filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			g.Or(
				g.Equal(9),
				g.Between(4, 6),
				g.LessThan(3),
			),
		),
		// [1 2 4 5 6 9]

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
		// [a b e jj]

		g.Filter([]dude{{"a", 10}, {"b", 30}, {"c", 40}},
			func(d dude) bool { return 18 < d.Age },
		),
		// [{b 30} {c 40}]
	)
}
