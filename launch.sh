#!/bin/bash
pkill main
sleep 1
nohup ./main & > /dev/null &
exit 0
