package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

var (
	lang          string
	langGitIgnore string
	found         bool
)

func main() {
	usage := fmt.Sprintf("USAGE: gitignore <language>")
	if len(os.Args) == 1 {
		fmt.Println(usage)
		os.Exit(1)
	}

	if os.Args[1] == "--help" {
		fmt.Println(usage)
		os.Exit(1)
	}

	lang = os.Args[1]
	langGitIgnore = strings.Title(lang) + ".gitignore"

	assets := AssetNames()
	for _, name := range assets {
		checkGitIgnore(name)
	}
	if !found {
		color.Red("✗ Could not find .gitignore for '%s'", lang)
	}
}

func checkGitIgnore(name string) {
	pathLangGitIgnore := fmt.Sprintf("gitignore/%s", langGitIgnore)
	if name == pathLangGitIgnore {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		checkErr(err)
		bytes, err := Asset(name)
		checkErr(err)
		ioutil.WriteFile(".gitignore", bytes, 0644)
		color.Green("✓ Created .gitignore for %s in %s", lang, dir)
		found = true
	}
}

func checkErr(err error) {
	if err != nil {
		color.Red("✗ %s", err)
		os.Exit(1)
	}
}
