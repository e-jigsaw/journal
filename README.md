# deprecated

# journal(激ヤバ鬼便利日報システム改)

[![wercker status](https://app.wercker.com/status/89474ea5446159910bb06c7113046e21/s/master "wercker status")](https://app.wercker.com/project/bykey/89474ea5446159910bb06c7113046e21)

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

OSX:

```
% brew tap e-jigsaw/journal
% brew install journal
```

## Author

jigsaw (http://jgs.me)

## License

The MIT License (MIT)

Copyright (c) 2014 motemen

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
