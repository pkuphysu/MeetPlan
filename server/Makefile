MODULE=github.com/pkuphysu/meetplan

.PHONY: api
api-init:
	mkdir -p cmd/api
	cd cmd/api && hz new -force -module ${MODULE}/cmd/api -I ../../../idl -idl ../../../idl/http/meetplan/service/api.proto && rm go.mod

.PHONY: gen-all
gen-all: gen-gorm gen-kitex gen-hertz

.PHONY: gen-gorm
gen-gorm:
	sleep 10
	go run cmd/gormgen/main.go

.PHONY: gen-kitex
gen-kitex: gen-kitex-user gen-kitex-plan

.PHONY: gen-kitex-user
gen-kitex-user:
	kitex --thrift-plugin validator -module github.com/pkuphysu/meetplan ./idl/user.thrift
	if [ ! -d "./cmd/user" ]; then mkdir cmd/user; fi
	cd cmd/user && kitex -service user -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../idl/user.thrift

.PHONY: gen-kitex-plan
gen-kitex-plan:
	kitex --thrift-plugin validator -module github.com/pkuphysu/meetplan ./idl/plan.thrift
	if [ ! -d "./cmd/plan" ]; then mkdir cmd/plan; fi
	cd cmd/plan && kitex -service plan -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../idl/plan.thrift

.PHONY: gen-hertz
gen-hertz:
	cd cmd/api && hz update -idl ../../../idl/api.proto

.PHONY: docker-reup
docker-reup:
	docker-compose down && docker-compose up -d