package main

import (
	"io/ioutil"
	"os/exec"
	"strings"
)

func createWorkspace(name, group string) {
	dat, err := ioutil.ReadFile("kubernetes/pod.yaml")
	if err != nil {
		panic(err)
	}
	txt := []byte(strings.ReplaceAll(string(dat), "workstation", name))
	err = ioutil.WriteFile(".remote/"+group+"/"+name, txt, 0777)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("kubectl", "apply", "-n", group, "-f", ".remote/"+group+"/"+name)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}

func deleteWorkspace(name, group string) {
	cmd := exec.Command("kubectl", "delete", "-n", group, "-f", ".remote/"+group+"/"+name)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func createNamespace(name string) {
	cmd := exec.Command("kubectl", "create", "ns", name)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func deleteNamespace(name string) {
	cmd := exec.Command("kubectl", "delete", "ns", name)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
