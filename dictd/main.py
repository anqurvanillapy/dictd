from . import server

import sys

if __name__ == "__main__":
    args = sys.argv[1:]
    if len(args) != 2:
        sys.exit("Usage: dictd HOST PORT")
    host, port = sys.argv[1:]
    server.run(host, port)
