#!/bin/bash

echo "Downloading dependencies"
go get github.com/gorilla/mux
echo "Building application"
go build main.go
echo "Saving application as /home/btwlinux/btw"
mv main /home/btwlinux/btw
chmod +X /home/btwlinux/btw
echo "Finished"