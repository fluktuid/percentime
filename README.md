# percentime

Executes a command n times and shows percentile of input numbers.

This command is very helpful if you need to get the percentile for a network request.

This Project is based on
- [yuya-takeyama/ntimes](https://github.com/yuya-takeyama/ntimes)
- [yuya-takeyama/percentile](https://github.com/yuya-takeyama/percentile)

## Installation

### plain (golang)

``` bash
$ go get github.com/fluktuid/percentime
```

### OS X (Homebrew)

``` bash
$ brew tap fluktuid/tap
$ brew install percentime
```

## Usage

### Execution Time
``` bash
$ percentime 100 -- sleep 0.1
50%:	0.1062
66%:	0.10707
75%:	0.10752
80%:	0.10811
90%:	0.10879
95%:	0.10911
98%:	0.10953
99%:	0.10999
100%:	0.11869
```

## Response Number
Append the `-c` flag to use the time `curl [...] -w "%{time_total}"` responds.
``` bash
$ percentime 100 -c -- curl -s -o /dev/null -w "%{time_total}" google.com
50%:	0.048027
66%:	0.049199
75%:	0.049855
80%:	0.051012
90%:	0.052951
95%:	0.056334
98%:	0.082925
99%:	0.095592
100%:	0.11349
```

``` bash
$ percentime 100 -c -- echo -n "0.01"
50%:	0.01
66%:	0.01
75%:	0.01
80%:	0.01
90%:	0.01
95%:	0.01
98%:	0.01
99%:	0.01
100%:	0.01

```


## License

The MIT License
