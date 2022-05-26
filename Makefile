.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

default:
	wget -O schema.yml https://raw.githubusercontent.com/goauthentik/authentik/master/schema.yml
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli:v6.0.0 generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	go get
	go fmt .
	go mod tidy
