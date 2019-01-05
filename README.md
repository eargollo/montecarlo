# MonteCarlo

This simple tool lets you use MonteCarlo simulation based on a data series to estimate
values in the future with different precision levels.

The reports will be generated at the standard output an it is a TSV list.

## Usage

First create a file with historical data, with one data point per line.

For example, let's say a team can execute the following amount of story points per sprint (saved in example.txt):
```
15
8
12
6
10
20
5
```

We want to estimate the next 12 sprints based on this historical data.

We run the `montecarlo` tool with the data file as argument and the 12 cycles for the future:
```
$ montecarlo.exe estimate --input ./example.txt --future 12
Generating randomized data...
Aggregating future data...
Calculating forecasts...
    0... 1... 2... 3... 4... 5... 6... 7... 8... 9... 10... 11...Done.
FuturePoints    12      Simulations     1000000
Conf%   1       2       3       4       5       6       7       8       9       10      11      12
100%    3       6       9       12      15      18      21      25      29      33      37      40
95%     3       7       11      15      20      24      29      33      38      43      47      52
90%     3       7       12      16      21      25      30      35      39      44      49      53
85%     4       8       12      17      21      26      31      36      40      45      50      55
80%     4       8       13      17      22      27      32      36      41      46      51      56
75%     4       8       13      18      23      27      32      37      42      47      51      56
70%     4       9       13      18      23      28      33      38      42      47      52      57
65%     4       9       14      19      24      28      33      38      43      48      53      58
60%     4       9       14      19      24      29      34      39      44      49      54      58
55%     5       9       14      19      24      29      34      39      44      49      54      59
50%     5       10      15      20      25      30      35      40      45      50      55      60
45%     5       10      15      20      25      30      35      40      45      50      55      60
40%     5       10      15      21      26      31      36      41      46      51      56      61
35%     5       11      16      21      26      31      36      42      47      52      57      62
30%     5       11      16      21      27      32      37      42      47      52      58      63
25%     6       11      17      22      27      32      38      43      48      53      58      63
20%     6       12      17      23      28      33      38      44      49      54      59      64
15%     6       12      18      23      29      34      39      44      50      55      60      65
10%     8       13      18      24      29      35      40      46      51      56      62      67
5%      8       14      20      25      31      36      42      47      53      58      64      69
0%      8       16      24      32      40      48      54      61      69      74      80      87
...

This means there is a 95% chance that the team will complete 52 user stories in the next 12 sprints. A 100% chance (based in past data) that it will complete 40 in the next 12 sprints. Or a 80% confidence it complete 27 user stories in the next 6 sprints.

This could be applied to story points or any other metric.


The tool runs 1 million simulations and predicts for every 5% increment as default. You can also change this value at the command line. To see more options:
```
$ montecarlo.exe -h
MonteCarlo: Given a sequence of measurements, and the amount of future data,
it applies MonteCarlo to predict future results based on past ones.

Usage:
  montecarlo [command]

Available Commands:
  estimate    Simulate future estimation
  help        Help about any command

Flags:
  -h, --help   help for montecarlo

Use "montecarlo [command] --help" for more information about a command.

$ montecarlo.exe estimate -h
Estimate towards the future based on past data

Usage:
  montecarlo estimate [flags]

Flags:
      --future int        Future data points (default 12)
  -h, --help              help for estimate
      --increment float   Percentual increment for each confidence data point. Default is 5, i.e. one data point for each 5%: 100%, 95%, 90%,...0% (default 5)
      --input string      Input data, one value per line (default "./input.csv")
      --simulations int   Amount of MonteCarlo simulations used (the bigger the number, the better the precision but it may take longer to simulate). (default 1000000)

```

