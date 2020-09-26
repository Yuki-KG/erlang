# erlang

## Files
This repository includes the following files:
- erlB_int.c
- erlB_int.py
- erlB_int.go
- erlB_cont.c
- erlB_cont.py
- erlB_cont.go
- dererl.c
- dererl.py
- dererl.go

## Usage
### erlB_int
On a terminal shell, type either of the following command depending on your system environment. Make sure you need Python 3.x to run .py files.

```bash
$ gcc erlB_int.c; ./a.out
$ python erlB_int.py
$ go run erlB_int.go
```

Then give parameters (s as integer number of servers, a as offered traffic in erlang) and you'll have Es(a).

### erlB_cont
On a terminal shell, type either of the following command depending on your system environment. Make sure you need Python 3.x to run .py files.

```bash
$ gcc erlB_cont.c; ./a.out
$ python erlB_cont.py
$ go run erlB_cont.go
```

Then give parameters (s as number of servers, a as offered traffic in erlang) and you'll have Es(a).

### dererl
On a terminal shell, type either of the following command depending on your system environment. Make sure you need Python 3.x to run .py files.

```bash
$ gcc dererl.c; ./a.out
$ python dererl.py
$ go run dererl.go
```

Then give parameters (s as integer number of servers, a as offered traffic in erlang) and you'll have psi, Es(a), dEs/da, dEs/ds and ds/da.
