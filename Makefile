GO=go

build:
	${GO} build -v ./...

test:
	${GO} test -race -v ./...
watch-test:
	reflex -t 50ms -s -- sh -c 'gotest -race -v ./...'

bench:
	${GO} test -benchmem -count 3 -bench ./...
watch-bench:
	reflex -t 50ms -s -- sh -c 'go test -benchmem -count 3 -bench ./...'

coverage:
	${GO} test -v -coverprofile=cover.out -covermode=atomic .
	${GO} tool cover -html=cover.out -o cover.html

tools:
	${GO} install github.com/cespare/reflex@latest
	${GO} install github.com/rakyll/gotest@latest
	${GO} install github.com/psampaz/go-mod-outdated@latest
	${GO} install github.com/jondot/goweight@latest
	${GO} install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	${GO} get -t -u golang.org/x/tools/cmd/cover
	${GO} install github.com/sonatype-nexus-community/nancy@latest
	${GO} mod tidy

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...
lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

audit:
	${GO} list -json -m all | nancy sleuth

outdated:
	${GO} list -u -m -json all | go-mod-outdated -update -direct

weight:
	goweight