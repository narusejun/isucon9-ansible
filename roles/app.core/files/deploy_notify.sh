#!/bin/bash

. /etc/profile
. deploy.sh 2>&1 | notify_slack
