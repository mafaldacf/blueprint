#!/bin/bash

WORKSPACE_NAME="queue_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    SHIPPING_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    SHIPPING_SERVICE_THRIFT_DIAL_ADDR=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR"
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


queue_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${SHIPQUEUE_DIAL_ADDR+x}" ]; then
		if ! shipqueue_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! shipping_service_thrift_dial_addr; then
			return $?
		fi
	fi

	run_queue_service_proc() {
		
        cd queue_service_proc
        ./queue_service_proc --shipqueue.dial_addr=$SHIPQUEUE_DIAL_ADDR --shipping_service.thrift.dial_addr=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR &
        QUEUE_SERVICE_PROC=$!
        return $?

	}

	if run_queue_service_proc; then
		if [ -z "${QUEUE_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting queue_service_proc: function queue_service_proc did not set QUEUE_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started queue_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting queue_service_proc due to exitcode ${exitcode} from queue_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running queue_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  SHIPPING_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  SHIPPING_SERVICE_THRIFT_DIAL_ADDR=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR"
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

	queue_service_proc
	
	wait
}

run_all
