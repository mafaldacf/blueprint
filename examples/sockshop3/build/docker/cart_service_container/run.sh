#!/bin/bash

WORKSPACE_NAME="cart_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${CART_DB_DIAL_ADDR+x}" ]; then
		echo "    CART_DB_DIAL_ADDR (missing)"
	else
		echo "    CART_DB_DIAL_ADDR=$CART_DB_DIAL_ADDR"
	fi
	if [ -z "${CART_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    CART_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    CART_SERVICE_THRIFT_BIND_ADDR=$CART_SERVICE_THRIFT_BIND_ADDR"
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


cart_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${CART_DB_DIAL_ADDR+x}" ]; then
		if ! cart_db_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${CART_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! cart_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_cart_service_proc() {
		
        cd cart_service_proc
        ./cart_service_proc --cart_db.dial_addr=$CART_DB_DIAL_ADDR --cart_service.thrift.bind_addr=$CART_SERVICE_THRIFT_BIND_ADDR &
        CART_SERVICE_PROC=$!
        return $?

	}

	if run_cart_service_proc; then
		if [ -z "${CART_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting cart_service_proc: function cart_service_proc did not set CART_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started cart_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting cart_service_proc due to exitcode ${exitcode} from cart_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running cart_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${CART_DB_DIAL_ADDR+x}" ]; then
		echo "  CART_DB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CART_DB_DIAL_ADDR=$CART_DB_DIAL_ADDR"
	fi
	
	if [ -z "${CART_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  CART_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CART_SERVICE_THRIFT_BIND_ADDR=$CART_SERVICE_THRIFT_BIND_ADDR"
	fi
		

	if [ "$missing_vars" -gt 0 ]; then
		echo "Aborting due to missing environment variables"
		return 1
	fi

	cart_service_proc
	
	wait
}

run_all
