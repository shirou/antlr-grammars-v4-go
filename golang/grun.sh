#!/usr/bin/env bash
export CLASSPATH=".:/usr/local/Cellar/antlr/4.6/antlr-4.6-complete.jar:$CLASSPATH"
SRC=/Users/shirou/Works/golang/src/github.com/shirou/golightan/lexer/testcase/golang/example.go

antlr4 -o java Golang.g4
cd java
grun Golang sourceFile -gui ${SRC}
