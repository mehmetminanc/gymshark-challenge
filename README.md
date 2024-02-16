# Overview
This repo contains an AWS SAM backend app and a rudimentary web UI to utilize that.

`README_SAM.md` contains more information on how the project is organized and setup.

`TODO.md` contains things I wish I had done, but ran out of time. 

## Public endpoints
The API is hosted on AWS APIGW: `https://7rcijo7emg.execute-api.eu-central-1.amazonaws.com/Prod/compute-packs/`
The curl examples would also work against this endpoint. Please don't call it over 1m times :)

The frontend app is hosted on S3 

## Endpoint
I decided to utilize a single POST endpoint in API GW pointing to the lambda.
The request is a JSON string and must contain `order` field with the target number as the value.
Optionally, one can provide `sizes` field with an array of integers to override default pack sizes defined in the doc.
Curl examples can be found below.

## Dockerfile & Local environment
It is intentionally left blank as the SAM CLI takes care of spinning up a local lambda environment. Execute 

```shell
sam local start-api
```
to start up a local stack. To test if everything is in order, execute
```shell
curl http://localhost:3000/compute-packs -d '{"order": 1002}'   
```
```
curl http://localhost:3000/compute-packs -d '{"order": 1002, "sizes":[31,42,99]}'
```

## Shabby UI
It has been a while since I coded any UI, so please wear eye protection while using the app. 
I lack sense of aesthetics, anyway. 
