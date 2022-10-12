## Description

Create go http middleware for REST endpoint to verify "Content-Type" header is set to "application/json".

## Run code

```
# start server
go run main.go

# test with curl
curl --header "Content-Type:application/json" localhost:3000
curl --header "Content-Type:text" localhost:3000
curl localhost:3000
```
