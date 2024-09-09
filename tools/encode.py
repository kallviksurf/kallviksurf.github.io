import sys


def encode(text):
  return "".join([f"&#{cp};" for cp in [ord(c) for c in text]])


print(encode(sys.argv[1]))

