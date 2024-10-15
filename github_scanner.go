// github_scanner.go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "os"

    "github.com/google/go-github/v41/github"
    "golang.org/x/oauth2"
)

func main() {
    ctx := context.Background()

    // Set up authentication
    ts := oauth2.StaticTokenSource(
        &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
    )
    tc := oauth2.NewClient(ctx, ts)

    client := github.NewClient(tc)

    // Specify the owner and repo
    owner := "owner_name" // Replace with actual owner
    repo := "repo_name"   // Replace with actual repo

    contents, _, _, err := client.Repositories.GetContents(ctx, owner, repo, "", nil)
    if err != nil {
        log.Fatalf("Error getting repository contents: %v", err)
    }

    scanFiles(contents)
}

func scanFiles(contents []github.RepositoryContent) {
    for _, content := range contents {
        if *content.Type == "file" {
            if endsWith(*content.Name, ".go") {
                fmt.Printf("Found Go file: %s\n", *content.Name)
            }
        }
    }
}

func endsWith(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}
