all:	xrpc chatter

chatter:
	go install ./...

xrpc:
	go install github.com/agentio/slink/cmd/slink-generate@v0.1.0
	slink-generate xrpc -m xrpc.json -l debug
