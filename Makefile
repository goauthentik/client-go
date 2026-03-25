.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

build:
	docker compose -f compose.yml run --rm --user "${UID}:${GID}" gen \
		generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	rm -rf ./test
	rm -f .travis.yml git_push.sh
	go fmt .
