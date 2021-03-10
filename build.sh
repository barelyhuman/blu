#!/bin/bash

rm -rf blu blu.app
go build 
fyne package -os darwin 
mkdir -p blu.app/Contents/Resources/scripts
cp ./scripts/* blu.app/Contents/Resources/scripts

