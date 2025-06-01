#!/bin/bash

WORKSPACE_NAME="movieid_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${MOVIEID_DB_DIAL_ADDR+x}" ]; then
		echo "    MOVIEID_DB_DIAL_ADDR (missing)"
	else
		echo "    MOVIEID_DB_DIAL_ADDR=$MOVIEID_DB_DIAL_ADDR"
	fi
	if [ -z "${MOVIEID_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    MOVIEID_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    MOVIEID_SERVICE_THRIFT_BIND_ADDR=$MOVIEID_SERVICE_THRIFT_BIND_ADDR"
	fi
		
	exit 1; 
}

while getopts "h" flag; do
	case $flag in
		*)
		usage
		;;
	esac
done


movieid_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${MOVIEID_DB_DIAL_ADDR+x}" ]; then
		if ! movieid_db_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${MOVIEID_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! movieid_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_movieid_proc() {
		
        cd movieid_proc
        ./movieid_proc --movieid_db.dial_addr=$MOVIEID_DB_DIAL_ADDR --movieid_service.thrift.bind_addr=$MOVIEID_SERVICE_THRIFT_BIND_ADDR &
        MOVIEID_PROC=$!
        return $?

	}

	if run_movieid_proc; then
		if [ -z "${MOVIEID_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting movieid_proc: function movieid_proc did not set MOVIEID_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started movieid_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting movieid_proc due to exitcode ${exitcode} from movieid_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running movieid_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${MOVIEID_DB_DIAL_ADDR+x}" ]; then
		echo "  MOVIEID_DB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEID_DB_DIAL_ADDR=$MOVIEID_DB_DIAL_ADDR"
	fi
	
	if [ -z "${MOVIEID_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  MOVIEID_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEID_SERVICE_THRIFT_BIND_ADDR=$MOVIEID_SERVICE_THRIFT_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	movieid_proc
	
	wait
}

run_all
