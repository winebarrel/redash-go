.PHONY: all
all: vet test

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: testacc
testacc:
	$(MAKE) test TEST_ACC=1

.PHONY: lint
lint:
	golangci-lint run

.PHONY: redash-setup
redash-setup:
	psql -U postgres -h localhost -p 15432 -f _etc/redash.sql
