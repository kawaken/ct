ct
====

`ct` is a simple counting tool for CLI.

## Description

`ct` counts up a number and displays it.   
The result is saved in a file for next time.

## Demo

**Counting**

![counting](https://kawaken.github.io/ct/ct.gif)

**Rotation**

![rotation](https://kawaken.github.io/ct/ctr.gif)

**Separate File**

![separate](https://kawaken.github.io/ct/ctf.gif)

## Requirement

Nothing. Download a binary via release page (in preparation).

## Usage

```
ct [options] [modifier] [number]
```

### Options

```bash
--rotate=number, -r    Specify the number to rotate.
--file=path,     -f    Specify the file to save number.
```

`--file`: Default file is `$HOME/.config/ct/count`.

### Modifier and Number

**Modifier:**

```bash
up       Increments the current stored number. Default modifier.
reset    Reset the number.
```

* `up` is default modifier.
* `reset` resets stored number to specific number.
 
**Number:**

Number requires a integer number.  

* `1` is default number.
* `0`: Not occur anything. Displayed current stored number.
* `Negative number`: count down.

## Install

Download a binary via relase page (in preparation).

Or `go get`.

```bash
$ go get github.com/kawaken/ct
```

## Contribution

## Licence

[MIT](https://github.com/kawaken/ct/blob/master/LICENSE)

## Author

[kawaken](https://github.com/kawaken)

