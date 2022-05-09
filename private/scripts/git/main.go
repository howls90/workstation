package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

var (
	STACKS = []string{"frontend", "backend"}
	CACHE_FILE = "private/scripts/git/.commits.toml"
)

type YamlConfig struct {
	Id string
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func getLastCommitId(fileName string) string {
	file, err := ioutil.ReadFile(fileName)
    if err != nil {
    	log.Fatal(err)
    }

	var cfg YamlConfig
	if err := toml.Unmarshal([]byte(file), &cfg); err != nil {
		log.Fatal(err)
	}
	return cfg.Id
}

var projects []string
func findProject(line string, stack string) {
	if strings.Contains(line, stack+"/apps") {
		projectName := strings.Split(line, "/")[2]
		if !contains(projects, stack+"/apps/"+projectName) {
			projects = append(projects, stack+"/apps/"+projectName)
			
			if _, err := os.Stat("./"+stack+"/apps/"+projectName); !os.IsNotExist(err) {
				// path/to/whatever exists
				fmt.Println(stack+"/apps/"+projectName)
			}
		}
	}
}

func main(){
	lastCommitId := getLastCommitId(CACHE_FILE)

	fileNameBytes, err := exec.Command("git", "diff", "--name-only", lastCommitId, "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}
	
    lines := strings.Split(string(fileNameBytes[:]), "\n")
	for _, line := range lines {
		for _, stack := range STACKS {
			findProject(line, stack)
		}
	}
}