QUEUEMETRE
===================================================================================================
queuemetre is a command for displaying the number of email queues like the iostat(8).

EXAMPLE
---------------------------------------------------------------------------------------------------
`queuemetre` automatically detects the SMTP daemon running on the host where it is executed.

### Sendmail
```
# /usr/local/sbin/queuemetre
      Date     Time    Total      MTA      MSA Deferred  Lost Quarantined  Proc:d/c  Conn:L/R   Load   MB:vsz/rss
2018-02-02 18:30:22        0        0        0        0     0           1       6/0       1/0   0.01        407/8
2018-02-02 18:31:22        0        0        0        0     0           1       6/0       2/0   0.00        407/8
2018-02-02 18:32:22        0        0        0        0     0           1       6/0       1/0   0.00        407/8
2018-02-02 18:33:22        0        0        0        0     0           1       4/0       1/0   0.00        338/5
2018-02-02 18:34:22        0        0        0        0     0           1       6/0       1/0   0.03        407/8
...
```

### Postfix
```
[root@mailout ~]# /tmp/queuemetre
      Date     Time    Total Incoming   Active Deferred     Hold Bounced  Proc:d/c  Conn:L/R   Load   MB:vsz/rss
2018-02-02 18:30:22      135        0       80       55        0       0     2/121     0/124   0.22   22241/2101
2018-02-02 18:31:22      127        0      125        2        0       0     2/122     0/127   0.25   22241/2104
2018-02-02 18:32:22      128        1      126        1        0       0     2/128      0/93   0.23   21244/2044
2018-02-02 18:33:22      127        3      103       21        0       0     2/126      0/90   0.08   21206/2032
2018-02-02 18:34:22      115        0      103       12        0       0     2/108      0/92   0.09   29278/2092
...
```


### OpenSMTPD
```
[root@p1 ~]# /tmp/queuemetre
      Date     Time    Total Incoming    Queue    Purge Offline  Proc:d/q  Conn:L/R   Load   MB:vsz/rss
2018-02-02 18:30:22        5        1        4        0       0       5/1       1/1   0.00       250/21
2018-02-02 18:31:22        5        1        4        0       0       5/1       1/1   0.00       250/21
2018-02-02 18:32:22        5        0        5        0       0       5/1       1/2   0.00       250/22
2018-02-02 18:33:22        5        0        5        0       0       5/1       0/1   0.00       250/22
2018-02-02 18:34:22        5        0        5        0       0       5/1       0/1   0.00       250/22
...
```

INSTALL
---------------------------------------------------------------------------------------------------
```
$ sudo make install
Password: ********
install -o root -m 0755 queuemetre /usr/local/sbin/queuemetre
```

```
$ /bin/cp ./queuemetre /path/to/somewhere/
$ chmod a+x /path/to/somewhere/queuemetre


Author
===================================================================================================
[@azumakuniyuki](https://twitter.com/azumakuniyuki)

Copyright
===================================================================================================
Copyright (C) 2024 azumakuniyuki, All Rights Reserved.

License
===================================================================================================
This software is distributed under The BSD 2-Clause License.

