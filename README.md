
# Mock AWS SES

[![OpenAPI Spec](https://img.shields.io/badge/Swagger-UI-green)](https://petstore.swagger.io/?url=https://raw.githubusercontent.com/libaishwarya/mock-aws-ses-go/refs/heads/main/openapi.yaml)


[![View API Docs](https://img.shields.io/badge/API-Docs-red)](https://redocly.github.io/redoc/?url=https://raw.githubusercontent.com/libaishwarya/mock-aws-ses-go/refs/heads/main/openapi.yaml)

Mock AWS SES is a mock server to mock the AWS SES API.

Currently the valid identities are (test@demtech.ai, test@test.com, test@gmail.com). Only they can send mail.
The default Max24HourSend is 10000 and maximum emails per second is 14.

Current all the email sending will be success. The list of destination email for which a email send would fail will be done later. (Check TODO)

Set rate limit: 5 requests per second with a burst of 10 (Should be changed according to the AWS) (Check TODO)

# Check the OpenAPI spec for more details about request and response.

# Deployed on:
https://small-mouse-libaishwarya-223bef55.koyeb.app/

# How to start server:
1. Clone the code `git clone git@github.com:libaishwarya/mock-aws-ses-go.git`
2. Run `go mod tidy` from the folder
3. Run `go run main.go`
4. Server will be running in `8080`


# To run tests:
`go test ./...`

Postman collection added for testing.


# TODO
* Fix error messages as per AWS spec
* Add configuration/list where email send would fail/reject
* Optional(Relational store)
* Time intervaled stats
* Ratelimiting as AWS
* Coverage check and show in readme using gihub workflow and code cov.

# API Reference
https://docs.aws.amazon.com/ses/latest/APIReference
