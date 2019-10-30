#!/bin/bash
#
#   Copyright © 2019 M.Watermann, 10247 Berlin, Germany
#               All rights reserved
#           EMail : <support@mwat.de>
#
# ---------------------------------------------------------------------------

set -u #-x

# Number of concurrent instances:
INSTANCES=5

# path/file of `krabbel` executable:
KRABBEL='./krabbel'

# ---------------------------------------------------------------------------
# NOTHING TO CHANGE BEYOND THIS POINT

# check whether the executable exists:
[ -x "${KRABBEL}" ] || : ${CANT_FIND:?"$KRABBEL"}

# check the URL to read:
URL="${1:?MISSING_URL_ARGUMENT}"

LOOP=0
while [ ${LOOP} -lt ${INSTANCES} ] ; do
	${KRABBEL} ${URL} >/dev/null &
	let LOOP=$[LOOP + 1]
done

echo
#_EoF_
