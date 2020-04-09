package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

var groupFlag string
var kubeconfig *string

func setFlags() {
	flag.StringVar(&groupFlag, "group", "remote", "The group you wanna do something in")
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
}

func main() {
	setFlags()

	switch flag.Args()[0] {
	case "group":
		group(flag.Args()[1:])
	case "workspace":
		workspace(flag.Args()[1:])
	case "help":
		help()
	default:
		fmt.Printf("Oops! %s isn't a recognized command.\n", flag.Args()[0])
	}
}

func group(args []string) {
	switch args[0] {
	case "create":
		err := os.MkdirAll(".remote/"+args[1], 0777)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created group %s\n", args[1])
	case "delete":
		err := os.Remove(".remote/" + args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Deleted group %s\n", args[1])
	case "list":
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
	case "help":
		help()
	default:
		fmt.Printf("Oops! %s isn't a recognized command.\n", args[0])
	}
}

func workspace(args []string) {
	switch args[0] {
	case "create":
		f, err := os.Create(".remote/" + groupFlag + "/" + args[1])
		f.Close()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created workspace %s\n", args[1])
	case "delete":
		err := os.Remove(".remote/" + groupFlag + "/" + args[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Created workspace %s\n", args[1])
	case "list":
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
}

func help() {
	fmt.Printf("Help is on the way! But we haven't written this yet...\n")
}
