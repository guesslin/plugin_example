## How to build this example

- run `make all` to generate all __binary__ and __plugins__
- run `make clean` to clean all __binary__ and __plugins__


## How to run this example


### Usage

```
$ ./example -h
Usage of ./example:
  -mode string
        load which kind [adder|greet] of plugin (default "adder")
  -path string
        path to plugin file
```

### Load Greeting

```
$ ./example -mode greet -path plugs/greeter/greeter.so
======= main =======
Greeting Mr. guesslin
```

### Load different Adder

- simple adder, implement `a + b`
```
$ ./example -mode adder -path plugs/adder/simple/simple.so
======= main =======
Adder 100 + 20 = 120 result
```

- mod adder, implement `(a + b) * 10 * alpha(0.5)`
```
$ ./example -mode adder -path plugs/adder/mod/mod.so
======= main =======
Adder 100 + 20 = 600 result
```

- average adder, implement `(a + b) / 2`
```
$ ./example -mode adder -path plugs/adder/average/average.so
======= main =======
Adder 100 + 20 = 60 result
```
