#!/bin/bash
maxtime=$1
looptime=1
if [ "$1" -gt "100" ] 
then
    maxtime=100
fi 
while [ "$maxtime" -ge "$looptime" ]
do
    echo "$looptime times I've printed nirabpudasaini"
    looptime=$(($looptime+1))
done
