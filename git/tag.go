package git

import "fmt"

func buscaGitTag(tag, repoOwner, repo string) {
	fmt.Println("Buscando commits pertencentes a git tag...")

	fmt.Printf("repo owner: %s \ntag: %s\n", repoOwner, tag)

}
