package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bradleyfalzon/ghinstallation"
	"github.com/google/go-github/v32/github"
	"net/http"
	"strings"
)

func main() {
	var privateKeyFile string
	var appID int64
	var baseURL string
	var repoArg string

	flag.StringVar(&privateKeyFile, "private-key-file", "", "Path to file containing GitHub app private key PEM")
	flag.Int64Var(&appID, "app-id", 0, "GitHub Application ID")
	flag.StringVar(&repoArg, "repo", "", "GitHub repository (owner/repo)")
	flag.StringVar(&baseURL, "base-url", "", "GitHub API base URL (e.g. https://github.example.com/api/v3/)")
	flag.Parse()

	if privateKeyFile == "" {
		panic("private-key-file is required")
	}
	if appID == 0 {
		panic("app-id is required")
	}
	if repoArg == "" {
		panic("repo is required")
	}
	repo := strings.SplitN(repoArg, "/", 2)
	if len(repo) != 2 {
		panic("invalid repo value")
	}

	// Prepare transport
	tr, err := ghinstallation.NewAppsTransportKeyFromFile(http.DefaultTransport, appID, privateKeyFile)
	if err != nil {
		panic(err)
	}
	tr.BaseURL = baseURL

	// New client
	client := github.NewClient(&http.Client{Transport: tr})

	// Get Installation ID
	inst, _, err := client.Apps.FindRepositoryInstallation(context.Background(), repo[0], repo[1])
	if err != nil {
		panic(err)
	}

	// Get Installation Token
	token, _, err := client.Apps.CreateInstallationToken(context.Background(), *inst.ID, nil)
	if err != nil {
		panic(err)
	}

	println("Expires:", token.ExpiresAt.String()) // to stderr
	fmt.Println(*token.Token)                     // to stdout
}
