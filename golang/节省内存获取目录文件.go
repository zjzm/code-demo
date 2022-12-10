package main

import (
  "fmt"
  "io/ioutil"
)

type DirectoryFiles struct {
  directory string
  files []string
  index int
}

func (df *DirectoryFiles) Next() bool {
  if df.index >= len(df.files) {
    df.files, _ = ioutil.ReadDir(df.directory)
    df.index = 0
  }
  if df.index >= len(df.files) {
    return false
  }
  df.index++
  return true
}

func (df *DirectoryFiles) Value() string {
  return df.files[df.index - 1].Name()
}

func main() {
  df := &DirectoryFiles{directory: "/path/to/directory"}
  for df.Next() {
    fmt.Println(df.Value())
  }
}
