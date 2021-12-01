# Advent of Code

[2015](/2015/README.md) | [2016](/2016/README.md) | [2017](/2017/README.md) | 
[2018](/2018/README.md) | [2019](/2020/README.md) | [2020](/2020/README.md) | [2021](/2021/README.md)

This repository contains my solutions for [Advent of Code](https://adventofcode.com/).

The folder for each day contains:
- Task of the day
- Input for the given task
- Code to solve the given task
- My solutions to the given input

## Requirements

- PHP 7.4
- Advent of Code account

## Installation

- `cp config.php.example config.php`
- Put your session token under: `config.php` file to be able to download descriptions and inputs.

## Usage

- `bin/task --year 2020 --day 1` - will prepare everything for the given task.
- `bin/run --year 2020 --day 1` - will execute all `code.*` files for the given task.
- `bin/readme --year 2020` - will create a `README.md` for the given year.
- `bin/test` - will test all `code.*` files. Beware this can overwrite your `benchmark.json`!!!

### Languages

You can add every language you want use. 
You only need to add the binary for executing your script into `binary-ext.map.php`.

Already added languages:
- PHP
- Javascript/Node (in `template/.ignore/`)
- Golang (in `template/.ignore/`)

To enable ignored templates for `bin/task`, you need to moved them manually to `template/`.