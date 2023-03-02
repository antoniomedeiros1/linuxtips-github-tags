package main

import (
	"github.com/antoniomedeiros1/linuxtips-github-tags/git"
)

func main() {

	var repoOwner string
	repoOwner = "kubernetes"
	repo := "kubernetes"
	tag := "v1.26.2"

	git.BuscaGitTag(tag, repoOwner, repo)
}
