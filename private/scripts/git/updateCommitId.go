package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/pelletier/go-toml/v2"
)

var (
	CACHE_FILE = "private/scripts/git/.commits.toml"
)

type YamlConfig struct {
	Id string
}

func readToml() YamlConfig{
	file, err := ioutil.ReadFile(CACHE_FILE)
    if err != nil {
    	log.Fatal(err)
    }

	var cfg YamlConfig
	if err := toml.Unmarshal([]byte(file), &cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}

func updateToml(cfg YamlConfig) {
	f, err := os.Create(CACHE_FILE)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

	bytes, err := toml.Marshal(cfg)
	if err != nil {
		panic(err)
	}
    _, err = f.WriteString(string(bytes))
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
	cfg := readToml()

	commitId, err := exec.Command("git", "rev-parse", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}
	
	cfg.Id = string(commitId[:len(commitId)-1])
	
	updateToml(cfg)
}