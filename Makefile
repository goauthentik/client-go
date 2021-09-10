.SHELLFLAGS += -x -e
PWD = $(shell pwd)
UID = $(shell id -u)
GID = $(shell id -g)

default:
	docker run \
		--rm -v ${PWD}:/local \
		--user ${UID}:${GID} \
		openapitools/openapi-generator-cli generate \
		--git-host goauthentik.io \
		--git-repo-id "" \
		--git-user-id "" \
		-i /local/schema.yml \
		-g go \
		-o /local/api \
		--additional-properties=packageName=api,enumClassPrefix=true,useOneOfDiscriminatorLookup=true
