# Healthcheck
 Checking website and response time when given the CSV list, calls the Healthcheck Report API to send the statistic of each website


## Local test
* `go get`
* `go run main.go *.csv`

## via docker test
 * push expected csv in the root folder
 * `docker build -t go-healthcheck .`
 * `docker run -it --rm go-healthcheck test.csv`
