.PHONY: clean
clean:
	rm -rf build abigenBindings

.PHONY: install
install:
	yarn install

.PHONY: build-contract
build-contract:
	make clean

	yarn truffle compile

	mkdir abigenBindings
	
	yarn truffle run abigen MondayHaterToken

	abigen --bin=abigenBindings/bin/MondayHaterToken.bin \
	--abi=abigenBindings/abi/MondayHaterToken.abi \
	--pkg=api --out=api/token.go

.PHONY: build
build:
	go build -o sc-interactor cmd/main.go

.PHONY: all
all: clean install build-contract build 
