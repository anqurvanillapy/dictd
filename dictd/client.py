import sys
import asyncio

__all__ = ("dial",)

_HELP = """
Commands:

    set KEY VALUE       Sey key-value pair
    get KEY             Get value by key
    del KEY             Delete value by key
    all                 Dump database
    clr                 Clear database
    siz                 Get number of pairs
"""

_CMD_INFO = {
    "set": {"argc": 2, "prefix": "+"},
    "get": {"argc": 1, "prefix": "="},
    "del": {"argc": 1, "prefix": "-"},
    "all": {"argc": 0, "prefix": "*"},
    "clr": {"argc": 0, "prefix": "!"},
    "siz": {"argc": 0, "prefix": "?"},
}


def _parse_cmd(cmd):
    cmd = cmd.split()

    if not cmd:
        return None

    cmd, *args = cmd

    try:
        info = _CMD_INFO[cmd.lower()]
        if len(args) != info["argc"]:
            return None
        return "".join([info["prefix"], "\r\n".join(args), "\r\n"])
    except KeyError:
        pass

    return None


def _show_help():
    print(_HELP, file=sys.stderr)


async def _handle_conn(host, port, loop):
    reader, writer = await asyncio.open_connection(host, port, loop=loop)

    while True:
        try:
            cmd = input("dictd> ").strip()

            if not cmd:
                continue
            elif cmd.lower() == "help":
                _show_help()
                continue

            msg = _parse_cmd(cmd)

            if msg is None:
                print("error: Invalid command", file=sys.stderr)
                _show_help()
                continue

            writer.write(msg.encode())

            data = await reader.read(100)
            print(data.decode())

        except EOFError:
            print("Bye")
            writer.close()
            await writer.wait_closed()
            break


def dial(host, port):
    loop = asyncio.get_event_loop()
    loop.run_until_complete(_handle_conn(host, port, loop))
    loop.close()


if __name__ == "__main__":
    _, *args = sys.argv
    if len(args) != 2:
        sys.exit("Usage: dictd.client HOST PORT")
    dial(*args)
