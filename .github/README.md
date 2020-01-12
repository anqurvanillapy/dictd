# Dictd

> **DISCLAIMER**: Toy project.

Dict, distributed.  A distributed KV database written in Python, for learning
to implement the Raft concensus algorithm.

- Commands/Protocols:
    + Set: `'+KEY\r\nVAL'`
    + Get: `'=KEY\r\n'`
    + Delete: `'-KEY\r\n'`
    + Dump all: `'*\r\n'`
    + Get size: `'?\r\n'`
    + Clear all: `'!\r\n'`

## License

MIT
