package main

import (
  "os"
  "github.com/codegangsta/cli"
)

func main() {
  newApp().Run(os.Args)
}

func newApp() *cli.App {
  app := cli.NewApp()
  app.Name = "journal"
  app.Usage = "Journal for Daily Reports"
  app.Version = "0.1.1"
  app.Author = "jigsaw"
  app.Email = "m@jgs.me"
  app.Commands = Commands
  return app
}
