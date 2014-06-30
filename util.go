package main

import (
  "path"
  "os/user"
)

func ZeroComp(str string) string {
  if len(str) == 1 {
    str = "0" + str
  }
  return str
}

func HomePath(file string) string {
  usr, _ := user.Current()
  return path.Join(usr.HomeDir, file)
}
