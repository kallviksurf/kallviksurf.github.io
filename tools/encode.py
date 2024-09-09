import sys
from html.entities import codepoint2name


def encode(text):
  return "".join(f"&{codepoint2name[cp]};" if cp in codepoint2name else f"&#{cp};" for cp in [ord(c) for c in text])

def main():
    print(encode(sys.argv[1]))

if __name__ == '__main__':
    main()
