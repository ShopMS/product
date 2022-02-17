tidy:
	go mod tidy

gen: tidy
	buf mod update && buf generate
