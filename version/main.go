package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	majorVersion = "3"
)

type Schema struct {
	Info struct {
		Version string `yaml:"version"`
	} `yaml:"info"`
}

func getUpstreamVersion() string {
	var schema Schema
	yamlFile, err := ioutil.ReadFile("../schema.yml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &schema)
	if err != nil {
		panic(err)
	}
	v := schema.Info.Version
	parts := strings.Split(v, ".")
	m, _ := strconv.Atoi(parts[1])
	parts[1] = fmt.Sprintf("%02d", m)
	fmt.Sprintln(parts)
	return strings.Join(parts, ".")
}

func getLatestTag() string {
	out, err := exec.Command("git", "describe", "--abbrev=0", "--tags").Output()
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(out))
}

func main() {
	// Make sure there are no suffixes, and remove all dots so we have the same as above
	upstream := strings.ReplaceAll(strings.Split(getUpstreamVersion(), "-")[0], ".", "")
	tag := strings.Split(getLatestTag(), ".")
	rev, err := strconv.Atoi(tag[2])
	if err != nil {
		panic(err)
	}
	if upstream != tag[1] {
		rev = 0
	}
	newVersion := fmt.Sprintf("v%s.%s.%d", majorVersion, upstream, rev+1)
	fmt.Print(newVersion)
}
