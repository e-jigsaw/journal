package main

import (
  "fmt"
  "os"
  "bufio"
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

func doComment(c *cli.Context) {
  scanner := bufio.NewScanner(os.Stdin)
  fout, err := os.OpenFile(HomePath(".comment"), os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fout, err = os.Create(HomePath(".comment"))
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
  conf, err := ioutil.ReadFile(HomePath(".journal.config.json"))
  schedule, err := ioutil.ReadFile(HomePath(".journal"))
  comment, err := ioutil.ReadFile(HomePath(".comment"))
  if err != nil {
    fmt.Println(err)
    return
  }
  var config Config
  json.Unmarshal(conf, &config)
  body := "From:" + config.Mail +
    "\r\nTo:" + config.To +
    "\r\nSubject:" + config.Subj +
    "\r\n\r\n# 本日の業務内容\n\ntime  | description\n----- | ----\n" + string(schedule) +
    "\n# 所感\n\n" + string(comment) +
    "\n---" +
    "\nこの日報は激ヤバ鬼便利日報システム改( https://github.com/e-jigsaw/journal )によって送信されました\n"
  auth := smtp.PlainAuth("", config.Mail, config.Pass, "smtp.gmail.com")
  err = smtp.SendMail("smtp.gmail.com:587", auth, config.Mail, []string{config.To}, []byte(body))
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println("Sent: ", body)
  err = os.Remove(HomePath(".journal"))
  if err != nil {
    fmt.Println(err)
    return
  }
  err = os.Remove(HomePath(".comment"))
  if err != nil {
    fmt.Println(err)
    return
  }
}

func doWrite(c *cli.Context) {
  hour, min, _ := time.Now().Clock()
  fout, err := os.OpenFile(HomePath(".journal"), os.O_RDWR|os.O_APPEND, 0660)
  if err != nil {
    fout, err = os.Create(HomePath(".journal"))
  }
  defer fout.Close()
  fout.WriteString(ZeroComp(strconv.Itoa(hour)) + ":" + ZeroComp(strconv.Itoa(min)) + " | ")
  for i := 0; i < len(c.Args()); i++ {
    fout.WriteString(c.Args()[i])
  }
  fout.WriteString("\n")
}
