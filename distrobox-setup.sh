#!/usr/bin/env bash

dboxname=gostak
dbpack="go nodejs npm"

if [ $# -eq 0 ]; then
    distrobox rm "${dboxname}" -f
    distrobox create \
        --image localhost/bt-archbase \
        --additional-packages "$dbpack"  \
        --name "$dboxname" \
        --init-hooks "$PWD/distrobox-setup.sh $USER"
    exit 0
fi

host_user=$1

if [ ! -f /etc/setup_complete ]; then
    # do initial setup stuff
    touch /etc/setup_complete
fi


# after initial setup, enter the distrobox and...
# go install github.com/air-verse/air@latest

