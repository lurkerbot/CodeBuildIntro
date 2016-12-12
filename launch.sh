#!/bin/bash
pkill main
./main &
disown -h
exit 0
