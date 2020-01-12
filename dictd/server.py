import asyncio
import logging as L

from . import db

__all__ = ("serve",)

L.basicConfig(format="[%(levelname)s] %(asctime)s %(message)s", level=L.INFO)


async def _handle_conn(reader, writer):
    addr = writer.get_extra_info("peername")

    while True:
        data = await reader.read(100)
        if not data:
            writer.close()
            L.info(f"{addr}: Closing ...")
            break

        msg = data.decode()

        resp = None

        cmd = db.parse_msg(msg)

        if cmd is None:
            L.warning(f"{addr}: Invalid command {data}")
            resp = "bad"
        else:
            L.info(f"{addr}: Command '{cmd}'")
            resp = cmd.execute()

        if isinstance(cmd, db.All):
            L.info(f"{addr}: Dumping database ...")
        else:
            L.info(f"{addr}: Sending '{repr(resp)}' ...")

        writer.write(resp.encode())
        await writer.drain()

    writer.close()


def serve(port):
    loop = asyncio.get_event_loop()
    coro = asyncio.start_server(_handle_conn, '0.0.0.0', port, loop=loop)
    server = loop.run_until_complete(coro)

    L.info(f"Listening to {server.sockets[0].getsockname()} ...")
    try:
        loop.run_forever()
    except KeyboardInterrupt:
        pass

    server.close()
    loop.run_until_complete(server.wait_closed())
    loop.close()


if __name__ == "__main__":
    import sys

    _, *args = sys.argv
    if len(args) != 1:
        sys.exit("Usage: dictd.server PORT")
    serve(*args)
