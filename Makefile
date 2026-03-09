all:	xrpc chatter

chatter:
	go install ./...

xrpc:
	go install github.com/agentio/slink/cmd/slink-generate@v0.1.6
	slink-generate xrpc -m xrpc.json -l debug

submodules:
	git submodule update --init --recursive
