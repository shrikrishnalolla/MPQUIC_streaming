# MPQUIC - webcam streaming

This is an implementation of [Multipath QUIC](https://github.com/qdeconinck/mp-quic/) transport layer protocol for live video (webcam) streaming. We attempted to perform webcam streaming in a loopback interface. 

First,setup MPQUIC as described in the following [link](https://multipath-quic.org/2017/12/09/artifacts-available.html). Note that the given method is for mininet VM and necessary adjustments need to be made to run it on real systems. This experiment was done on a real system. 

The code is written in Go(1.9.2) and Python(3.83). It requires [Gocv](https://gocv.io/) package of Go which is used to capture images from the webcam. 
