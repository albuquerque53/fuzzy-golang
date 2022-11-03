run:
	go run main.go
test:
	go test ./... -v
fuzz:
	sh run_all_fuzz.sh