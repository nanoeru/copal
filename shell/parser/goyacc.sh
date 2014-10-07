#!/bin/sh
source ~/.bash_profile
go tool yacc -o parser.go -v parser.output parser.go.y