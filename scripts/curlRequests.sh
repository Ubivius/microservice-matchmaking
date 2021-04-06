#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/queue/42
curl localhost:9090/queue -XPOST --header "Content-Type: application/json" -d '{"user_id":"98af650c-96f4-11eb-a8b3-0242ac130003", "user_ip":"123.123.123.123"}'
curl localhost:9090/queue/42 -XDELETE
