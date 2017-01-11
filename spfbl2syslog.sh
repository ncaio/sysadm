#!/bin/bash
#
#       Noilson Caio 10-01-2007
#
_logpath="/var/log/spfbl/"
_logdate="$(date +%Y-%m-%d)"
#
#
#
tail -f "$_logpath"spfbl."$_logdate".log |
#
#
#
while read -r line
do
   newline="$(echo "$line" | awk '{$1=$2=""; print $0}')"
   logger -p local3.notice -t SPFBL "$newline"
done
#
#
#
