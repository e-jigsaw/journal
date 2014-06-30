package main

import (
  "fmt"
  "os"
  "os/user"
  "bufio"
  "path"
  "time"
  "strconv"
  "net/smtp"
  "io/ioutil"
  "encoding/json"
  "github.com/codegangsta/cli"
)

var Commands = []cli.Command {
  commandComment,
  commandSend,
  commandWrite,
}

var commandComment = cli.Command {
  Name: "comment",
  ShortName: "c",
  Usage: "write comment",
  Description: "write comment",
  Action: doComment,
}

var commandSend = cli.Command {
  Name: "send",
  ShortName: "s",
  Usage: "send mail",
  Description: "send mail",
  Action: doSend,
}

var commandWrite = cli.Command {
  Name: "write",
  ShortName: "w",
  Usage: "write journal",
  Description: "write journal",
  Action: doWrite,
}

func ZeroComp(str string) string {
  if len(str) == 1 {
    str = "0" + str
  }
  return str
}

func doComment(c *cli.Context) {
  scanner := bufio.NewScanner(os.Stdin)
  usr, _ := user.Current()
  fout, err := os.OpenFile(path.Join(usr.HomeDir, ".comment"), os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fout, err = os.Create(path.Join(usr.HomeDir, ".comment"))
  }
  defer fout.Close()
  fmt.Print("Title: ")
  for scanner.Scan() {
    fout.WriteString("## " + scanner.Text() + "\n\n")
    break
  }
  for i := 1; i <= 3; i++ {
    fmt.Print("Comment", i, ": ")
    for scanner.Scan() {
      fout.WriteString("* " + scanner.Text() + "\n")
      break
    }
  }
  fout.WriteString("\n")
}

type Config struct {
  Mail string
  Pass string
  To   string
  Subj string
}

func doSend(c *cli.Context) {
  usr, _ := user.Current()
  conf, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".journal.config.json"))
  schedule, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".journal"))
  comment, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".comment"))
  if err != nil {
    fmt.Println(err)
    return
  }
  var config Config
  json.Unmarshal(conf, &config)
  body := "To: " + config.To +
    "\r\nSubject: " + config.Subj +
    "\r\n\r\n# 本日の業務内容\n\ntime  | description\n----- | ----\n" + string(schedule) +
    "\n# 所感\n\n" + string(comment) +
    "\n---" +
    "\nこの日報は激ヤバ鬼便利日報システム改( https://github.com/e-jigsaw/journal )によって送信されました\n"
  auth := smtp.PlainAuth("", config.Mail, config.Pass, "smtp.gmail.com")
  from := make([]string, 1)
  from[0] = string(config.Mail)
  err = smtp.SendMail("smtp.gmail.com:587", auth, config.To, from, []byte(body))
  if err != nil {
    fmt.Println(err)
    return
  }
  err = os.Remove(path.Join(usr.HomeDir, ".journal"))
  if err != nil {
    fmt.Println(err)
    return
  }
  err = os.Remove(path.Join(usr.HomeDir, ".comment"))
  if err != nil {
    fmt.Println(err)
    return
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
  fout.WriteString(ZeroComp(strconv.Itoa(hour)) + ":" + ZeroComp(strconv.Itoa(min)) + " | ")
  for i := 0; i < len(c.Args()); i++ {
    fout.WriteString(c.Args()[i])
  }
  fout.WriteString("\n")
}
