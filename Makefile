run:
	air

swagger:
	swag init

test-auth:
	go test ./auth/tests

test-auth-cover:
	go test ./auth/tests -cover