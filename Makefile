binary:
	go build -o main main.go

test:
	go test -timeout 30s -count=1 -cover 