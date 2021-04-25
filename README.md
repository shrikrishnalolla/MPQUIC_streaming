# MPQUIC - Webcam Streaming

This is an implementation of [Multipath QUIC](https://github.com/qdeconinck/mp-quic/) transport layer protocol for live video (webcam) streaming. We attempted to perform webcam streaming in a loopback interface. 

## Installation

First,setup MPQUIC as described in the following [link](https://multipath-quic.org/2017/12/09/artifacts-available.html). Note that the given method is for mininet VM and necessary adjustments need to be made to run it on real systems. This experiment was done on a real system. This also installs go 1.9.2.

Next, clone the repository or download the file as a zip.
```
git clone https://github.com/shrikrishnalolla/mpquic_streaming
```
The code requires [Gocv](https://gocv.io/) package of Go which is used to capture images from the webcam. Install it from the official page of [Gocv](https://gocv.io/)

## Running the Code

Clear the frame_save folder, this folder must be empty before running the code. 

Start the server(receiver.go) code.

~~~
go run receiver.go <insert_path_to_frame_save_directory>
~~~

Start the client(sender.go) code.

~~~
go run sender.go <camera_id> <server_ip>
~~~

In this implementation, the server ip is set to 0.0.0.0:4242 and the camera id as 0

Lastly, to view the video, run viewer.py

~~~
python viewer.py
~~~

We use matplotlib to display the video. We obtain 20-30 fps which is decent. The code deletes the frame once displayed.
