package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	az, err := NewAuthZ(ctx)
	checkErr(err)

	// sample input as struct
	run(ctx, az, AuthZInput{
		Action: "read",
		Object: "server123",
		Token: map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{"hr"},
				},
			},
		},
	})

	// sample input as map
	run(ctx, az, map[string]any{
		"action": "read",
		"object": "server123",
		"token": map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{"hr", "engineering"},
				},
			},
		},
	})
}

func run(ctx context.Context, az *AuthZ, input any) {
	allow, err := az.Authorize(ctx, input)
	checkErr(err)

	log.Printf("allow: %t\n", allow)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
