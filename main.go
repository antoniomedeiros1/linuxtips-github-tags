package main

import (
	"github.com/antoniomedeiros1/linuxtips-github-tags/git"
)

func main() {

	var repoOwner string
	repoOwner = "kubernetes"
	repo := "kubernetes"
	tag := "v1.26.2"

	b := git.BuscadorGit{
		Repo:      repo,
		RepoOwner: repoOwner,
	}

	b.BuscaGitTag(tag)

}
