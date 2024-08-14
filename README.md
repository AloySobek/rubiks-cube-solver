# Rubik's Cube solver

[Rubik's Cube](https://en.wikipedia.org/wiki/Rubik%27s_Cube) solver written in Go.

## Overview

Rubik's cube is a puzzle in the form of a plastic cube covered with
multicoloured squares, which the player attempts to twist and turn
so that all the squares on each face are of the same colour.

## Usage

```
NAME:
   Rubik - Rubik's cube solver

USAGE:
   Rubik [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -d          Colored 2D cube visualization (default: false)
   -v          Detailed description of algorithm's steps (default: false)
   -r value    Generate move randomly with specified number of rotations (default: 30)
   --help, -h  show help (default: false)
```

The super-flip scramble: R L U2 F U’ D F2 R2 B2 L U2 F’ B’ U R2 D F2 U R2 U

https://rubikscu.be/
