#!/bin/bash
BUILD_DATE=`date +%Y%m%d_%H`
docker build --tag innodepcloud.azurecr.io/dms-demo-service:dev-latest --tag innodepcloud.azurecr.io/dms-demo-service:$BUILD_DATE .
docker push innodepcloud.azurecr.io/dms-demo-service:dev-latest
docker push innodepcloud.azurecr.io/dms-demo-service:$BUILD_DATE