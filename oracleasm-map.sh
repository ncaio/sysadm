#!/bin/bash
#
#
#
v() {
        ls -l /dev/* | grep disk | grep "$1," > /tmp/.lista
	while read linha
	do
	k="$(echo "$linha" | awk '{print $6}')"	
	if [ "$k" -eq "$2" ]
		then
			echo "$linha" | awk '{print $10}'
		fi
	done < /tmp/.lista
}
#
#
#
for i in $(oracleasm listdisks)
do 
	id=$(oracleasm querydisk -d "$i" | awk '{print $10}' | tr -d "]" | tr -d "[")
	p1=$(echo "$id" | cut -d',' -f1)
	p2=$(echo "$id" | cut -d',' -f2)
	#echo $id $p1 $p2
	echo -n "$i : "
	v "$p1" "$p2"
done
