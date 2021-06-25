# Benchmark for interprocess communication data transfer

This benchmarks aims to measure the theoretical maximum data rate achievable with various IPC methods.
Currently two methods are tested:

- Standard in/out: `stdio` server
- TCP/IP loopback: `tcp` server

The test tries to transfer 10GB of data with various blocksize.

The test result on my system (a laptop with an Intel(R) Core(TM) i7-8550U CPU @ 1.80GHz, 8GB DDR4 RAM 2666MHz) are as follows:

### Ubuntu Linux 4.15.0-145-generic

```
stdio BS=   1024: Read 10737418240 bytes in  10.562 sec: 969.498 MB/sec
stdio BS=   2048: Read 10737418240 bytes in   5.939 sec: 1724.132 MB/sec
stdio BS=   4096: Read 10737418240 bytes in   4.158 sec: 2462.911 MB/sec
stdio BS=   8192: Read 10737418240 bytes in   2.767 sec: 3700.944 MB/sec
stdio BS=  16384: Read 10737418240 bytes in   3.899 sec: 2626.530 MB/sec
stdio BS=  32768: Read 10737418240 bytes in   4.321 sec: 2369.885 MB/sec
stdio BS=  65536: Read 10737418240 bytes in   4.591 sec: 2230.405 MB/sec
stdio BS= 131072: Read 10737418240 bytes in   4.529 sec: 2261.034 MB/sec
stdio BS= 262144: Read 10737418240 bytes in   4.705 sec: 2176.544 MB/sec
stdio BS= 524288: Read 10737418240 bytes in   4.648 sec: 2202.912 MB/sec
stdio BS=1048576: Read 10737418240 bytes in   4.696 sec: 2180.791 MB/sec
tcpip BS=   1024: Read 10737418240 bytes in  35.416 sec: 289.136 MB/sec
tcpip BS=   2048: Read 10737418240 bytes in  15.675 sec: 653.271 MB/sec
tcpip BS=   4096: Read 10737418240 bytes in   8.271 sec: 1238.127 MB/sec
tcpip BS=   8192: Read 10737418240 bytes in   4.429 sec: 2312.237 MB/sec
tcpip BS=  16384: Read 10737418240 bytes in   2.739 sec: 3738.057 MB/sec
tcpip BS=  32768: Read 10737418240 bytes in   2.510 sec: 4080.480 MB/sec
tcpip BS=  65536: Read 10737418240 bytes in   2.750 sec: 3723.722 MB/sec
tcpip BS= 131072: Read 10737418240 bytes in   2.707 sec: 3782.208 MB/sec
tcpip BS= 262144: Read 10737418240 bytes in   2.634 sec: 3887.456 MB/sec
tcpip BS= 524288: Read 10737418240 bytes in   2.694 sec: 3800.393 MB/sec
tcpip BS=1048576: Read 10737418240 bytes in   2.616 sec: 3914.336 MB/sec
```

As we can see `stdio` handles better small block size, it reach the top speed at `8192` and quickly decrease leveling at roughly half of the top speed, as block size gets bigger.
`tcpip` instead seems to suffer when data is transferred with a lot of fragmentation but, as the block size increase, the protocol overhead becomes less and less relevant and the top speed is quickly reached and maintained even with very large block size.

### Windows 10 Virtual Machine

```
stdio BS=   1024: Read 10737418240 bytes in  43.387 sec: 236.018 MB/sec
stdio BS=   2048: Read 10737418240 bytes in  31.195 sec: 328.255 MB/sec
stdio BS=   4096: Read 10737418240 bytes in  20.602 sec: 497.039 MB/sec
stdio BS=   8192: Read 10737418240 bytes in  17.375 sec: 589.350 MB/sec
stdio BS=  16384: Read 10737418240 bytes in  16.727 sec: 612.202 MB/sec
stdio BS=  32768: Read 10737418240 bytes in   9.913 sec: 1032.972 MB/sec
stdio BS=  65536: Read 10737418240 bytes in   5.207 sec: 1966.657 MB/sec
stdio BS= 131072: Read 10737418240 bytes in   3.483 sec: 2940.062 MB/sec
stdio BS= 262144: Read 10737418240 bytes in   2.545 sec: 4023.986 MB/sec
stdio BS= 524288: Read 10737418240 bytes in   2.492 sec: 4108.759 MB/sec
stdio BS=1048576: Read 10737418240 bytes in   2.194 sec: 4666.860 MB/sec
tcpip BS=   1024: Read 10737418240 bytes in  76.799 sec: 133.335 MB/sec
tcpip BS=   2048: Read 10737418240 bytes in  38.160 sec: 268.340 MB/sec
tcpip BS=   4096: Read 10737418240 bytes in  19.173 sec: 534.089 MB/sec
tcpip BS=   8192: Read 10737418240 bytes in  10.776 sec: 950.255 MB/sec
tcpip BS=  16384: Read 10737418240 bytes in   8.079 sec: 1267.445 MB/sec
tcpip BS=  32768: Read 10737418240 bytes in  10.353 sec: 989.088 MB/sec
tcpip BS=  65536: Read 10737418240 bytes in  11.929 sec: 858.429 MB/sec
tcpip BS= 131072: Read 10737418240 bytes in  11.138 sec: 919.358 MB/sec
tcpip BS= 262144: Read 10737418240 bytes in  12.237 sec: 836.837 MB/sec
tcpip BS= 524288: Read 10737418240 bytes in   7.707 sec: 1328.642 MB/sec
tcpip BS=1048576: Read 10737418240 bytes in   6.786 sec: 1508.924 MB/sec
```

In this case we can see that `stdio` constantly outperforms network, and the performance scales better at higher block size, while the network stack seems to struggle.
