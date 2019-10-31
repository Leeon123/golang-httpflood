# Golang-httpflood ![](https://img.shields.io/badge/Version-1.6-brightgreen.svg) ![](https://img.shields.io/badge/license-MIT-blue.svg)
Using Golang(net/socket) to httpflood

**Warning: Please use command "ulimit -n 999999" before use this in linux**

Using tcp socket to send Http request

**1 Threads =  1 connection, Normal website about using 300-1000 connections will down in 10s(specially apache server LOL)**
 
**This is golang and threads are just goroutines so you set more higher threads like 1000-5000 is fine.**

**Why can it run over 1000 threads(goroutines)? [Read this](http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/)**

## INFO

 - [x] HTTP Get Flood
 - [x] HTTP Post Flood
 - [x] Random url(http get flood)
 - [x] Self edit header(You can use "nil" to use default header)
 - [x] Improved threading control
 - [x] More powerful flood
 -----------------------------------------------------
 Default header setting:
 - [x] Random user-agents
 - [x] Random data(http post flood) 
 - [x] Random Accpet

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
    ./httpflood  <host/ip> <port> <threads> <page> <get/post> <seconds> <header.txt/nil>
