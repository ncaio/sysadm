#!/bin/bash
#
#       ESTE SCRIPT REALIZA A LEITURA DOS LOGs DO SPFBL 
#       E TRANSPORTA PARA O SYSLOG VIA COMMANDO logger
#       TOPICO: https://groups.google.com/forum/#!topic/spfbl/Xo8dbYKq4IU
#       Noilson Caio 10-01-2007
#
_logpath="/var/log/spfbl/"
_logdate="$(date +%Y-%m-%d)"
_facility="local3"
_level="notice"
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
   logger -p "$_facility"."$_level" -t SPFBL "$newline"
done
#
#
#
