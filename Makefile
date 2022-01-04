sca:
	@sh -c "'$(CURDIR)/scripts/sca.sh'"

static-code-analysis:
	@sh -c "'$(CURDIR)/scripts/sca.sh'"

build: sca
	@sh -c "'$(CURDIR)/scripts/build.sh'"

.NOTPARALLEL:

.PHONY: sca static-code-analysis build