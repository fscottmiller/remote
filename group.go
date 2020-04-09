package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func groupCreate(args []string) {
	err := os.MkdirAll(".remote/"+args[0], 0777)
	if err != nil {
		log.Fatal(err)
	}
	createNamespace(args[0])
}

func groupDelete(args []string) {
	err := os.Remove(".remote/" + args[0])
	if err != nil {
		log.Fatal(err)
	}
	deleteNamespace(args[0])
}

func groupList() {
	files, err := ioutil.ReadDir(".remote")
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		fmt.Println("No groups found")
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
