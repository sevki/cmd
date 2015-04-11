package main

import (
	"flag"
	"log"
	"os/exec"
	"strings"

	"github.com/nightlyone/go-stackdriver/stackdriver"
)

func main() {
	var revision string
	var author string
	if out, err := exec.Command("git",
		"describe",
		"--always").Output(); err != nil {
		revision = "deadbeef"
	} else {
		revision = strings.TrimSpace(string(out))
	}
	if out, err := exec.Command("git",
		"show",
		"--format=%aN",
		revision).Output(); err != nil {
		author = "deadbeef"
	} else {
		author = strings.TrimSpace(strings.Split(string(out), "\n")[0])
	}

	apikey := flag.String("apikey", "", "Apikey")
	apiendpoint := flag.String("apiendpoint", "https://gateway.google.stackdriver.com/v1", "Api Endpoint")
	commit := flag.String("commit", revision, "Revision")
	to := flag.String("to", "", "Deplyoed to")
	repo := flag.String("repo", "", "repo")
	flag.Parse()

	stackdriver.APIEndpoint = *apiendpoint
	stackdriver.APIKey = *apikey
	deploy := stackdriver.Deploy{
		RevisionID: *commit,
		DeployedBy: author,
	}
	if *to != "" {
		deploy.DeployedTo = *to
	}
	if *repo != "" {
		deploy.Repository = *repo
	}

	err := deploy.Submit()
	if err != nil {
		log.Fatal(err)
	}
}
