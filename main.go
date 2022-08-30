package main

import (
  "log"
  "fmt"
  "os"
  "reflect"
  "io"
)

func main() {
  LOCAL_DIR := "./local"
  REMOTE_DIR := "./remote"

  ScannedLocal := scan(LOCAL_DIR)

  defer ScannedLocal.Close()

  local,_ := ScannedLocal.Readdir(0);
  ScannedRemote := scan(REMOTE_DIR)

  defer ScannedLocal.Close()
  remote,_ := ScannedRemote.Readdir(0);

  compared := reflect.DeepEqual(local, remote)
  if compared {
    fmt.Printf("True")
  } else {
    for _, localFile := range local {
      existed := false
  
      fmt.Println(len(remote))

      if len(remote) != 0 {
        for _, remoteFile := range remote {
          if reflect.DeepEqual(localFile, remoteFile) {
            existed = true
            fmt.Printf("MATCH")
          }
        }
      } else {
        src := LOCAL_DIR +"/" + localFile.Name()
        dst := REMOTE_DIR +"/" + localFile.Name()
        copiedByte, err := copy(src, dst)
        if err != nil {
          log.Fatalf("Failed to copy the files, %s", err)
        }
        fmt.Printf("Copied %s", copiedByte)
        existed = true;
      }

      if !existed {
        src := LOCAL_DIR +"/" + localFile.Name()
        dst := REMOTE_DIR +"/" + localFile.Name()
        copiedByte, err := copy(src, dst)
        if err != nil {
          log.Fatalf("Failed to copy the files, %s", err)
        }
        fmt.Printf("Copied %s", copiedByte)
      }
    }
  }
}

func scan(dir string) *os.File {
  local, err := os.Open(dir)
  if err != nil {
    log.Fatalf("Failed to open the directory, %s", err)
  }

  return local
}

func copy(src, dst string) (int64, error) {
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