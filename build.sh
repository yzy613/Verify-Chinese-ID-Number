#!/bin/bash

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o verify-chinese-ID-number ./main.go
tar -czvf linux_amd64.tar.gz verify-chinese-ID-number area_code.json
rm -f verify-chinese-ID-number

CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o verify-chinese-ID-number.exe ./main.go
tar -czvf windows_amd64.tar.gz verify-chinese-ID-number.exe area_code.json
rm -f verify-chinese-ID-number.exe