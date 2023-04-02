MODULE=github.com/pkuphysu/meetplan

.PHONY: api
api:
	mkdir -p cmd/api
	cd cmd/api && hz new

.PHONY: user
user:
	mkdir -p cmd/user
	cd cmd/user && \
	kitex -service user -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/user.thrift

.PHONY: gen
gen:
	bash ./kitex-all.sh
	mkdir -p cmd/user && cd cmd/user && kitex -service user -module ${MODULE} -use github.com/pkuphysu/meetplan/kitex_gen/  ../../../idl/user.thrift
	go run cmd/gormgen/main.go
