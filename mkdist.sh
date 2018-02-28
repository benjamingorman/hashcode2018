#!/bin/bash
mkdir -p dist
cp solutions/* dist
zip dist/src.zip *.go
