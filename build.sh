#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o Verify-Chinese-ID-Number ./main.go
tar -czvf linux_amd64.tar.gz Verify-Chinese-ID-Number
rm -f Verify-Chinese-ID-Number

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o Verify-Chinese-ID-Number.exe ./main.go
tar -czvf windows_amd64.tar.gz Verify-Chinese-ID-Number.exe
rm -f Verify-Chinese-ID-Number.exe