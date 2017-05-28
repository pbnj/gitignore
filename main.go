package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
)

var (
	debugFlag = flag.Bool("d", false, "Debug")
	langFlag  = flag.String("l", "", "Language")
	writeFlag = flag.Bool("w", false, "Write to .gitignore file")

	usage = fmt.Sprintf("USAGE: gitignore -l <language>")
)

func main() {

	flag.Parse()

	if flag.NFlag() == 0 {
		color.New(color.Bold).Println(usage)
		os.Exit(1)
	}

	if *debugFlag {
		log.SetLevel(log.DebugLevel)
	}

	langCapitalized := strings.Title(*langFlag) + ".gitignore"

	resBytes := getGitIgnore(langCapitalized)

	color.New(color.FgGreen, color.Bold).Println("✓ Found", langCapitalized)

	if *writeFlag {
		writeGitIgnore(resBytes)
	} else {
		fmt.Println(string(resBytes))
	}

}

func getGitIgnore(l string) []byte {
	gitURL := "https://raw.githubusercontent.com/github/gitignore/master/" + l

	log.Debugf("Checking for .gitignore at: %+v\n", gitURL)

	resp, err := http.Get(gitURL)
	checkErr(err)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		color.Red("✗ Could not find .gitignore for '%s'", l)
		return nil
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	checkErr(err)

	return bytes

}

func writeGitIgnore(b []byte) {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	checkErr(err)
	log.Debugf("Current Directory: %+v\n", currentDir)

	gitIgnoreFile := filepath.Join(currentDir, ".gitignore")
	log.Debugf(".gitignore file path to be created: %+v\n", gitIgnoreFile)

	err = ioutil.WriteFile(gitIgnoreFile, b, 0644)
	checkErr(err)

	fmt.Println("Created:", gitIgnoreFile)
}

func checkErr(err error) {
	if err != nil {
		color.Red("✗ %s", err)
		os.Exit(1)
	}
}
