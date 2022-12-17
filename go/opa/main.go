package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	az, err := NewAuthZ(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// input as struct

	bob := AuthZInput{
		Action: "read",
		Object: "database456",
		User:   "bob",
	}

	ok, err := az.Authorize(ctx, bob)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("allow bob: %t\n", ok)

	// input as map

	alice := map[string]string{
		"action": "read",
		"object": "database456",
		"user":   "alice",
	}

	ok, err = az.Authorize(ctx, alice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("allow alice: %t\n", ok)
}
