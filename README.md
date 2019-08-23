# Golang-httpflood ![](https://img.shields.io/badge/Version-1.2-brightgreen.svg) ![](https://img.shields.io/badge/license-MIT-blue.svg)
Using Golang(net/socket) to httpflood

**Using tcp socket to send Http request**

## INFO

 - [x] HTTP Get Flood
 - [x] HTTP Post Flood
 - [x] Random url(http get flood)
 - [x] Self edit header(You can use "nil" to use default header)
 -----------------------------------------------------
 Default header setting:
 - [x] Random user-agents
 - [x] Random data(http post flood) 
 - [x] Accept: */*

## Download
***Please download the F\*cking golang at first.***

Then:

    git clone https://github.com/Leeon123/golang-httpflood.git

Header.txt format:

    Accept: text/html
    User-agent: Wget
    Referer: http://google.com

Or anything else of http header. If you don't have any idea of this please just use nil for using default random header.
## Usage

    cd golang-httpflood
    go build httpflood.go
    ./httpflood  <ip> <port> <threads> <page> <get/post> <seconds> <header.txt/nil>
