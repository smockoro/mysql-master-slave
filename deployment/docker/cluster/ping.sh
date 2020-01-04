#!/bin/sh

if [ "$1" = 'ndb_mgmd' ]; then
	# Don't proceed until the hostnames are resolvable; otherwise, ndb_mgmd will quit
    HOSTNAMES=${HOSTNAMES:-ndb1 ndb2 management1}
    HOSTNAMES_ARR=(${HOSTNAMES})
    
    for HOSTNAME in "${HOSTNAMES_ARR[@]}"; do
	    echo "Waiting for ${HOSTNAME}"
	    while ! getent hosts "${HOSTNAME}"; do
		    sleep 1
	    done
    done
fi

./entrypoint.sh $1

