MODULE=meetplan

.PHONY: api
api-init:
	hz new -force -module ${MODULE} -I ../idl -idl ../idl/service/meetplan.proto --customize_package=./pkg/tpl/hertz/package.yaml
	rm -rf script
	go mod tidy

.PHONY: gen-all
gen-all: gen-gorm gen-hertz

.PHONY: gen-gorm
gen-gorm:
	go run cmd/gormgen/main.go

.PHONY: gen-hertz
gen-hertz:
	hz update -I ../idl -idl ../idl/service/meetplan.proto --customize_package=./pkg/tpl/hertz/package.yaml

.PHONY: docker-reup
docker-reup:
	docker-compose down && docker-compose up -d

lint:
	find . -depth -name '*.go' ! -name '*pb.go' ! -name 'router_gen.go' ! -path './biz/router/service/*' ! -path './biz/gorm_gen/*' -type f -print0 | xargs -0 goimports -w -local ${MODULE}