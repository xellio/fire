package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/xellio/fire"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 1 {
		showUsage(nil)
		return
	}

	jsonFile, err := os.Open(arguments[0])
	if err != nil {
		showUsage(err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		showUsage(errors.New("Problem reading the JSON file."))
		return
	}

	var requests []*fire.Request
	err = json.Unmarshal(byteValue, &requests)
	if err != nil {
		showUsage(errors.New("Problem parsing the JSON file."))
		return
	}

	var wg sync.WaitGroup
	for _, req := range requests {
		wg.Add(1)
		go func(req *fire.Request) {
			defer wg.Done()
			_, err := req.Fire()
			if err != nil {
				log.Println(err)
			}
		}(req)
	}
	wg.Wait()

	return

}

func showUsage(err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Usage: fire ./requests.json")
}
