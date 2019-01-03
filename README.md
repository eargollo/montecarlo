# MonteCarlo

This simple tool lets you use MonteCarlo simulation based on a data series to predict
into the future with different precision levels.

The reports will be generated as a TSV list.

## Usage

First create a file with historical data:
```
$ cat input.csv
5
3
6
4
8
4
5
```

Now run the tool with data file as argument:
```
$ ./montecarlo --input ./input.csv --future 12
...

To see more options:
```
$ montecarlo -h
```
