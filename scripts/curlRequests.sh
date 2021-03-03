#!/bin/bash

# A collection of simple curl requests that can be used to manually test endpoints before and while writing automated tests

curl localhost:9090/queue
curl localhost:9090/queue/42
curl localhost:9090/queue -XPOST --header "Content-Type: application/json" -d '{"userid":13}'
curl localhost:9090/delete/42
