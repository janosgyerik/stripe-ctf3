# Gitcoin Round 146

Welcome to the wonderful world of the Gitcoin mining round! This is
where the fun really begins, as well as the opportunity to score many
points.

## Overview

As before, the balances are all contained in `LEDGER.txt`, and we've
provided the same sample `miner` script.

The only commits that can be pushed are ones that:

- Increments an existing ledger entry by 1, or adds a new ledger entry
  with balance: 1; and
- Has a SHA1 starting with the number of zeros specified in `difficulty.txt`.

Your goal is to mine as many Gitcoins as possible. The round will
end at:

  2014-01-24 07:45:00 +0000

At that point, the next round will begin, accessible both at:

  <username>@stripe-ctf.com:current-round, and
  <username>@stripe-ctf.com:round147

You'll still be able to clone this code as:

  <username>@stripe-ctf.com:round146

## Scoring

After each round, we'll order all of the participants by number of
Gitcoins mined. We'll then run Elo [1] to update their rating
(starting from 1000). Your level score will be your rating, which
sticks around from round to round, plus any bonus points you've
accumulated.

Bonus points are earned in each round where you've mined at least one
Gitcoin. For each such round, we'll give you the number of bonus
points specified in that round's `points.txt` (for this round,
5).

For a detailed view of how the round is progressing, check out [2].

## Catalog

Note we'll dynamically adjust the values in the `.txt` files in order
to rebalance the contest (just like the Gamemakers from the Hunger
Games).

- `difficulty.txt`: The number of leading zeros in a valid Gitcoin
  commit.

- `miner`: A sample Gitcoin mining script.

- `round.txt`: The round number.

- `points.txt`: The number of bonus points that participation in this
  round is worth.

- `end-time.txt`: The round end time, as a Unix timestamp [3].

- `start-time.txt`: The round start time, as a Unix timestamp [3].

- `LEDGER.txt`: Contains the current Gitcoin balances.

- `README.md`: This file.

## Nonce

This nonce prevents mining before the round begins: tGyKEvwzGB

[1] http://en.wikipedia.org/wiki/Elo_rating_system
[2] https://stripe-ctf.com/leaderboard/1
[3] http://www.epochconverter.com/
