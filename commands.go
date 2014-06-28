package main

import (
  "fmt"
  "os"
  "os/user"
  "path"
  "time"
  "strconv"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command {
  commandWrite,
}

var commandWrite = cli.Command {
  Name: "write",
  ShortName: "w",
  Usage: "write journal",
  Description: "write journal",
  Action: doWrite,
}

func doWrite(c *cli.Context) {
  usr, _ := user.Current()
  hour, min, _ := time.Now().Clock()
  fout, err := os.Create(path.Join(usr.HomeDir, ".journal"))
  if err != nil {
    fmt.Println(err)
    return
  }
  defer fout.Close()
  fout.WriteString(strconv.Itoa(hour) + ":" + strconv.Itoa(min) + " | ")
  for i := 0; i < len(c.Args()); i++ {
    fout.WriteString(c.Args()[i])
  }
  fout.WriteString("\n")
}
