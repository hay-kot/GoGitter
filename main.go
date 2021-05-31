package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func setDefaults() {
	viper.SetDefault("sorByLanguage", true)
	viper.SetDefault("zipAll", false)
	viper.SetDefault("forks", false)
	viper.SetDefault("pull", true)
}

func getConfig() {
	setDefaults()

	confPath := os.Args[1]
	confDir := filepath.Dir(confPath)

	viper.SetConfigName("GoGitter")
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(confDir)
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	getConfig()

	bearer := "Bearer " + viper.GetString("source.github.token")
	req_url := "https://api.github.com/user/repos"

	req, _ := http.NewRequest("GET", req_url, nil)
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	response, err := client.Do(req)

	CheckError(err)

	responseData, err := ioutil.ReadAll(response.Body)

	var repos Repositories

	json.Unmarshal(responseData, &repos)

	CheckError(err)

	for _, repo := range repos {
		if repo.Fork && !viper.GetBool("forks") {
			continue
		}
		repo.clone()
	}

}
