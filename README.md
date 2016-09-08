# tripeg
Brute force solver of the triangle peg jumping game you see at Cracker Barrel.

I'm well aware that this is a solved game, I just wanted to write a program to solve it on my own.

1. `git clone` or `go get` this repo
1. In the tripegd directory type `go build`
1. `./tripegd` to run the solver for a random missing peg.
1. `.tripegd 1` to run the solver for the top missing peg.
  1. You can enter 1-15 as a first argument to specify a missing peg.  Pegs are numbers top to bottom, left to right.

The solver will print out a solution to the console.  
The red * is the peg that should be moved and the green O is the hole where it should end up.

Just follow along and you'll be left with 1 peg on the board:
<img width="559" alt="screen shot 2016-09-07 at 10 06 45 pm" src="https://cloud.githubusercontent.com/assets/921877/18334735/cb00db4a-7547-11e6-8272-9757da58c6c9.png">

And so on...
