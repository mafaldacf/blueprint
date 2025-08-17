#!/bin/bash

WORKSPACE_NAME="order_service_container"
WORKSPACE_DIR=$(pwd)

usage() { 
	echo "Usage: $0 [-h]" 1>&2
	echo "  Required environment variables:"
	
	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    CART_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    CART_SERVICE_THRIFT_DIAL_ADDR=$CART_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${ORDER_DB_DIAL_ADDR+x}" ]; then
		echo "    ORDER_DB_DIAL_ADDR (missing)"
	else
		echo "    ORDER_DB_DIAL_ADDR=$ORDER_DB_DIAL_ADDR"
	fi
	if [ -z "${ORDER_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "    ORDER_SERVICE_THRIFT_BIND_ADDR (missing)"
	else
		echo "    ORDER_SERVICE_THRIFT_BIND_ADDR=$ORDER_SERVICE_THRIFT_BIND_ADDR"
	fi
	if [ -z "${PAYMENT_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    PAYMENT_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    PAYMENT_SERVICE_THRIFT_DIAL_ADDR=$PAYMENT_SERVICE_THRIFT_DIAL_ADDR"
	fi
	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "    SHIPPING_SERVICE_THRIFT_DIAL_ADDR (missing)"
	else
		echo "    SHIPPING_SERVICE_THRIFT_DIAL_ADDR=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR"
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


order_service_proc() {
	cd $WORKSPACE_DIR
	
	if [ -z "${USER_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! user_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! cart_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${PAYMENT_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! payment_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		if ! shipping_service_thrift_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${ORDER_DB_DIAL_ADDR+x}" ]; then
		if ! order_db_dial_addr; then
			return $?
		fi
	fi

	if [ -z "${ORDER_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		if ! order_service_thrift_bind_addr; then
			return $?
		fi
	fi

	run_order_service_proc() {
		
        cd order_service_proc
        ./order_service_proc --user_service.thrift.dial_addr=$USER_SERVICE_THRIFT_DIAL_ADDR --cart_service.thrift.dial_addr=$CART_SERVICE_THRIFT_DIAL_ADDR --payment_service.thrift.dial_addr=$PAYMENT_SERVICE_THRIFT_DIAL_ADDR --shipping_service.thrift.dial_addr=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR --order_db.dial_addr=$ORDER_DB_DIAL_ADDR --order_service.thrift.bind_addr=$ORDER_SERVICE_THRIFT_BIND_ADDR &
        ORDER_SERVICE_PROC=$!
        return $?

	}

	if run_order_service_proc; then
		if [ -z "${ORDER_SERVICE_PROC+x}" ]; then
			echo "${WORKSPACE_NAME} error starting order_service_proc: function order_service_proc did not set ORDER_SERVICE_PROC"
			return 1
		else
			echo "${WORKSPACE_NAME} started order_service_proc"
			return 0
		fi
	else
		exitcode=$?
		echo "${WORKSPACE_NAME} aborting order_service_proc due to exitcode ${exitcode} from order_service_proc"
		return $exitcode
	fi
}


run_all() {
	echo "Running order_service_container"

	# Check that all necessary environment variables are set
	echo "Required environment variables:"
	missing_vars=0
	if [ -z "${CART_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  CART_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  CART_SERVICE_THRIFT_DIAL_ADDR=$CART_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${ORDER_DB_DIAL_ADDR+x}" ]; then
		echo "  ORDER_DB_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  ORDER_DB_DIAL_ADDR=$ORDER_DB_DIAL_ADDR"
	fi
	
	if [ -z "${ORDER_SERVICE_THRIFT_BIND_ADDR+x}" ]; then
		echo "  ORDER_SERVICE_THRIFT_BIND_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  ORDER_SERVICE_THRIFT_BIND_ADDR=$ORDER_SERVICE_THRIFT_BIND_ADDR"
	fi
	
	if [ -z "${PAYMENT_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  PAYMENT_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  PAYMENT_SERVICE_THRIFT_DIAL_ADDR=$PAYMENT_SERVICE_THRIFT_DIAL_ADDR"
	fi
	
	if [ -z "${SHIPPING_SERVICE_THRIFT_DIAL_ADDR+x}" ]; then
		echo "  SHIPPING_SERVICE_THRIFT_DIAL_ADDR (missing)"
		missing_vars=$((missing_vars+1))
	else
		echo "  SHIPPING_SERVICE_THRIFT_DIAL_ADDR=$SHIPPING_SERVICE_THRIFT_DIAL_ADDR"
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

	order_service_proc
	
	wait
}

run_all
