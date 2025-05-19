.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: clean update build diff

clean:
	rm -f *.go
	rm -rf docs/

update:
	mv schema.yml schema-old.yml
	cp ../authentik/schema.yml schema.yml

build:
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		docker.io/openapitools/openapi-generator-cli:v6.2.0 generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	rm -rf ./test
	rm -f .travis.yml git_push.sh
	go get
	go fmt .

diff:
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		docker.io/openapitools/openapi-diff:2.1.1 \
		--markdown /local/diff.test \
		/local/schema-old.yml /local/schema.yml || echo > diff.test
	rm schema-old.yml
	mv diff.test /tmp/diff.test
	echo "Update API Client\n\n" > diff.test
	cat /tmp/diff.test >> diff.test
