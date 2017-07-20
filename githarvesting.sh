#!/bin/bash
#
#
#
_repo="$1"
_list="/tmp/githarvesting.log"
#
#
#
if [ -z "$_repo" ] 
then
	echo "./script /path/to/repo"
	exit 1
fi
#
#
#
if [ ! -z "$_repo" ] && [ ! -d "$_repo" ]
then
	echo "$_repo is not a directory"
	exit 1
fi
#
#
#
if [ ! -d "$_repo"/.git ]
then
	echo "$_repo Not a git repository"
	exit
else
	cd "$_repo"
	git log > "$_list"	
fi
#
#
#
echo "First Name,Last Name,Position,Email"
#
#
#
while read _line;
do
	echo "$_line" | grep -q ^"Author:"
	_key="$?"
	if [ "$_key" == "0" ]
	then
		_name=$(echo "$_line" | cut -d':' -f 2 | cut -d'<' -f 1 | tr -d '"' | sed 's/^ *//')
		_firstname="$(echo "$_name" | awk '{print $1}')"
		if [ -z "$_firstname" ]
		then
			_firstname="First Name"
		fi
		_lastname="$(echo $_name | awk '{print $NF}')"
		_position="$(echo "$_line"|awk '{print $1}'|tr -d ':')"
		if [ -z "$_position" ]
		then
			_position="Author"
		fi
		_email="$(echo "$_line" | grep -E -o "\b[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,6}\b")"
		if [ ! -z "$_firstname" ] && [ ! -z "$_lastname" ] && [ ! -z "$_position" ] && [ ! -z "$_email" ]
		then
			echo "$_firstname,$_lastname,$_position,$_email"
		fi
	fi
done < "$_list"
