#!/bin/bash

WORKSPACE_NAME="frontend_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    CART_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    CART_SERVICE_THRIFT_DIAL_ADDR=$CART_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${CATALOGUE_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    CATALOGUE_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    CATALOGUE_SERVICE_THRIFT_DIAL_ADDR=$CATALOGUE_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${FRONTEND_HTTP_BIND_ADDR+x}" ]; then
		echo "    FRONTEND_HTTP_BIND_ADDR (missing)"
	else
		echo "    FRONTEND_HTTP_BIND_ADDR=$FRONTEND_HTTP_BIND_ADDR"
	fi
	if [ -z "${ORDER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    ORDER_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    ORDER_SERVICE_THRIFT_DIAL_ADDR=$ORDER_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${USER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    USER_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    USER_SERVICE_THRIFT_DIAL_ADDR=$USER_SERVICE_THRIFT_DIAL_ADDR"
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


frontend_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${USER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! user_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${CATALOGUE_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! catalogue_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! cart_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${ORDER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! order_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${FRONTEND_HTTP_BIND_ADDR+x}" ]; then
		if ! frontend_http_bind_addr; then
			return $?
		fi
	fi

	run_frontend_service_proc() {
		
        cd frontend_service_proc
        ./frontend_service_proc --user_service.thrift.dial_addr=$USER_SERVICE_THRIFT_DIAL_ADDR --catalogue_service.thrift.dial_addr=$CATALOGUE_SERVICE_THRIFT_DIAL_ADDR --cart_service.thrift.dial_addr=$CART_SERVICE_THRIFT_DIAL_ADDR --order_service.thrift.dial_addr=$ORDER_SERVICE_THRIFT_DIAL_ADDR --frontend.http.bind_addr=$FRONTEND_HTTP_BIND_ADDR &
        FRONTEND_SERVICE_PROC=$!
        return $?

	}

	if run_frontend_service_proc; then
		if [ -z "${FRONTEND_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting frontend_service_proc: function frontend_service_proc did not set FRONTEND_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started frontend_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting frontend_service_proc due to exitcode ${exitcode} from frontend_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running frontend_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  CART_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CART_SERVICE_THRIFT_DIAL_ADDR=$CART_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${CATALOGUE_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  CATALOGUE_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CATALOGUE_SERVICE_THRIFT_DIAL_ADDR=$CATALOGUE_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${FRONTEND_HTTP_BIND_ADDR+x}" ]; then
		echo "  FRONTEND_HTTP_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  FRONTEND_HTTP_BIND_ADDR=$FRONTEND_HTTP_BIND_ADDR"
	fi
	
	if [ -z "${ORDER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  ORDER_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  ORDER_SERVICE_THRIFT_DIAL_ADDR=$ORDER_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${USER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  USER_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  USER_SERVICE_THRIFT_DIAL_ADDR=$USER_SERVICE_THRIFT_DIAL_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	frontend_service_proc
	
	wait
}

run_all
