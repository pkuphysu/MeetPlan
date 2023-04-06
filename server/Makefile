MODULE=github.com/pkuphysu/meetplan

.PHONY: api
api:
	mkdir -p cmd/api
	cd cmd/api && hz new

.PHONY: gen
gen-all: gen-gorm gen-kitex
	mkdir -p cmd/user && cd cmd/user && kitex -service user -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/user.thrift

.PHONY: gen-gorm
gen-gorm:
	go run cmd/gormgen/main.go

.PHONY: gen-kitex
gen-kitex: gen-kitex-user gen-kitex-plan gen-kitex-order

.PHONY: gen-kitex-user
gen-kitex-user:
	kitex --thrift-plugin validator -module github.com/pkuphysu/meetplan ../idl/user.thrift
	if [ ! -d "./cmd/user" ]; then mkdir cmd/user; fi
	cd cmd/user && kitex -service user -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/user.thrift

.PHONY: gen-kitex-plan
gen-kitex-plan:
	kitex --thrift-plugin validator -module github.com/pkuphysu/meetplan ../idl/plan.thrift
	if [ ! -d "./cmd/plan" ]; then mkdir cmd/plan; fi
	cd cmd/plan && kitex -service plan -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/plan.thrift

.PHONY: gen-kitex-order
gen-kitex-order:
	kitex --thrift-plugin validator -module github.com/pkuphysu/meetplan ../idl/order.thrift
	if [ ! -d "./cmd/order" ]; then mkdir cmd/order; fi
	cd cmd/order && kitex -service order -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/order.thrift
