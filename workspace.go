package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func workspaceCreate(args []string) {
	f, err := os.Create(".remote/" + groupFlag + "/" + args[0])
	f.Close()
	if err != nil {
		log.Fatal(err)
	}
	createWorkspace(args[0], groupFlag)
}

func workspaceDelete(args []string) {
	deleteWorkspace(args[0], groupFlag)
	err := os.Remove(".remote/" + groupFlag + "/" + args[0])
	if err != nil {
		log.Fatal(err)
	}
}

func workspaceList() {
	files, err := ioutil.ReadDir(".remote/" + groupFlag)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		fmt.Println("No workspaces found")
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
