
# Mock AWS SES

Mock AWS SES is a mock server to mock the AWS SES API.

Currently the valid identities are (test@demtech.ai, test@test.com, test@gmail.com). Only they can send mail.
The default Max24HourSend is 10000 and maximum emails per second is 14.

Current all the email sending will be success. The list of destination email for which a email send would fail will be done later. (Check TODO)


# How to start server:
1. Clone the code `git clone git@github.com:libaishwarya/mock-aws-ses-go.git`
2. Run `go mod tidy` from the folder
3. Run `go run main.go`
4. Server will be running in `8080`


# To run tests:
`go test ./...`


# TODO
* Fix error messages as per AWS spec
* OpenAPI Spec
* Add configuration/list where email send would fail/reject
* Optional(Relational store)
* Time intervaled stats
* Ratelimiting
* Coverage check

# API Reference
https://docs.aws.amazon.com/ses/latest/APIReference
