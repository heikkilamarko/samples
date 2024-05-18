package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	if err := generate(parseFlags()); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parseFlags() (int, string) {
	var n int
	var o string

	fs := flag.NewFlagSet("csvgen", flag.ExitOnError)

	fs.IntVar(&n, "n", 10, "row count")
	fs.StringVar(&o, "o", "data.csv", "output file")

	fs.Parse(os.Args[1:])

	return n, o
}

func generate(n int, o string) error {
	f, err := os.Create(o)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)

	defer w.Flush()

	w.Write([]string{
		"id",
		"name",
		"age",
		"height",
		"is_active",
		"created_at",
	})

	for i := 1; i <= n; i++ {
		w.Write([]string{
			id(i),
			name(i),
			age(i),
			height(),
			isActive(i),
			createdAt(),
		})
	}

	return nil
}

func id(i int) string {
	return fmt.Sprintf("%d", i)
}

func name(i int) string {
	return fmt.Sprintf("demo person %d", i)
}

func age(i int) string {
	return fmt.Sprintf("%d", i%100)
}

func height() string {
	return fmt.Sprintf("%f", 1.80)
}

func isActive(i int) string {
	return fmt.Sprintf("%t", i%2 == 0)
}

func createdAt() string {
	return time.Now().Format("2006-01-02T15:04:05Z")
}
