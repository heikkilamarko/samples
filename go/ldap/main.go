package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/cap/ldap"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("ldap.env")
	checkErr(err)

	username, password := parseFlags()

	config := &ldap.ClientConfig{
		URLs:           []string{os.Getenv("LDAP_URL")},
		BindDN:         getBindDN(username, os.Getenv("LDAP_DOMAIN")),
		BindPassword:   password,
		UserDN:         os.Getenv("LDAP_USER_DN"),
		GroupDN:        os.Getenv("LDAP_GROUP_DN"),
		UserFilter:     os.Getenv("LDAP_USER_FILTER"),
		UseTokenGroups: os.Getenv("LDAP_USE_TOKEN_GROUPS") == "true",
	}

	ctx := context.Background()

	client, err := ldap.NewClient(ctx, config)
	checkErr(err)

	defer func() { client.Close(ctx) }()

	result, err := client.Authenticate(ctx, username, password, ldap.WithGroups())
	checkErr(err)

	printUsername(username)

	if result.Success {
		printGroups(result.Groups)
	}
}

func parseFlags() (username, password string) {
	fs := flag.NewFlagSet("ldap-sample", flag.ExitOnError)

	fs.StringVar(&username, "u", "", "username")
	fs.StringVar(&password, "p", "", "password")

	fs.Parse(os.Args[1:])

	if username == "" || password == "" {
		fs.Usage()
		os.Exit(2)
	}

	return
}

func getBindDN(username, domain string) string {
	username = ldap.EscapeValue(username)
	if _, _, found := strings.Cut(username, "@"); found {
		return username
	}
	return fmt.Sprintf("%s@%s", username, domain)
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

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
