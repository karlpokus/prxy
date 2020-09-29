# prxy
proxy a to b. WIP

# use-case
Stubborn process running on localhost on remote host.

# usage
```bash
# args format is ip:port
$ prxy <src> <dest>
```

# test
```bash
$ go test -v
```

# todos
- [ ] debug flag
- [ ] if arg is - then use stdin or stdout so we can put tee in the middle to dump traffic
- [ ] optional port in args
