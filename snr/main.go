package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

var (
	oldFile = flag.String("of", "", "The text that is going to be replaced")
	newFile = flag.String("nf", "", "The text that is going to replace")
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 0 {
		fmt.Println("You need to specify a glob pattern")
	}
	if *oldFile == "" || *newFile == "" {
		flag.PrintDefaults()
	}
	old, err := ioutil.ReadFile(*oldFile)
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	new, err := ioutil.ReadFile(*newFile)
	if err != nil {
		log.Fatalf("read: %s", err)
	}
	for _, arg := range args {
		matches, err := filepath.Glob(arg)
		if err != nil {
			log.Fatalf("glob: %s", err)
		}
		for _, file := range matches {
			log.Println(file)

			bytz, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("read: %s", err)
			}
			nb := bytes.Replace(bytz, old, new, -1)
			log.Printf("%d %d", len(bytz), len(nb))

			err = ioutil.WriteFile(file, nb, 0644)
			if err != nil {
				log.Fatalf("write: %s", err)
			}

		}
	}

}
