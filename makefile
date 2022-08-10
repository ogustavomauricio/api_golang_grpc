build:
	go build -o ./bin/currency-coin/server 
	
server:
	./bin/currency-coin/server

test:
	cd src/server && go test

mod:
	go mod tidy