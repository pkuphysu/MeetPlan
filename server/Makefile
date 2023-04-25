MODULE=meetplan

.PHONY: api
api-init:
	mkdir -p cmd/api
	cd cmd/api && hz new -force -module ${MODULE}/cmd/api -I ../../../idl -idl ../../../idl/http/service/meetplan.proto --customize_package=../../tpl/hertz/package.yaml && rm go.mod

.PHONY: gen-all
gen-all: gen-gorm gen-kitex gen-hertz

.PHONY: gen-gorm
gen-gorm:
	go run cmd/gormgen/main.go

.PHONY: gen-kitex
gen-kitex: gen-kitex-user gen-kitex-plan

.PHONY: gen-kitex-user
gen-kitex-user:
	kitex --thrift-plugin validator -module meetplan ./idl/user.thrift
	if [ ! -d "./cmd/user" ]; then mkdir cmd/user; fi
	cd cmd/user && kitex -service user -module ${MODULE} -use meetplan/kitex_gen/  ../../idl/user.thrift

.PHONY: gen-kitex-plan
gen-kitex-plan:
	kitex --thrift-plugin validator -module meetplan ./idl/plan.thrift
	if [ ! -d "./cmd/plan" ]; then mkdir cmd/plan; fi
	cd cmd/plan && kitex -service plan -module ${MODULE} -use meetplan/kitex_gen/  ../../idl/plan.thrift

.PHONY: gen-hertz
gen-hertz:
	cd cmd/api && hz update -I ../../../idl -idl ../../../idl/http/service/meetplan.proto --customize_package=../../tpl/hertz/package.yaml

.PHONY: docker-reup
docker-reup:
	docker-compose down && docker-compose up -d