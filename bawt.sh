#!/bin/bash

if [[ "x$REDIS_NAME" != "x" ]]; then
    export HAL_REDIS_URL=$REDIS_PORT_6379_TCP_ADDR:$REDIS_PORT_6379_TCP_PORT
fi

/bawt/bin/bawt
