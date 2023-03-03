package git

import (
	"fmt"
	"net/http"
)

type Buscador interface {
	BuscaGitTag(tag string)
}

type BuscadorGit struct {
	Repo      string
	RepoOwner string
}

func (b *BuscadorGit) BuscaGitTag(tag string) {
	fmt.Println("Buscando commits pertencentes a git tag...")

	fmt.Printf("repo owner: %s \ntag: %s\n", b.RepoOwner, tag)

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s",
		b.RepoOwner, b.Repo)

	// go nao tem exception
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(nil)
	}
	if resp.StatusCode == 200 {
		fmt.Println(resp.Body)
	}

}
