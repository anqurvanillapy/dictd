"""DB server"""

import asyncio
import logging as L

L.getLogger().setLevel(L.INFO)

from . import db

__all__ = ("run",)


async def _handle_conn(reader, writer):
    data = await reader.read(100)
    msg = data.decode()

    resp = None

    addr = writer.get_extra_info("peername")
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
        L.info(f"{addr}: Sending '{resp}' ...")

    writer.write(resp.encode())
    await writer.drain()

    writer.close()


def run(host, port):
    loop = asyncio.get_event_loop()
    coro = asyncio.start_server(_handle_conn, host, port, loop=loop)
    server = loop.run_until_complete(coro)

    L.info(f"Listening to {host}:{port} ...")
    try:
        loop.run_forever()
    except KeyboardInterrupt:
        pass

    server.close()
    loop.run_until_complete(server.wait_closed())
    loop.close()
