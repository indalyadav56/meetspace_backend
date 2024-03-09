include .env

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

postgres-start:
	docker run --name postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=indal_db -p 5432:5432 -v pgdata:/var/lib/postgresql/data -d postgres

postgres-remove:
	docker rm -f postgres

postgres-shell:
	docker exec -it postgres psql -U postgres

redis-start:
	docker run -d --name redis-container -p 6379:6379 redis:latest --requirepass redis@123

redis-remove:
	docker rm -f redis-container

redis-cli:
	docker exec -it meetspace_redis redis-cli -a $(REDIS_PASSWORD)

livekit-local:
	# docker run --rm -v$PWD:/output livekit/generate --local
	# docker run -d --rm \
    # -p 7880:7880 \
    # -p 7881:7881 \
    # -p 7882:7882/udp \
    # -v $PWD/livekit.yaml:/livekit.yaml \
    # livekit/livekit-server \
    # --config /livekit.yaml \
    # --node-ip=127.0.0.1

livekit-deploy:
	docker run --rm -it -v$PWD:/output livekit/generate
