ct
====

`ct` is simple counting tool for CLI.

## Description

`ct` counts up or counts down a number and displays it.   
The result is saved in a file for next time.

`ct` is __zero based__.

## Demo

## Requirement

* Go

## Usage

```
ct [options] [modifier] [number]
```

### Options

```bash
--mod=number, -m    Specify the number of modulo.
--file=path, -f     Specify the file to save number.
```

`--file`: Default file is `$HOME/.config/ct/count`.

### Modifier and Number

```bash
up                 Increments the current stored number. Default modifier.
down               Decrements the current stored number.
reset              Reset the number.
```

`up` is default modifier.
`1` is default number.

#### Up/Down

Number requires a integer number.

**Zero**

Not occur anything. Displayed current stored number.

**Negative number**

`up` causes **Decrement**.
`down` causes **Increment**.

#### Reset

Reset stored number to specific number.

## Install

```bash
$ go get github.com/kawaken/ct
```

## Contribution

## Licence

[MIT](https://github.com/kawaken/ct/blob/master/LICENSE)

## Author

[kawaken](https://github.com/kawaken)

