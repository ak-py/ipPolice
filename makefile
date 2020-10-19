

compile:
	go build ./...

test:
	go test ./...

exec:
	env GOOS=linux GOARCH=amd64 go build -o ipPolice ipAddressPolice.go

clean:
	rm -rf logs */logs