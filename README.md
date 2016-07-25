ct
====

`ct` is simple counting tool for CLI.

## Description

`ct` counts up or counts down a number and displays it. The result is saved in a file for next time.

## Demo

## Requirement

* Go

## Usage

```
ct [options] [modifier]
```

### Options

```bash
-step=number, -s       Specify the number of counting step.
-mod=number, -m        Specify the number of modulo.
```

### Modifier

```bash
up                     Increments the current stored number. Default modifier.
down                   Decrements the current stored number.
```

`up` is default modifier.

## Install

```
$ go get github.com/kawaken/ct
```

## Contribution

## Licence

[MIT](https://github.com/kawaken/ct/blob/master/LICENSE)

## Author

[kawaken](https://github.com/kawaken)

