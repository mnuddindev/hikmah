.PHONY: all asas run-ufuq fmt lint clean

all: asas

asas:
	go build -o bin/asasd ./asas/cmd/asasd

run-asas: asas
	./bin/asasd

run-ufuq:
	qs -p ufuq/

fmt:
	gofmt -w ./asas
	qmllint ufuq/*.qml hikmahui/*.qml || true

lint:
	go vet ./asas/...

clean:
	rm -rf bin