# FIN-WAIT-1

Client Side: Socket stuck in FIN-WAIT-1

Server Side: Socket remains in ESTAB

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
ESTAB      0      966798          127.0.0.1:46394          127.0.0.1:8080  users:(("client",pid=16528,fd=3))
ESTAB      81785  0      [::ffff:127.0.0.1]:8080  [::ffff:127.0.0.1]:46394 users:(("server",pid=16407,fd=4))
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
FIN-WAIT-1 0      966799          127.0.0.1:46394          127.0.0.1:8080
ESTAB      81785  0      [::ffff:127.0.0.1]:8080  [::ffff:127.0.0.1]:46394 users:(("server",pid=16407,fd=4))
```
