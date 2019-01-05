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
$ ./montecarlo.exe estimate --input ./example.txt --future 12
Generating randomized data...
Aggregating future data...
Calculating forecasts...
    0... 1... 2... 3... 4... 5... 6... 7... 8... 9... 10... 11...Done.
FuturePoints    12      Simulations     1000000
Conf%   1       2       3       4       5       6       7       8       9       10      11      12
100%    5       10      15      20      25      30      35      40      47      54      60      69
95%     5       11      20      28      37      46      56      65      74      84      93      103
90%     5       13      22      31      40      50      60      69      79      89      99      109
85%     6       14      23      33      43      52      62      72      82      92      102     113
80%     6       15      25      35      45      55      65      75      85      95      105     116
75%     6       16      26      36      46      57      67      77      87      98      108     118
70%     8       17      28      38      48      58      69      79      90      100     110     121
65%     8       18      29      39      49      60      70      81      92      102     113     123
60%     8       20      30      40      51      62      72      83      93      104     115     125
55%     10      20      31      42      52      63      74      85      95      106     117     128
50%     10      21      32      43      54      65      75      86      97      108     119     130
45%     10      22      33      44      55      66      77      88      99      110     121     132
40%     12      23      35      46      57      68      79      90      101     112     123     134
35%     12      25      36      47      58      69      81      92      103     114     125     136
30%     12      25      37      48      60      71      83      94      105     116     128     139
25%     15      26      38      50      62      73      85      96      108     119     130     142
20%     15      28      40      52      63      75      87      99      110     122     133     145
15%     15      30      42      54      66      78      90      101     113     125     136     148
10%     20      32      45      57      69      81      93      105     117     129     141     153
5%      20      35      47      60      73      86      98      111     123     135     147     159
0%      20      40      60      80      100     120     140     155     172     185     205     225
...

This means there is a 95% chance that the team will complete 103 user stories in the next 12 sprints. A 100% chance (based in past data) that it will complete at least 69 in the next 12 sprints. Or a 80% confidence the team will complete 55 user stories in the next 6 sprints.

This could be applied to story points or any other metric.

If you want a file that can be easily imported into MS Excel, you can pass the option --csv and the output to a file:

```
$ montecarlo.exe estimate --input ./example.txt --future 12 --csv > restult.csv
```

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

