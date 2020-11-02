#!/bin/bash

go build -ldflags="-s -w" goshell.go && upx --brute goshell
