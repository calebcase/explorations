# FIN-WAIT-2

Client Side: Socket stuck in FIN-WAIT-2
Server Side: Socket stuck in CLOSE-WAIT

Start the server:

```
cd server
go build && ./server
```

Start the client:

```
cd client
go build && ./client
```

Run `ss` to find out what state the sockets are in:

```
sudo ss -tnp | grep ':8080'
```

Should see something like:

```
ESTAB      0      0               127.0.0.1:46164          127.0.0.1:8080  users:(("client",pid=6962,fd=3))
ESTAB      7      0      [::ffff:127.0.0.1]:8080  [::ffff:127.0.0.1]:46164 users:(("server",pid=6835,fd=4))
```

Wait 2 seconds...

Client will send itself the terminate signal (as if you'd `ctrl-c`'d it) and
exit.

Run `ss` to find out what state the sockets are in:

```
sudo ss -tnp | grep ':8080'
```

Should see something like:

```
FIN-WAIT-2 0      0               127.0.0.1:46164          127.0.0.1:8080
CLOSE-WAIT 8      0      [::ffff:127.0.0.1]:8080  [::ffff:127.0.0.1]:46164 users:(("server",pid=6835,fd=4))
```
