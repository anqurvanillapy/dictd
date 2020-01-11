from enum import Enum, auto

__all__ = ("parse_msg", "Set", "Get", "Del", "All", "Clr", "Siz")

_DB = {}


def _set(key, val):
    _DB[key] = val
    return repr(key)


def _get(key):
    try:
        return repr(_DB[key])
    except KeyError:
        return "bad"


def _del(key):
    try:
        del _DB[key]
        return "ok"
    except KeyError:
        return "bad"


def parse_msg(msg):
    msg = msg.strip()
    if len(msg) == 0:
        return None

    cmd, txt = msg[0], msg[1:]

    if len(msg) == 1:
        if cmd == "*":
            return All()
        elif cmd == "!":
            return Clr()
        elif cmd == "?":
            return Siz()

    if cmd == "+":
        k, _, v = txt.partition("\r\n")
        if k == "" or v == "":
            return None
        return Set(k, v)

    k = txt
    if k == "":
        return None

    if cmd == "=":
        return Get(k)
    elif cmd == "-":
        return Del(k)

    return None


class Set:
    def __init__(self, key, val):
        self.key = key
        self.val = val

    def __str__(self):
        return f"Set(key={repr(self.key)}, val={repr(self.val)})"

    def execute(self):
        return _set(self.key, self.val)


class Get:
    def __init__(self, key):
        self.key = key

    def __str__(self):
        return f"Get(key={repr(self.key)})"

    def execute(self):
        return _get(self.key)


class Del:
    def __init__(self, key):
        self.key = key

    def __str__(self):
        return f"Del(key={repr(self.key)})"

    def execute(self):
        return _del(self.key)


class All:
    def __str__(self):
        return "All"

    def execute(self):
        return str(_DB)


class Clr:
    def __str__(self):
        return "Clr"

    def execute(self):
        _DB.clear()
        return "ok"


class Siz:
    def __str__(self):
        return "Siz"

    def execute(self):
        return str(len(_DB))
