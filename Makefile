CURRENT_DIR=$(shell pwd)
DBURL := postgres://postgres:03212164@localhost:5432/auth_commers_service?sslmode=disable

proto-gen:
	./scripts/gen-proto.sh ${CURRENT_DIR}


mig-up:
	migrate -path migrations -database '${DBURL}' -verbose up

mig-down:
	migrate -path migrations -database '${DBURL}' -verbose down

mig-force:
	migrate -path migrations -database '${DBURL}' -verbose force 1

create_migrate:
	@echo "Enter file name: "; \
	read filename; \
	migrate create -ext sql -dir migrations -seq $$filename
swag:
	~/go/bin/swag init -g ./api/router.go -o api/docs
run-service:
	go run cmd/main.go
