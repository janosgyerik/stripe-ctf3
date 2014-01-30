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
Task: improve `level0`. In whatever language. Your goal is to
make it as fast as possible compared to the benchmark.

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

The relevant part is: `You/Benchmark: 15.312883`

That's very poor. Try to get that rate below 0.5.


level1
------
TODO


level2
------
TODO


level3
------
TODO


level4
------
TODO
