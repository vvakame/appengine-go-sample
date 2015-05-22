test:
	goapp test ./...

coverage:
	goapp test -coverprofile=coverage.cov sample
	go tool cover -html=coverage.cov -o=coverage.html
