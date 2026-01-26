.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

all: clean update diff build

clean:
	rm -f *.go
	rm -rf docs/

update:
	mv schema.yml schema-old.yml
	cp ../authentik/schema.yml schema.yml

diff:
	docker compose -f scripts/compose.yml run --rm --user "${UID}:${GID}" diff \
		--markdown \
		/local/diff.test \
		/local/schema-old.yml \
		/local/schema.yml
	rm schema-old.yml
	mv diff.test /tmp/diff.test
	echo "Update API Client\n\n" > diff.test
	echo >> diff.test
	echo >> diff.test
	cat /tmp/diff.test >> diff.test

build:
ifndef version
	$(error Usage: make build version=20xx.xx.xx)
endif
	docker compose -f scripts/compose.yml run --rm --user "${UID}:${GID}" gen \
		generate \
		-i /local/schema.yml \
		-g go \
		-o /local \
		-c /local/config.yaml
	rm -rf ./test
	rm -f .travis.yml git_push.sh
	git checkout -- go.mod go.sum
	go get
	go fmt .
