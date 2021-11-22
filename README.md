# gotf - go time formatter

`gotf` converts epoch times in inputs to human readable strings
and outputs them to stdout.

```
$ cat log.txt
[1524241219] Time is on my side
[1555777220] Yes it is
[1587399621] A time in ecpoch millis: 1637419929123
[1618935621] A time in ecpoch nanos: 1637419929123456789

$ gotf -g log.txt
[12:20:19] Time is on my side
[12:20:20] Yes it is
[12:20:21] A time in ecpoch millis: 09:52:09.123
[12:20:21] A time in ecpoch nanos: 09:52:09.123456789
```

## Usage

```
$ gotf --help
usage:  ./gotf <options> [file1 [file2 [...]]]

GOlang Time Formatter

Reads text files, converting epoch times to human readable, outputting to stdout.

10-digits are interpreted as seconds, 13 as milliseconds,
16 as microseconds, and 19 as nanoseconds.

If no filenames or only '-' is passed, stdin is processed.

example:
echo 1637421447 | gotf

gotf -g log.txt | head

  -b, --block           use block buffering (default: line buffering)
  -d, --date            default format with '2006-01-02 15:04:05'
  -f, --format string   golang Time.Format string (default: '15:04:05')
  -g, --global          global match
  -h, --help            show help```
```

----

## Installing

Binaries for multiple platforms are [released on GitHub](https://github.com/neomantra/gotf/releases) through [GitHub Actions](https://github.com/neomantra/gotf/actions).

You can also install for various platforms with [Homebrew](https://brew.sh) from [`neomantra/homebrew-tap`](https://github.com/neomantra/homebrew-tap):

```
brew tap neomantra/homebrew-tap
brew install gotf
```

----

## Example Usage

Raw log:
```
$ cat log.txt
[1524241219] Time is on my side
[1555777220] Yes it is
[1587399621] A time in ecpoch millis: 1637419929123
[1618935621] A time in ecpoch nanos: 1637419929123456789
```

Basic usage, piping from `stdin`:
```
$ cat log.txt | gotf
[12:20:19] Time is on my side
[12:20:20] Yes it is
[12:20:21] A time in ecpoch millis: 1637419929123
[12:20:21] A time in ecpoch nanos: 1637419929123456789
```

Date conversion:
```
$ gotf -d log.txt 
[2018-04-20 12:20:19] Time is on my side
[2019-04-20 12:20:20] Yes it is
[2020-04-20 12:20:21] A time in ecpoch millis: 1637419929123
[2021-04-20 12:20:21] A time in ecpoch nanos: 1637419929123456789
````

Global match, converting sub-second times:
```
$ gotf -gd log.txt
[2018-04-20 12:20:19] Time is on my side
[2019-04-20 12:20:20] Yes it is
[2020-04-20 12:20:21] A time in ecpoch millis: 2021-11-20 09:52:09.123
[2021-04-20 12:20:21] A time in ecpoch nanos: 2021-11-20 09:52:09.123456789
```
----

## Building

Building is performed with [task](https://taskfile.dev/#/):

```
$ task
task: [test] go test neomantra/gotf/internal/gotf
ok      neomantra/gotf/internal/gotf
task: [build] go build -o gotf cmd/gotf/main.go
```

----

## Credits and License

Copyright (c) 2021 Neomantra BV

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE.txt](./LICENSE.txt).
