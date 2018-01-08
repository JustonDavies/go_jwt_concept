#!/usr/bin/env bash

#-- Requirements -------------------------------------------------------------------------------------------------------
command -v jq   >/dev/null 2>&2 || { echo >&2 "jq is required but is not available, aborting...";   exit 1; }
command -v curl >/dev/null 2>&2 || { echo >&2 "curl is required but is not available, aborting..."; exit 1; }

#-- Variables ----------------------------------------------------------------------------------------------------------
path=/auth/login
method=POST
server=http://localhost:8080

authorization=''
input=@./examples/create_session_data.json

output=./tmp/create_session_response.json

#-- Pre-conditions -----------------------------------------------------------------------------------------------------
rm -f $output

#-- Action -------------------------------------------------------------------------------------------------------------
curl -X $method                                      \
     --verbose                                       \
     --output $output                                \
     --header "Content-Type: application/json"       \
     --header "Authorization: Bearer $authorization" \
     --data $input                                   \
     $server$path

#-- Post-Conditions ----------------------------------------------------------------------------------------------------
sleep 0.1 && cat $output