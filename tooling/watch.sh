#!/run/current-system/sw/bin/bash

DIR_TO_WATCH=${1}
COMMAND=${2}

DIR_TO_WATCH="."
COMMAND="rsync -rltgD --recursive --human-readable --progress -v /home/mattn/Code/thermostat thermo:/"

trap "echo Exited!; exit;" SIGINT SIGTERM
while [[ 1=1 ]]
do
  watch --chgexit -n 1 "ls --all -l --recursive --full-time ${DIR_TO_WATCH} | sha256sum" && ${COMMAND}
  sleep 1
done
