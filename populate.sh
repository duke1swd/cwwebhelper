#!/bin/sh
#
# Read the output of "cwwebhelper users" and make sure those
# "users" exist and are subscribed as appropriate.
#

MM3=../mm3util/mm3utilx

while IFS=, read -r email cwmail www burp
do
	${MM3} user add $email
	if [ $? -ne 0 ] ; then
		echo cannot create user $email 1>&2
		exit 1
	fi
	if $cwmail == "true"; then
		${MM3} subscribe cwmail $email
	fi
	if $www == "true"; then
		${MM3} subscribe www $email
	fi
	if $burp == "true"; then
		${MM3} subscribe burp $email
	fi
done
