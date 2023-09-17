package main

import (
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

func doNew(appName string) {
	appName = strings.ToLower(appName)

	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[(len(exploded) - 1)]
	}

	color.Green("\tCloning repository...")

	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		URL:      "git@github.com:IrakliGiorgadze/celeritas-app.git",
		Depth:    1,
		Progress: os.Stdout,
	})
	if err != nil {
		exitGracefully(err)
	}
}
