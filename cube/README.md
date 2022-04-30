## Cube module

This module provide API for cube rotations. For better performance i've decided to implement cube structure as a bitboard. What is a bitboard? Well essentially this is cube's representation in form of bits inside integers. For example in our case we have 6 64 bits integers for each cube's side. Each cube's side have 9 cubies(smaller cubes in the big cube) and they can be different colors. We can skip central cubie because the whole integer itself represents this central cubie(as central cubie represents side color).

### Reference

[What is bitboard](https://en.wikipedia.org/wiki/Bitboard)
