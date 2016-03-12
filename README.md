## About

Companion package to https://github.com/gologs/log.

Implementation of gologs/log/logger [Logger](https://godoc.org/github.com/gologs/log/logger#Logger)
that pipes messages to syslog.

## Demo

```sh
$ !go
go run cmd/demo/demo.go 
demo: hello % world
demo: this is info
demo: this is warn
demo: this is error
demo: this is fatal
exit status 1
```
