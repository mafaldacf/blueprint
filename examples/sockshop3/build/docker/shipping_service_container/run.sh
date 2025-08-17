#!/bin/bash

WORKSPACE_NAME="shipping_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${SHIPDB_DIAL_ADDR+x}" ]; then
		echo "    SHIPDB_DIAL_ADDR (missing)"
	else
		echo "    SHIPDB_DIAL_ADDR=$SHIPDB_DIAL_ADDR"
	fi
	if [ -z "${SHIPPING_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    SHIPPING_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    SHIPPING_SERVICE_THRIFT_BIND_ADDR=$SHIPPING_SERVICE_THRIFT_BIND_ADDR"
	fi
	if [ -z "${SHIPQUEUE_DIAL_ADDR+x}" ]; then
		echo "    SHIPQUEUE_DIAL_ADDR (missing)"
	else
		echo "    SHIPQUEUE_DIAL_ADDR=$SHIPQUEUE_DIAL_ADDR"
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


shipping_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${SHIPQUEUE_DIAL_ADDR+x}" ]; then
		if ! shipqueue_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${SHIPDB_DIAL_ADDR+x}" ]; then
		if ! shipdb_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${SHIPPING_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! shipping_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_shipping_service_proc() {
		
        cd shipping_service_proc
        ./shipping_service_proc --shipqueue.dial_addr=$SHIPQUEUE_DIAL_ADDR --shipdb.dial_addr=$SHIPDB_DIAL_ADDR --shipping_service.thrift.bind_addr=$SHIPPING_SERVICE_THRIFT_BIND_ADDR &
        SHIPPING_SERVICE_PROC=$!
        return $?

	}

	if run_shipping_service_proc; then
		if [ -z "${SHIPPING_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting shipping_service_proc: function shipping_service_proc did not set SHIPPING_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started shipping_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting shipping_service_proc due to exitcode ${exitcode} from shipping_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running shipping_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${SHIPDB_DIAL_ADDR+x}" ]; then
		echo "  SHIPDB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  SHIPDB_DIAL_ADDR=$SHIPDB_DIAL_ADDR"
	fi
	
	if [ -z "${SHIPPING_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  SHIPPING_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  SHIPPING_SERVICE_THRIFT_BIND_ADDR=$SHIPPING_SERVICE_THRIFT_BIND_ADDR"
	fi
	
	if [ -z "${SHIPQUEUE_DIAL_ADDR+x}" ]; then
		echo "  SHIPQUEUE_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  SHIPQUEUE_DIAL_ADDR=$SHIPQUEUE_DIAL_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	shipping_service_proc
	
	wait
}

run_all
