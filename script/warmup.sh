#!/bin/bash

set -o xtrace

env >> /etc/environment

service syslog-ng start

/app/script/updateAllCharacterIds.sh

cp cronjob /etc/cron.d/updatefile

service cron start

logger "Ready to serve API"

tail -f /var/log/syslog
