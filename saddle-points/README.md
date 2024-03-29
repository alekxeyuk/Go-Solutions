# Saddle Points

Welcome to Saddle Points on Exercism's Go Track.
If you need help running the tests or submitting your code, check out `HELP.md`.

## Instructions

Detect saddle points in a matrix.

So say you have a matrix like so:

```text
    1  2  3
  |---------
1 | 9  8  7
2 | 5  3  2     <--- saddle point at row 2, column 1, with value 5
3 | 6  6  7


rowsMax [9, 5, 7]
colsMin [5, 3, 2]

for r := 0; r < rSize; r++ {
  for c := 0; c < cSize; c++ {
    if m[r][c] >= rowsMax[r] && m[r][c] <= colsMin[c] {
      print(m[r][c], "saddle found", r, c)
    }
  }
}
```

It has a saddle point at row 2, column 1.

It's called a "saddle point" because it is greater than or equal to
every element in its row and less than or equal to every element in
its column.

A matrix may have zero or more saddle points.

Your code should be able to provide the (possibly empty) list of all the
saddle points for any given matrix.

The matrix can have a different number of rows and columns (Non square).

Note that you may find other definitions of matrix saddle points online,
but the tests for this exercise follow the above unambiguous definition.

## Source

### Created by

- @soniakeys

### Contributed to by

- @alebaffa
- @bitfield
- @dvrkps
- @ekingery
- @ferhatelmas
- @hilary
- @kytrinyx
- @leenipper
- @petertseng
- @robphoenix
- @sebito91
- @tleen
- @eklatzer

### Based on

J Dalbey's Programming Practice problems - http://users.csc.calpoly.edu/~jdalbey/103/Projects/ProgrammingPractice.html