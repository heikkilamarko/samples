package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.Background()

	az, err := NewAuthZ(ctx)
	checkErr(err)

	// allow: false
	run(ctx, az, nil)

	// allow: false
	run(ctx, az, AuthZInput{})

	// allow: false
	run(ctx, az, AuthZInput{
		Permission: "",
		Token:      nil,
	})

	// allow: false
	run(ctx, az, AuthZInput{
		Permission: "sample.write",
		Token:      nil,
	})

	// allow: false
	run(ctx, az, AuthZInput{
		Permission: "sample.write",
		Token: map[string]any{
			"resource_access": map[string]any{
				"wrong-api": map[string]any{
					"roles": []string{},
				},
			},
		},
	})

	// allow: false
	run(ctx, az, AuthZInput{
		Permission: "sample.write",
		Token: map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{},
				},
			},
		},
	})

	// allow: false
	run(ctx, az, AuthZInput{
		Permission: "sample.write",
		Token: map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{"sample.reader"},
				},
			},
		},
	})

	// allow: true
	run(ctx, az, AuthZInput{
		Permission: "sample.write",
		Token: map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{"sample.writer"},
				},
			},
		},
	})

	// allow: true
	run(ctx, az, map[string]any{
		"permission": "sample.write",
		"token": map[string]any{
			"resource_access": map[string]any{
				"sample-api": map[string]any{
					"roles": []string{"sample.reader", "sample.admin"},
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
