run:
	air

build:
	docker-compose up --build -d

swagger:
	swag init

test-auth:
	go test ./auth/tests -v

test-auth-cover:
	go test ./auth/tests -cover
