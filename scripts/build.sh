#!/bin/bash

PROGRAM_NAME="gidx"
DIST_DIR="./dist"

rm -rf $DIST_DIR
go build -o $PROGRAM_NAME && mkdir dist && mv ./$PROGRAM_NAME $DIST_DIR

