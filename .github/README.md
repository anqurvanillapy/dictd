# Dictd

> **DISCLAIMER**: Toy project.

Dict, distributed.  A distributed KV database written in Python, for learning
to implement the Raft concensus algorithm.

- Commands/Protocols:
    + Set: `'+KEY\r\nVAL\r\n'`
    + Get: `'=KEY\r\n'`
    + Delete: `'-KEY\r\n'`
    + Dump all: `'*\r\n'`
    + Get size: `'?\r\n'`
    + Clear all: `'!\r\n'`

Start the server:

```bash
$ python -m dictd.server localhost 8080
```

Easiest way to send client command:

```bash
$ echo -ne "+foo\r\nbar\r\n" | nc localhost 8080
```

Start the client:

```bash
$ python -m dictd.client localhost 8080
```

## License

MIT
