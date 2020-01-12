# Dictd

> **DISCLAIMER**: Toy project.

Dict, distributed.  A distributed KV database written in Python, for learning
to implement the Raft concensus algorithm.

- Commands/protocols:
    + Set key-value pair: `set KEY VAL`/`'+KEY\r\nVAL\r\n'`
    + Get value by key: `get KEY`/`'=KEY\r\n'`
    + Delete value by key: `del KEY`/`'-KEY\r\n'`
    + Dump database: `all`/`'*\r\n'`
    + Clear database: `clr`/`'!\r\n'`
    + Get number of pairs: `siz`/`'?\r\n'`

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
