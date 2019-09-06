#!/bin/sh

if [ -z $1 ]; then
	echo Usage: $0 [filepath]
	exit 1
elif [ ! -e $1 ]; then
	echo No such file: $1
	exit 1
fi

curl -X POST https://slack.com/api/files.upload \
-F token={{ slack_token }} \
-F channels={{ slack_channel_stdout }} \
-F file=@$1
