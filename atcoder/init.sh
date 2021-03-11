#!/bin/bash

# example: init.sh abc173_d

DIRECTORY=$1
mkdir $DIRECTORY
cp template.go $DIRECTORY/main.go
cd $DIRECTORY