package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const (
	STAR_FILE = ".star"
	DIR       = "📁"
	FILE      = "📄"
)

type SearchOptions struct {
	editor string
	delete bool
}

func openStar() *os.File {
	starFile := filepath.Join(os.Getenv("HOME"), STAR_FILE)
	file, err := os.OpenFile(starFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file error: %v", err)
	}
	return file
}

func searchStar(file *os.File, options SearchOptions) {
	file.Seek(0, 0)
	content := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := scanner.Text()
		fileInfo, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		if fileInfo.IsDir() {
			content += fmt.Sprintf("%s %s\n", DIR, path)
		} else {
			content += fmt.Sprintf("%s %s\n", FILE, path)
		}
	}
	formatedContent := strings.ReplaceAll(content, os.Getenv("HOME"), "~")
	cmd := exec.Command("fzf")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Printf("Error getting fzf stdin: %v", err)
	}
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Start(); err != nil {
		fmt.Println("Error starting fzf command:", err)
		return
	}
	if _, err := stdin.Write([]byte(formatedContent)); err != nil {
		fmt.Println("Error writing to stdin pipe:", err)
		return
	}
	stdin.Close()
	if err := cmd.Wait(); err != nil {
		return
	}
	result := out.String()
	result = strings.Replace(result, "~", os.Getenv("HOME"), -1)
	result = strings.Split(result, " ")[1]
	if options.delete {
		deleteStar(file, result)
	} else if options.editor != "" {
		openInEditor(result, options.editor)
	} else {
		fmt.Print(result)
	}
}

func deleteStar(file *os.File, record string) bool {
	deleted := false
	file.Seek(0, 0)
	buffer := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() + "\n"
		if strings.TrimSpace(line) == strings.TrimSpace(record) {
			deleted = true
			continue
		}
		buffer += line
	}
	file.Truncate(0)
	file.Seek(0, 0)
	file.WriteString(buffer)
	if deleted {
		fmt.Printf("\x1b[31m✘ %s was unstarred ✘\x1b[0m\n",
			strings.TrimSpace(strings.Replace(record, os.Getenv("HOME"), "~", -1)),
		)
	}
	return deleted
}

func toggleStar(file *os.File, path string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("get current path error: %v", err)
	}
	fullpath := filepath.Join(pwd, path)
	if os.IsNotExist(err) {
		fmt.Println("This file or directory do not exist")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() == fullpath {
			deleteStar(file, fullpath)
			return
		}
	}

	file.WriteString(fullpath + "\n")
	shortPath := strings.Replace(fullpath, os.Getenv("HOME"), "~", -1)
	fmt.Printf("\x1b[33m⭑ %s was starred ⭑\x1b[0m\n", shortPath)
}

func openInEditor(path, editor string) {
	trimmedPath := strings.TrimSpace(path)
	cmd := exec.Command(editor, trimmedPath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = trimmedPath
	err := cmd.Run()
	if err != nil {
		fmt.Println("Couldn't start specified editor, make sure it is installed")
	}
}
