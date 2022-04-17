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

.NOTPARALLEL:

.PHONY: sca build buildOnly test testOnly clean docker-build docker-release