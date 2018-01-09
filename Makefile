
build:
	dep ensure
	go build .

run: build
	./goreg ls

clean:
	rm ./goreg
