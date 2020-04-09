package main

import (
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/util/homedir"
)

var groupFlag string
var kubeconfig string

func setFlags() {
	flag.StringVar(&groupFlag, "group", "remote", "The group you wanna do something in")
	if home := homedir.HomeDir(); home != "" {
		flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
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
		groupCreate(args[1:])
		fmt.Printf("Created group %s\n", args[1])
	case "delete":
		groupDelete(args[1:])
		fmt.Printf("Deleted group %s\n", args[1])
	case "list":
		groupList()
	case "help":
		help()
	default:
		fmt.Printf("Oops! %s isn't a recognized command.\n", args[0])
	}
}

func workspace(args []string) {
	switch args[0] {
	case "create":
		workspaceCreate(args[1:])
		fmt.Printf("Created workspace %s\n", args[1])
	case "delete":
		workspaceDelete(args[1:])
		fmt.Printf("Deleted workspace %s\n", args[1])
	case "list":
		workspaceList()
	}
}

func help() {
	fmt.Printf("Help is on the way! But we haven't written this yet...\n")
}
