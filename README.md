# HackerRank challenges

This repo contains solutions to the challenges on HackerRank I am taking for fun
and to prepare for interviews. Roughly one challenge a day will be added.

Solutions are stored under `solutions/`. Each challenge has a root folder named
after itself. Each of such folders contains the solution under
`pkg/solution/<problem_name>.go`. For instance, the solution to the problem with
name [SockMerchant](SockMerchant) is located under
[`solutions/SockMerchant/pkg/solution/sockmerchant.go`](solutions/SockMerchant/pkg/solution/sockmerchant.go).
Under the root dir of each challenge there's a README with a link to the
challenge statement and other information.

All solutions are in Go.
In the files containing the solutions there are comments stating time and
space complexity.

## Solved challenges

* Easy
  * [SockMerchant](solutions/SockMerchant)
  * [CountingValleys](solutions/CountValleys)
  * [JumpingOnClouds](solutions/JumpingOnClouds)
  * [RepeatedString](solutions/RepeatedString)
* Medium
* Hard

## Support utils

This repo dir structure is quite regular, which makes it easy to automate some
tasks with scripts. [utils](utils) contains the scripts I use to
automate such tasks.