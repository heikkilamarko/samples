package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/cap/ldap"
)

func main() {
	ctx := context.Background()

	urls := []string{"ldap://..."}
	domain := "..."
	username := "..."
	password := "..."

	config := &ldap.ClientConfig{
		URLs:           urls,
		UserDN:         "DC=...,DC=...",
		BindDN:         fmt.Sprintf("%s@%s", username, domain),
		BindPassword:   password,
		UserFilter:     fmt.Sprintf("(sAMAccountName=%s)", username),
		UseTokenGroups: true,
	}

	client, err := ldap.NewClient(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { client.Close(ctx) }()

	result, err := client.Authenticate(ctx, username, password, ldap.WithGroups())
	if err != nil {
		log.Fatal(err)
	}

	if result.Success {
		fmt.Println("--------------------------------")
		fmt.Println(username)
		fmt.Println("--------------------------------")
		if 0 < len(result.Groups) {
			for _, group := range result.Groups {
				fmt.Println(group)
			}
		}
		fmt.Println("--------------------------------")
	}

	fmt.Println(isMemberOf("...", result.Groups))
}

func isMemberOf(group string, groups []string) bool {
	for _, g := range groups {
		if g == group {
			return true
		}
	}
	return false
}
