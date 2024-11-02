# ccwc

'wc' unix command from scratch in go.

```
Usage: ./ccwc [OPTIONS] [FILE]

Options:
  -c    to count bytes
  -l    to count lines
  -m    to count words

If no flags are specified, defaults to counting lines, words, and bytes.
If no file is specified, reads from standard input.
```

## Usage

```bash
$ ./ccwc sample.txt
7143 58164 342143 sample.txt
$ cat sample.txt | ./ccwc
7143 58164 342143
```

>

## How to run?

```bash
$ git clone https://github.com/mshra/ccwc
$ cd ccwc
$ go build .
$ ./ccwc sample.txt
7143 58164 342143 sample.txt
```

---

Raise a pull request or an issue if you see a bug/refactor.
