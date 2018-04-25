#!/bin/ash

etcdctl user add guest
etcdctl user add root
etcdctl user add client
etcdctl user add server

etcdctl role add guest
etcdctl role add root
etcdctl role add client
etcdctl role add server

etcdctl role grant client -path '/example/clients/*' -readwrite
etcdctl role grant server -path '/example/*' -readwrite
