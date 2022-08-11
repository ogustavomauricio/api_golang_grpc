build:
	go build -o ./bin/currency-coin/server 
	
server:
	./bin/currency-coin/server

test:
	go test -v

mod:
	go mod tidy