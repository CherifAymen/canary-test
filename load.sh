#!/bin/bash

ADDR=$(kubectl get svc istio-ingressgateway -n istio-system --output jsonpath="{.status.loadBalancer.ingress[0].ip}")
#ADDR="localhost:8080"



while true; do \
    curl  http://$ADDR/sentiment  -H "Content-type: application/json" \
    -d '{"sentence": "I love microservices"}' \
    -s -w "\t Time: %{time_total}s \t Status: %{http_code} \n" -o -; \
    sleep 1;\
    done

