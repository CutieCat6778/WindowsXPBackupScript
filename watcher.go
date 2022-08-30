package main

import (
    "log"
	"fmt"
    "github.com/fsnotify/fsnotify"
	"strings"
	"os"
	"io"
)

var LOCAL_DIR = ".//local"
var REMOTE_DIR = ".//remote"

func main() {

    // Create new watcher.
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        log.Fatal(err)
    }
    defer watcher.Close()

    // Start listening for events.
    go func() {
		log.Println("Starting to watch for files changes.")
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                log.Println("event:", event)
				nodash := strings.Split(event.Name, "\\")
				fileName := strings.TrimSpace(nodash[len(nodash) - 1])
				fileDir := strings.Join(nodash[1:], "\\")
				if strings.Contains(fileName, ".") {
					if event.Op == fsnotify.Remove {
						DeleteFile(fileName)
					} else {
						if fileName != "" {
							filename := ""
							if len(fileName) > 0 {
								filename = fileName
							}
							src := LOCAL_DIR + "//" + filename
							dst := REMOTE_DIR + "//" + filename
							fmt.Println(src, dst);
							copiedByte, err := copy2(src, dst)
							if err != nil {
								log.Println("error:", err)
							}
							fmt.Println(copiedByte)
						}
					}
				} else if !strings.Contains(fileName, ".") {
					dst := REMOTE_DIR + "//" + fileDir
					if event.Op != fsnotify.Remove {
						createFolder(dst)
					} else {
						removeFolder(dst)
					}
				}
				
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("error:", err)
            }
        }
    }()

    // Add a path.
    err = watcher.Add(LOCAL_DIR)
    if err != nil {
        log.Fatal(err)
    }

    // Block main goroutine forever.
    <-make(chan struct{})
}

func DeleteFile(fileName string) {
	err := os.Remove(LOCAL_DIR+fileName)

	if err != nil {
	  fmt.Println(err)
	  return
	}
}

func copy2(src, dst string) (int64, error) {
	fmt.Printf(src)
	sourceFileStat, err := os.Stat(src)
	if err != nil {
			return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
			return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
			return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
			return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func createFolder(name string) {
	os.Mkdir(name, os.ModeDir)
}

func removeFolder(name string) {
	os.RemoveAll(name);
}