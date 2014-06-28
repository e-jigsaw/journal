# journal(a.k.a.激ヤバ鬼便利日報システム改)

## Name

journal - Write journal for daily report mail

## Description

Make easy to write journal, For hackers.

```
% journal write "do something"
% journal comment
Title: Golang is awesome
Comment1: Golang is easy to make cli tool
Comment2: Golang is fun
Comment3: I am beginner gopher
% journal send
```

Sent mail:

```
# 本日の業務内容

time | description
---- | -----------
12:00| do something

# 所感

## Golang is awesome

* Golang is easy to make cli tool
* Golang is fun
* I am beginner gopher
```

## Commands

### write

`write` is logging what and when.

#### alias

* w

```
% journal write "something to do"
% journal w "some"
```

### comment

`comment` is tell your notice, thought and idea. Based by "一個三個" framework. It is extreamly simple to tell something others.

#### alias

* c

```
% journal comment
Title: Some thoughts
Comment1: first
Comment2: second
Comment3: third
```

### send

`send` is send mail by configured destination.

#### alias

* s

```
% journal send
```

## Configuration

Journal read `~/.journal.config.json`.

Sample:

```
{
  "Mail": "from@gmail.com",
  "Pass": "passwd",
  "To": "to@gmail.com, to2@gmail.com",
  "Subj": "【日報】"
}
```

## Installation

If installed go:

```
% go get github.com/e-jigsaw/journal
```

## Author

jigsaw (http://jgs.me)
