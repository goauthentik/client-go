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
		-o /local \
		--additional-properties=packageName=api,enumClassPrefix=true,useOneOfDiscriminatorLookup=true
	find api/ -type f -exec \
		sed -i 's/goauthentik\.io\/GIT_USER_ID\/GIT_REPO_ID/goauthentik\.io\/api/g' {} +
	git checkout -- README.md
	git checkout -- go.mod
