#!/bin/bash

curl --cacert certs/ca.pem --cert certs/etcd.pem --key certs/etcd.key https://127.0.0.1:2379/v2/keys/foo -XPUT -d value=bar -v
