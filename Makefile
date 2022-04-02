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

.NOTPARALLEL:

.PHONY: sca build buildOnly test testOnly clean