# tf - CLI Unix Time formatter

`tf` scans for UNIX epoch times in input and outputs them
as human readable strings to stdout.

```
$ cat log.txt
[1524241219] Time is on my side
[1555777220] Yes it is
[1587399621] A time in epoch millis: 1637419929123
[1618935621] A time in epoch nanos: 1637419929123456789

$ tf -g log.txt
[12:20:19] Time is on my side
[12:20:20] Yes it is
[12:20:21] A time in epoch millis: 09:52:09.123
[12:20:21] A time in epoch nanos: 09:52:09.123456789
```

## Usage

```
$ tf --help
usage:  tf <options> [file1 [file2 [...]]]

UNIX Time Formatter (tf)

Scans for UNIX epoch times in input and outputs them
as human readable strings to stdout.

10-digits are interpreted as seconds, 13 as milliseconds,
16 as microseconds, and 19 as nanoseconds.

If no filenames or only '-' is passed, stdin is processed.

examples:
  $ echo 1637421447 | tf
  $ tf -g log.txt | head

The time formatting uses Golang Time.Format layouts:
  https://pkg.go.dev/time#Time.Format

options:
  -b, --block               use block buffering (default: line buffering)
  -z, --block-size uint32   block buffer size (default 4096)
  -d, --date                output with date, same as --format '2006-01-02 15:04:05'
  -f, --format string       output with Golang Time.Format layout (default: '15:04:05')
  -g, --global              global match (default: convert only first match in line)
  -h, --help                show help
```

----

## Installing

Binaries for multiple platforms are [released on GitHub](https://github.com/neomantra/tf/releases) through [GitHub Actions](https://github.com/neomantra/tf/actions).

You can also install for various platforms with [Homebrew](https://brew.sh) from [`neomantra/homebrew-tap`](https://github.com/neomantra/homebrew-tap):

```
brew tap neomantra/homebrew-tap
brew install tf
```

----

## Example Usage

Raw log:
```
$ cat log.txt
[1524241219] Time is on my side
[1555777220] Yes it is
[1587399621] A time in epoch millis: 1637419929123
[1618935621] A time in epoch nanos: 1637419929123456789
```

Basic usage, piping from `stdin`:
```
$ cat log.txt | tf
[12:20:19] Time is on my side
[12:20:20] Yes it is
[12:20:21] A time in epoch millis: 1637419929123
[12:20:21] A time in epoch nanos: 1637419929123456789
```

Date conversion:
```
$ tf -d log.txt
[2018-04-20 12:20:19] Time is on my side
[2019-04-20 12:20:20] Yes it is
[2020-04-20 12:20:21] A time in epoch millis: 1637419929123
[2021-04-20 12:20:21] A time in epoch nanos: 1637419929123456789
````

Global match, converting sub-second times:
```
$ tf -gd log.txt
[2018-04-20 12:20:19] Time is on my side
[2019-04-20 12:20:20] Yes it is
[2020-04-20 12:20:21] A time in epoch millis: 2021-11-20 09:52:09.123
[2021-04-20 12:20:21] A time in epoch nanos: 2021-11-20 09:52:09.123456789
```
----

## Building

Building is performed with [task](https://taskfile.dev/#/):

```
$ task
task: [test] go test neomantra/tf/internal/tf
ok      neomantra/tf/internal/tf
task: [build] go build -o tf cmd/tf/main.go
```

----

## Credits and License

Copyright (c) 2021 Neomantra BV

Released under the [MIT License](https://en.wikipedia.org/wiki/MIT_License), see [LICENSE.txt](./LICENSE.txt).
