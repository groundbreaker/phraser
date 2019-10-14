.PHONY := default init test release
.DEFAULT_GOAL = default

default:
	@ mmake help

# init project
init:
	@ go mod vendor

# run tests
test:
	@ go test ./...

# create release VERSION on github
#
# VERSION should being with a v and be in SemVer format.
release: test
	$(eval VERSION=$(filter-out $@, $(MAKECMDGOALS)))
	$(if ${VERSION},@true,$(error "VERSION is required"))
	git commit --allow-empty -am ${VERSION}
	git push
	hub release create -m ${VERSION} -e ${VERSION}

%:
	@
