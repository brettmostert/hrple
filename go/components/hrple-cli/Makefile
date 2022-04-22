test: sca testOnly

build: sca testOnly buildOnly

docker-build: 
	@sh -c "'$(CURDIR)/scripts/docker-build.sh'"

buildOnly:
	@sh -c "'$(CURDIR)/scripts/build.sh'"

docker-release: 
	@sh -c "'$(CURDIR)/scripts/docker-release.sh'"

# Global Scripts / Commands below for easy access

testOnly:
	@sh -c "'../../scripts/test.sh'"

sca: 
	@sh -c "'../../scripts/sca.sh'"

clean:
	@sh -c "'../../scripts/clean.sh'"

docker-clean-all:
	@sh -c "'../../scripts/docker-clean-all.sh'"

docker-clean:
	@sh -c "'../../scripts/docker-clean.sh'"

.PHONY: sca build buildOnly test testOnly clean docker-build docker-release