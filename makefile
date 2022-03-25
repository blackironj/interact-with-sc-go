.PHONY: clean
clean:
	rm -rf build abigenBindings

.PHONY: build
build:
	make clean

	yarn truffle compile

	mkdir abigenBindings
	
	yarn truffle run abigen MondayHaterToken

	abigen --bin=abigenBindings/bin/MondayHaterToken.bin \
	--abi=abigenBindings/abi/MondayHaterToken.abi \
	--pkg=token --out=token/token.go
	
	abigen --abi=abigenBindings/abi/MondayHaterToken.abi --pkg=token --out=token/token.go
