# chatter

This is a small demo that uses [slink](https://github.com/agentio/slink)-generated handlers to call XRPC methods.

`chatter` uses the Bluesky Chat API to send messages from one user to another.
It requires a password for the sending user (please use an [App Password](https://bsky.app/settings/app-passwords)),
and the receiving user will probably need to be following the sending user.

Here's an example. `CHATPW` is an app password of the sending user. Note that here we are doing some verbose logging. The default log level is `warn` and is much quieter.
```
$ chatter --from timburks.me --to agent.io --password $CHATPW --text "Hey! How are you?" -l info
2026/01/22 16:29:41 INFO Found DID with DNS
2026/01/22 16:29:42 INFO POST https://inkcap.us-east.host.bsky.network/xrpc/com.atproto.server.createSession
2026/01/22 16:29:43 INFO 200 (1484 bytes)
2026/01/22 16:29:43 INFO eyJ0eX...
2026/01/22 16:29:43 INFO Found DID with DNS
2026/01/22 16:29:43 INFO did:plc:ahr5yhciwadehhwm7fotyfju did:plc:bnr33h7nafe5nk4zzlshvana https://inkcap.us-east.host.bsky.network
2026/01/22 16:29:43 INFO GET https://inkcap.us-east.host.bsky.network/xrpc/chat.bsky.convo.getConvoForMembers?members=did%3Aplc%3Aahr5yhciwadehhwm7fotyfju&members=did%3Aplc%3Abnr33h7nafe5nk4zzlshvana
2026/01/22 16:29:43 INFO authorization: Bearer eyJ0eX...
2026/01/22 16:29:43 INFO atproto-proxy: did:web:api.bsky.chat#bsky_chat
2026/01/22 16:29:44 INFO 200 (8510 bytes)
2026/01/22 16:29:44 INFO convo id 3md2h2kakjk22
2026/01/22 16:29:44 INFO POST https://inkcap.us-east.host.bsky.network/xrpc/chat.bsky.convo.sendMessage
2026/01/22 16:29:44 INFO authorization: Bearer eyJ0eX...
2026/01/22 16:29:44 INFO atproto-proxy: did:web:api.bsky.chat#bsky_chat
2026/01/22 16:29:45 INFO 200 (174 bytes)
2026/01/22 16:29:45 INFO &{LexiconTypeID: Embed:<nil> Facets:[] Id:3md2h2l7tit2p Reactions:[] Rev:22222247kqldr Sender:0xc00050a0c0 SentAt:2026-01-23T00:29:45.751Z Text:Hey! How are you?}
```

Generated support code is produced by the `make xrpc` target in the Makefile. This code is generated with `slink` and uses the `lexicons` directory, which is included as a submodule. 

Everything here is released under the [AGPL](LICENSE).

If you try this, let me know, but please don't depend on it yet! `slink` is prerelease and I'm thinking about making a few more changes to the generated code.

-- Tim Burks [@timburks.me](https://bsky.app/profile/timburks.me) | [@agent.io](https://bsky.app/profile/agent.io) 
