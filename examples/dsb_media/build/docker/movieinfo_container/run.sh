#!/bin/bash

WORKSPACE_NAME="movieinfo_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${MOVIEINFO_DB_DIAL_ADDR+x}" ]; then
		echo "    MOVIEINFO_DB_DIAL_ADDR (missing)"
	else
		echo "    MOVIEINFO_DB_DIAL_ADDR=$MOVIEINFO_DB_DIAL_ADDR"
	fi
	if [ -z "${MOVIEINFO_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    MOVIEINFO_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    MOVIEINFO_SERVICE_THRIFT_BIND_ADDR=$MOVIEINFO_SERVICE_THRIFT_BIND_ADDR"
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


movieinfo_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${MOVIEINFO_DB_DIAL_ADDR+x}" ]; then
		if ! movieinfo_db_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${MOVIEINFO_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! movieinfo_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_movieinfo_proc() {
		
        cd movieinfo_proc
        ./movieinfo_proc --movieinfo_db.dial_addr=$MOVIEINFO_DB_DIAL_ADDR --movieinfo_service.thrift.bind_addr=$MOVIEINFO_SERVICE_THRIFT_BIND_ADDR &
        MOVIEINFO_PROC=$!
        return $?

	}

	if run_movieinfo_proc; then
		if [ -z "${MOVIEINFO_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting movieinfo_proc: function movieinfo_proc did not set MOVIEINFO_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started movieinfo_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting movieinfo_proc due to exitcode ${exitcode} from movieinfo_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running movieinfo_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${MOVIEINFO_DB_DIAL_ADDR+x}" ]; then
		echo "  MOVIEINFO_DB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEINFO_DB_DIAL_ADDR=$MOVIEINFO_DB_DIAL_ADDR"
	fi
	
	if [ -z "${MOVIEINFO_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  MOVIEINFO_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEINFO_SERVICE_THRIFT_BIND_ADDR=$MOVIEINFO_SERVICE_THRIFT_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	movieinfo_proc
	
	wait
}

run_all
