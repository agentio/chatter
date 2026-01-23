all:
	go install ./...

xrpc:
	slink generate xrpc -l debug
