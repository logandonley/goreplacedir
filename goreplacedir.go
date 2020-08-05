package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func getEnv() map[string]string {
	m := make(map[string]string)
	for _, item := range os.Environ() {
		split := strings.Split(item, "=")
		m[split[0]] = split[1]
	}
	return m
}

func parse(path string, dest string) {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Println(err)
		return
	}
	f, err := os.Create(dest)
	if err != nil {
		log.Println(err)
		return
	}
	config := getEnv()

	err = t.Execute(f, config)
	if err != nil {
		log.Println(err)
		return
	}
	_ = f.Close()
}


func main() {
	argLength := len(os.Args[1:])
	if argLength != 2 {
		fmt.Println("Requires 2 arguments. Syntax should look like:")
		fmt.Println("./goreplacedir <target dir> <output dir>")
		os.Exit(1)
	}
	target := os.Args[1]
	dest := os.Args[2]
	targetSlice := strings.Split(target, "/")
	targetLen := len(targetSlice)
	destSlice := strings.Split(dest, "/")
	err := filepath.Walk(target, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() == false {
			pathSlice := strings.Split(path, "/")
			combinedSlice := append(destSlice, pathSlice[targetLen:]...)
			err = os.MkdirAll(strings.Join(combinedSlice[:len(combinedSlice)-1], "/"), os.ModePerm)
			if err != nil {
				log.Fatalln("Issue with creating directories: ", err)
			}
			destPath := strings.Join(combinedSlice, "/")
			fmt.Println(path)
			fmt.Println(destPath)
			parse(path, destPath)
		}
		return nil
	})
	if err != nil {
		log.Fatalln(err)
	}
}
