#!/bin/bash

WORKSPACE_NAME="catalogue_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${CATALOGUE_DB_DIAL_ADDR+x}" ]; then
		echo "    CATALOGUE_DB_DIAL_ADDR (missing)"
	else
		echo "    CATALOGUE_DB_DIAL_ADDR=$CATALOGUE_DB_DIAL_ADDR"
	fi
	if [ -z "${CATALOGUE_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    CATALOGUE_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    CATALOGUE_SERVICE_THRIFT_BIND_ADDR=$CATALOGUE_SERVICE_THRIFT_BIND_ADDR"
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


catalogue_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${CATALOGUE_DB_DIAL_ADDR+x}" ]; then
		if ! catalogue_db_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${CATALOGUE_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! catalogue_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_catalogue_service_proc() {
		
        cd catalogue_service_proc
        ./catalogue_service_proc --catalogue_db.dial_addr=$CATALOGUE_DB_DIAL_ADDR --catalogue_service.thrift.bind_addr=$CATALOGUE_SERVICE_THRIFT_BIND_ADDR &
        CATALOGUE_SERVICE_PROC=$!
        return $?

	}

	if run_catalogue_service_proc; then
		if [ -z "${CATALOGUE_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting catalogue_service_proc: function catalogue_service_proc did not set CATALOGUE_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started catalogue_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting catalogue_service_proc due to exitcode ${exitcode} from catalogue_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running catalogue_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${CATALOGUE_DB_DIAL_ADDR+x}" ]; then
		echo "  CATALOGUE_DB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CATALOGUE_DB_DIAL_ADDR=$CATALOGUE_DB_DIAL_ADDR"
	fi
	
	if [ -z "${CATALOGUE_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  CATALOGUE_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CATALOGUE_SERVICE_THRIFT_BIND_ADDR=$CATALOGUE_SERVICE_THRIFT_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	catalogue_service_proc
	
	wait
}

run_all
