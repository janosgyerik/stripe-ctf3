Capture The Flag 3: Distributed Systems
=======================================

[This edition](https://stripe-ctf.com/about) of the Capture The Flag contest by [@Stripe](https://twitter.com/stripe)
was about distributed systems and performance:

> This Capture the Flag is all about distributed systems.
> There are five levels, each one focused on a different problem
> in the field. In each case, the problem is one you've likely
> read about many times but never had a chance to try out in practice.

This repository contains the source code of the original levels,
*without solutions*. You can find solutions and write-ups by
many participants in https://github.com/ctfs/write-ups

The purpose of this repo is to make it easy to implement your
own solutions and compare them against a benchmark.

Each level contains a README.md, which should get you started.

level0
------
To test your implementation against the benchmark, simply run:

    ./test/harness

Example run:

```
No test case supplied. Randomly choosing among defaults.
About to run test case: level0-R5ez7L9mIu
Beginning run.
Finished run
Test case passed. Your time: 9.513440 seconds. Benchmark time: 0.621270 seconds. You/Benchmark: 15.312883
```

`You/Benchmark: 15.312883` is very poor.
Try to get that rate below 0.5


level1
------
TODO


level2
------
To test your implementation against the benchmark, simply run:

    ./test/harness

Example run:

```
... (skipped)
Number of total responses 528
Number of good responses: 30
Number of responses less than ideal: 538.6666666666667
Finished run
Test case passed. Your score: 0.010000. Benchmark score: 124.791667. You/Benchmark: 0.000080.
```

`You/Benchmark: 0.000080` is very poor.
Try to get that rate below 2.0


level3
------
TODO


level4
------
TODO
