# tripeg
Brute force solver of the triangle peg jumping game you see at Cracker Barrel.

I'm well aware that this is a solved game, I just wanted to write a program to solve it on my own.

1. `git clone` or `go get` this repo
1. In the tripegd directory type `go build`
1. run `go get github.com/genghisjahn/xlog`
1. `./tripegd` to run the solver for a random missing peg for a 5 row triangle.
1. `./tripegd 1` to run the solver for the top missing peg.  `./tripeg 0` will solve for a random missing peg to start for a 5 row triangle. 
  1. You can enter an integer as a first argument to specify a missing peg.  Pegs are numbers top to bottom, left to right.  Valid arguments are based on the number of rows (the 2nd parameter).
  1. You can enter an integer that's 5 or greater for the number of rows[default=5].  5 and 6 row triangles are solved quickly.  8 row triangles are sometimes solved in less than a minute([proof](https://gist.github.com/genghisjahn/6305d886454ffa87ed8efbf0a6ee949b)), but sometimes never return a solution.  7 and 9 row triangles have never returned  a solution.  A 10 row triangle has returned a solution a few times.  Haven't tried anything above 10 rows.

The solver will print out a solution to the console.  
* `+` is the peg that should be moved
* `0` is the hole where it should end up.  
* `O` represents an empty hole that isn't involved in the jump.  
* `*` represents a filled hole that isn't involved with the jump.

Just follow along and you'll be left with 1 peg on the board:
```
INFO: 2016/09/08 23:35:41 main.go:17: Tripeg Main
Move: 1
    0
   * *
  + * *
 * * * *
* * * * *

Move: 2
    *
   O *
  0 * *
 * * * *
+ * * * *

Move: 3
    *
   0 *
  * * *
 O * + *
O * * * *

Move: 4
    *
   + *
  * O *
 0 * O *
O * * * *

Move: 5
    *
   O *
  O 0 *
 * * O *
O + * * *

Move: 6
    *
   O +
  O * *
 * 0 O *
O O * * *

Move: 7
    *
   O O
  O O *
 * * O *
O 0 * + *

Move: 8
    *
   O 0
  O O *
 * * O +
O * O O *

Move: 9
    *
   O *
  O O O
 + * 0 O
O * O O *

Move: 10
    +
   O *
  O O 0
 O O * O
O * O O *

Move: 11
    O
   O O
  O O +
 O O * O
O * 0 O *

Move: 12
    O
   O O
  O O O
 O O O O
O + * 0 *

Move: 13
    O
   O O
  O O O
 O O O O
O O 0 * +
```
