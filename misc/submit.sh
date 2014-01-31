#!/bin/bash

maxperiod=10
randmin=10
randmod=100

git_dir=$1/.git
work_tree=$1

workdir=submit/$1/$2

n=$3
while ((i < n)); do
    mkdir -p $workdir
    ((i++))
    out=$workdir/out.$(printf %04d $i)
    echo $(date) run $i begins ...
    (git --git-dir $git_dir --work-tree $work_tree push -f origin master >$out 2>&1)&
    pid=$!
    j=0
    while ((j < maxperiod)); do
        sleep=$(($randmin + $RANDOM % $randmod))
        echo $(date) sleeping for $sleep seconds ...
        sleep $sleep
        if ! ps -p $pid >/dev/null; then
            break
        fi
    done
    if ps -p $pid >/dev/null; then
        echo process still running, killing it
        kill -9 $pid
    fi
    grep -Eio 'your.*score.*\.|You scored.*points' $out
done
