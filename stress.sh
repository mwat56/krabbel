#!/bin/bash
#
#   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
#               All rights reserved
#           EMail : <support@mwat.de>
#
# Script to run several instances of `krabbel` concurrently
# to simulate some kind of request-stress for a given URL.
# --------------------------------------------------------------------------
set -u #-x

# Number of concurrent instances:
declare -r -i INSTANCES=5

# path/file of `krabbel` executable:
declare -r KRABBEL='./krabbel'

# --------------------------------------------------------------------------
# NOTHING TO CHANGE BEYOND THIS POINT

# check whether the executable exists:
[ -x "${KRABBEL}" ] || : ${CANT_FIND:?"$KRABBEL"}

# check the URL to read:
URL="${1:?MISSING_URL_ARGUMENT}"

declare -i LOOP=0
while [ ${LOOP} -lt ${INSTANCES} ] ; do
	${KRABBEL} -url="${URL}" -quiet=true &
	let LOOP=$[LOOP + 1]
done

echo
#_EoF_
