all: bin

SHORT_NAME=okta-viewer
REPOSITORY=docker.io/fleetingclarity/$(SHORT_NAME)
VERSION=latest
FULL_IMAGE=$(REPOSITORY):$(VERSION)
PLATFORM=local
CI_COMMIT_SHA ?= $(shell git rev-parse HEAD)
# supported platforms
# make PLATFORM=linux/amd64
# make PLATFORM=darwin/amd64
# make PLATFORM=windows/amd64
# make PLATFORM=linux/arm

.PHONY: test
test:
	@echo "running go tests..."
	@docker build . --target unit-test-coverage --output coverage/

.PHONY: bin
bin:
	@docker build . --target bin \
	--output bin/ \
	--platform ${PLATFORM}
	chmod +x bin/*

.PHONY: tag-and-push
tag-and-push:
	@docker tag $(FULL_IMAGE) $(REPOSITORY):$(CI_COMMIT_SHA)
	@docker push $(FULL_IMAGE)
	@docker push $(REPOSITORY):$(CI_COMMIT_SHA)

.PHONY: print-repository
print-repository:
	@echo $(REPOSITORY)

.PHONY: clean
clean:
	rm -rf bin/ coverage/

.PHONY: run-image
run-image:
	@docker run -d -p $(PORT):8080 --name $(SHORT_NAME) $(FULL_IMAGE)

.PHONY: docker-stop
docker-stop:
	@docker stop $(SHORT_NAME)

.PHONY: docker-clean
docker-clean:
	@docker rm $(SHORT_NAME)
