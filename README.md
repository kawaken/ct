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
ct [options] [modifier] [step]
```

### Options

```bash
-mod=number, -m        Specify the number of modulo.
```

### Modifier

```bash
up                     Increments the current stored number. Default modifier.
down                   Decrements the current stored number.
```

`up` is default modifier.

### Step

Step is the number for couting. Default is `1`. Step requires a integer number.

**Zero**

Not occur anything. Displayed current stored number.

**Negative number**

`up` causes **Decrement**.
`down` causes **Increment**.

## Install

```
$ go get github.com/kawaken/ct
```

## Contribution

## Licence

[MIT](https://github.com/kawaken/ct/blob/master/LICENSE)

## Author

[kawaken](https://github.com/kawaken)

