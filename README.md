# CW-Hydra
## Ory Hydra client for coldwire's purposes

```sh
go get codeberg.org/coldwire/cwhydra
```

## Exemple
```go
// Initialize API
api := cwhydra.New(cwhydra.Config{
  Url: "http://127.0.0.1:4445", // Admin Endpoint
  Http: http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Timeout: 0,
	},
})

log.Println("Creating a client named 'test'")
create, err := cwhydra.ClientManager(api).Create(cwhydra.OAuth2Client{
	ClientName: "test",
})
if err != nil {
	t.Error(err)
}

log.Println("> The client id is:", create.ClientId)

log.Println("Getting the created client")
g, err := cwhydra.ClientManager(api).Get(create.ClientId)
if err != nil {
	t.Error(err)
}

log.Println(">", "Name of the client:", g.ClientName)

log.Println("List all clients")
clients, err := cwhydra.ClientManager(api).List()
if err != nil {
	t.Error(err)
}

for i, v := range clients {
	log.Println(">", i, v.ClientId)
}

log.Println("Deleting client:", create.ClientId)
err = cwhydra.ClientManager(api).Delete(create.ClientId)
if err != nil {
	t.Error(err)
}
```