package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

const BACKEND_PATHS = "./backend/apps"

func main() {
    files, err := ioutil.ReadDir(BACKEND_PATHS)
    if err != nil {
        log.Fatal(err)
    }
 
    res := "docker-compose"
    for _, f := range files {
        res +=  " -f backend/apps/"+f.Name()+"/docker-compose.yml"
    }
    res += " up"

    fmt.Println(res)
}