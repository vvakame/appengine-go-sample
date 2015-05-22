prepare:
	go fmt ./...
	go generate ./...

test: prepare
	goapp test ./...

coverage: prepare
	goapp test -coverprofile=coverage.cov sample
	go tool cover -html=coverage.cov -o=coverage.html

ci: coverage
	goveralls -v -service=circle-ci -coverprofile=coverage.cov
