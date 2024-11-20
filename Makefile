.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: clean build

clean:
	rm -f *.go
	rm -rf docs/

build:
	mv schema.yml schema-old.yml
	wget -O schema.yml https://raw.githubusercontent.com/goauthentik/authentik/main/schema.yml
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		docker.io/openapitools/openapi-diff:2.1.0-beta.11 \
		--markdown /local/diff.test \
		/local/schema-old.yml /local/schema.yml || echo > diff.test
	rm schema-old.yml
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		docker.io/openapitools/openapi-generator-cli:v7.10.0 generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	rm -rf ${PWD}/test ${PWD}/.travis.yml ${PWD}/git_push.sh
	go run golang.org/x/tools/cmd/goimports@latest -w .
	go get
	go fmt .
	mv ${PWD}/diff.test /tmp/diff.test
	echo "Update API Client\n\n" > ${PWD}/diff.test
	cat /tmp/diff.test >> ${PWD}/diff.test
