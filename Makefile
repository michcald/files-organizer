SHELL:=/bin/bash

export SERVICE_NAME:=files-organizer
export NETWORK:=$(SERVICE_NAME)-net
export MINIMUM_TEST_COVERAGE:=90
export GO_FMT_IMAGE=$(SERVICE_NAME)-go-fmt

default: clean prepare fmt analyze test coverage clean

prepare:
	docker network create $(NETWORK)

clean:
	docker-compose -f ./.docker/compose.yml rm --force
	rm -rf coverage.out
	docker network rm $(NETWORK) || true

fmt: # Formats the source code
	docker build -t $(GO_FMT_IMAGE) --target golang .docker/image/app
	docker run --rm \
		-v $(PWD):/app \
		-w /app \
		$(GO_FMT_IMAGE) sh -c "go mod tidy; go fmt ./..."

test:
	docker-compose -f ./.docker/compose.yml up --build --force-recreate --abort-on-container-exit

coverage: # Requires tests to run first
	docker run --rm \
		-e MINIMUM_COVERAGE=${MINIMUM_TEST_COVERAGE} \
		-e README_FILE=/app/README.md \
		-e COVERAGE_FILE=/app/coverage.out \
		-v $(PWD):/app \
		-w /app \
		michcald/go-coverage

analyze:
	docker run --rm -v $(PWD):/app michcald/go-analyzer
