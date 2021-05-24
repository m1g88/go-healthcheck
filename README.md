# Healthcheck
 Checking website and response time when given the CSV list, calls the Healthcheck Report API to send the statistic of each website

## Environment using `GoDotEnv`( https://github.com/joho/godotenv )
In `.env` file contain required key below
```
ACCESS_TOKEN=
HEALTHCEHECK_ENDPOINT=https://backend-challenge.line-apps.com/healthcheck/report
REQUEST_TIMEOUT=5
```

How to get `ACCESS_TOKEN` [here](https://developers.line.biz/en/docs/line-login/integrate-line-login/)

> Included ACCESS_TOKEN Valid for 30 days after 23-05-2021
## Local
* `go get`
* `go run main.go test.csv`

## via Docker
 * Push expected csv in the root folder name `test.csv`
 * `docker build -t go-healthcheck .`
 * `docker run -it --rm go-healthcheck test.csv`
