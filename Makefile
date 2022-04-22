test: sca testOnly

build: sca testOnly buildOnly

sca: 
	@sh -c "'$(CURDIR)/scripts/sca.sh'"

buildOnly:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

testOnly:
	@sh -c "'$(CURDIR)/scripts/test.sh'"

clean:
	@sh -c "'$(CURDIR)/scripts/clean.sh'"

docker-build: 
	@sh -c "'$(CURDIR)/scripts/docker-build.sh'"

docker-release: 
	@sh -c "'$(CURDIR)/scripts/docker-release.sh'"

docker-clean-all:
	@sh -c "'$(CURDIR)/scripts/docker-clean-all.sh'"

docker-clean:
	@sh -c "'$(CURDIR)/scripts/docker-clean.sh'"

# GIT_SHA1 = $(shell git rev-parse --verify HEAD)
# IMAGES_TAG = ${shell git describe --exact-match --tags 2> /dev/null || echo 'latest'}
# IMAGE_PREFIX = my-super-awesome-monorepo-

IMAGE_DIRS = $(wildcard ./components/*)

# Build all images
build-all: ${IMAGE_DIRS}

# Build and tag a single image
${IMAGE_DIRS}:	
	$(eval COMPONENT_FOLDER = $(@))
	$(eval COMPONENT_NAME = $(shell basename ${COMPONENT_FOLDER}))
	@echo $(COMPONENT_NAME)


.PHONY: all ${IMAGE_DIRS} sca build buildOnly test testOnly clean docker-build docker-release