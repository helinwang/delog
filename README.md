# delog
extract the values of given keys from the log in logfmt

## Usage

```bash
$ go get github.com/helinwang/delog
$ echo 'a=foo b=10ms c=cat E="123"'| delog -k a,b
foo,10ms
```
