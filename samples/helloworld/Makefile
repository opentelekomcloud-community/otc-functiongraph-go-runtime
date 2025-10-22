build:
	go build -o target/helloworld helloworld.go

zip:
	zip -j deploy.zip target/helloworld

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip
