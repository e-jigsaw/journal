package main

import (
  "fmt"
  "os"
  "os/user"
  "bufio"
  "path"
  "time"
  "strconv"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command {
  commandComment,
  commandWrite,
}

var commandComment = cli.Command {
  Name: "comment",
  ShortName: "c",
  Usage: "write comment",
  Description: "write comment",
  Action: doComment,
}

var commandWrite = cli.Command {
  Name: "write",
  ShortName: "w",
  Usage: "write journal",
  Description: "write journal",
  Action: doWrite,
}

func doComment(c *cli.Context) {
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Print("Title: ")
  for scanner.Scan() {
    break
  }
  for i := 1; i <= 3; i++ {
    fmt.Print("Comment", i, ": ")
    for scanner.Scan() {
      break
    }
  }
}

func doWrite(c *cli.Context) {
  usr, _ := user.Current()
  hour, min, _ := time.Now().Clock()
  fout, err := os.OpenFile(path.Join(usr.HomeDir, ".journal"), os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fout, err = os.Create(path.Join(usr.HomeDir, ".journal"))
  }
  defer fout.Close()
  fout.WriteString(strconv.Itoa(hour) + ":" + strconv.Itoa(min) + " | ")
  for i := 0; i < len(c.Args()); i++ {
    fout.WriteString(c.Args()[i])
  }
  fout.WriteString("\n")
}
