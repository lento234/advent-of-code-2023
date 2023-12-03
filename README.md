<div align="center">

# 🎅🎄 Advent of Code 2023 ☃️❄️

![Last commit](https://img.shields.io/github/last-commit/lento234/advent-of-code-2023)
[![GitHub license](https://img.shields.io/github/license/lento234/advent-of-code-2023?color=blue)](https://github.com/lento234/advent-of-code-2023/blob/main/LICENSE)

</div>

## Description

- [Advent of code 2023](https://adventofcode.com/2023/)
- [Awesome Advent of Code](https://github.com/Bogdanp/awesome-advent-of-code)

Advent of Code is an Advent calendar of small programming puzzles for a variety
of skill sets and skill levels that can be solved in any programming language
you like. People use them as a speed contest, interview prep, company training,
university coursework, practice problems, or to challenge each other.

## Calendar

| M       | T   | W   | T   | F    | *S*  | *S*  |
|---------|-----|-----|-----|------|------|------|
|         |     |     |     | ⭐⭐ | ⭐⭐ | ⭐⭐ |
| 4       | 5   | 6   | 7   | 8    | 9    | 10   |
| 11      | 12  | 13  | 14  | 15   | 16   | 17   |
| 18      | 19  | 20  | 21  | 22   | 23   | 24   |
| [🎁][1] |     |     |     |      |      |      |

## Setup

1. Install go: [go.dev](https://go.dev/dl/).

2. Build project:

    ```bash
    go build
    ```
3. Lint:

    ```bash
    golangci-lint run
    ```

## Usage

**Initializing day:**

```bash
./aoc2023 init --day <day>
```

**Run challenge of the day:**

```bash
./aoc2023 run --day <day>
```

**Run tests:**

```bash
go test -bench -v ./day*
```

**Run benchmarks:** (with memory allocation statistics and `GOMAXPROCS` set to 1, 4 and 8, 24)

```bash
go test -run=^$ -bench=. -benchmem -cpu=1,4,8,24 ./...
```

**Easter eggs**

```bash
./aoc2023 xmas
```

[1]: https://youtu.be/mkF7xLtNzPc?si=jQ7NB9oxtYNauYwd&t=27
