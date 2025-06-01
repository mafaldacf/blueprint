#!/bin/bash

WORKSPACE_NAME="api_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${API_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "    API_SERVICE_HTTP_BIND_ADDR (missing)"
	else
		echo "    API_SERVICE_HTTP_BIND_ADDR=$API_SERVICE_HTTP_BIND_ADDR"
	fi
	if [ -z "${MOVIEID_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    MOVIEID_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    MOVIEID_SERVICE_THRIFT_DIAL_ADDR=$MOVIEID_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR=$MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR"
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


api_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${MOVIEID_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! movieid_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! movieinfo_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${API_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		if ! api_service_http_bind_addr; then
			return $?
		fi
	fi

	run_api_proc() {
		
        cd api_proc
        ./api_proc --movieid_service.thrift.dial_addr=$MOVIEID_SERVICE_THRIFT_DIAL_ADDR --movieinfo_service.thrift.dial_addr=$MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR --api_service.http.bind_addr=$API_SERVICE_HTTP_BIND_ADDR &
        API_PROC=$!
        return $?

	}

	if run_api_proc; then
		if [ -z "${API_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting api_proc: function api_proc did not set API_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started api_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting api_proc due to exitcode ${exitcode} from api_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running api_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${API_SERVICE_HTTP_BIND_ADDR+x}" ]; then
		echo "  API_SERVICE_HTTP_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  API_SERVICE_HTTP_BIND_ADDR=$API_SERVICE_HTTP_BIND_ADDR"
	fi
	
	if [ -z "${MOVIEID_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  MOVIEID_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEID_SERVICE_THRIFT_DIAL_ADDR=$MOVIEID_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR=$MOVIEINFO_SERVICE_THRIFT_DIAL_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	api_proc
	
	wait
}

run_all
