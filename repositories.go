package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/viper"
)

type Owner struct {
	Login string `json:"login,omitempty"`
}

type Repository struct {
	Owner    Owner  `json:"owner,omitempty"`
	Name     string `json:"name,omitempty"`
	FullName string `json:"full_name,omitempty"`
	Language string `json:"language,omitempty"`
	Fork     bool   `json:"fork,omitempty"`
}

func (r Repository) pull(p string) {
	repo, err := git.PlainOpen(p)

	CheckError(err)

	w, err := repo.Worktree()

	CheckError(err)

	fmt.Printf("Local Repo Exists... pulling %v \n", r.FullName)
	w.Pull(&git.PullOptions{RemoteName: "origin"})

}

func checkKeyWords(name string, lang string) string {
	keyWords := viper.GetStringMapString("keywords")

	for key, val := range keyWords {
		if strings.Contains(strings.ToLower(name), strings.ToLower(key)) {
			return val
		}
	}
	return lang
}

func (r Repository) clone() {
	destPath := viper.GetString("destination.local")
	if viper.GetBool("sortByLanguage") {
		if r.Language == "" {
			destPath = destPath + "/MiscLang/"
		} else {
			destPath = destPath + "/" + checkKeyWords(r.Name, r.Language) + "/"
		}
	}

	absPath, e := filepath.Abs(destPath + r.Name)

	CheckError(e)

	_, err := git.PlainClone(absPath, false, &git.CloneOptions{
		URL:      "https://github.com/" + r.FullName,
		Progress: os.Stdout,
		Auth: &http.BasicAuth{
			Username: viper.GetString("source.github.user"),
			Password: viper.GetString("source.github.token"),
		},
	})

	if err != nil {
		if err.Error() == "repository already exists" {
			if viper.GetBool("pull") {
				r.pull(absPath)
			} else {
				fmt.Printf("Pulling Disabled, skipping repo %v \n", r.Name)
			}
		} else {
			fmt.Println(err)
		}
	}

}

type Repositories []Repository

func CheckError(e error) {
	if e != nil {
		fmt.Println("Pull Error", e)
	}
}
