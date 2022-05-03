package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Root struct {
	Root  []File `json:"root"`
	count int
}

type File struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Files []File `json:"files"`
}

func (root *Root) createFiletree(path string) ([]File, error) {
	root.count++
	fileList, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	elements := make([]File, 0)
	for _, f := range fileList {
		e := File{}
		e.Name = f.Name()
		if f.IsDir() {
			e.Type = "folder"
			files, err := root.createFiletree(filepath.Join(path, e.Name))
			if err != nil {
				return nil, err
			}
			e.Files = files
		}
		if !f.IsDir() {
			e.Type = "file"
			e.Files = make([]File, 0)
		}
		elements = append(elements, e)
	}
	return elements, nil
}

func main() {
	var path string
	var pretty bool
	flag.BoolVar(&pretty, "pretty", false, "print pretty")
	flag.StringVar(&path, "path", ".", "path to get filetree from")
	flag.Parse()

	root := new(Root)
	root.Root = make([]File, 0)
	elements, err := root.createFiletree(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	root.Root = elements
	var data []byte
	if pretty {
		data, err = json.MarshalIndent(root, "", "    ")
	} else {
		data, err = json.Marshal(root)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
