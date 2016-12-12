#!/bin/bash
pkill main
nohup ./main &
disown
exit 0
