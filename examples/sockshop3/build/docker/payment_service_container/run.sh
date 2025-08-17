#!/bin/bash

WORKSPACE_NAME="payment_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${PAYMENT_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    PAYMENT_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    PAYMENT_SERVICE_THRIFT_BIND_ADDR=$PAYMENT_SERVICE_THRIFT_BIND_ADDR"
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


payment_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${PAYMENT_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! payment_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_payment_service_proc() {
		
        cd payment_service_proc
        ./payment_service_proc --payment_service.thrift.bind_addr=$PAYMENT_SERVICE_THRIFT_BIND_ADDR &
        PAYMENT_SERVICE_PROC=$!
        return $?

	}

	if run_payment_service_proc; then
		if [ -z "${PAYMENT_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting payment_service_proc: function payment_service_proc did not set PAYMENT_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started payment_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting payment_service_proc due to exitcode ${exitcode} from payment_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running payment_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${PAYMENT_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  PAYMENT_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  PAYMENT_SERVICE_THRIFT_BIND_ADDR=$PAYMENT_SERVICE_THRIFT_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	payment_service_proc
	
	wait
}

run_all
