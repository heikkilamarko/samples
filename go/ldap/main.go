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
		BindDN:         fmt.Sprintf("%s@%s", username, domain),
		BindPassword:   password,
		UserDN:         "DC=...,DC=...",
		GroupDN:        "DC=...,DC=...",
		UserFilter:     fmt.Sprintf("(sAMAccountName=%s)", username),
		UseTokenGroups: false,
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

	printUsername(username)

	if result.Success {
		printGroups(result.Groups)
	}
}

func printUsername(username string) {
	fmt.Println("[username]")
	fmt.Printf("\t%s\n", username)
}

func printGroups(groups []string) {
	fmt.Printf("[groups (%d)]\n", len(groups))
	for i, group := range groups {
		fmt.Printf("\t%3d. %s\n", i+1, group)
	}
}
