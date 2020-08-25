# Golang-httpflood ![](https://img.shields.io/badge/Version-2.0-brightgreen.svg) ![](https://img.shields.io/badge/license-MIT-blue.svg)
Using Golang(net/socket) to httpflood

**Warning: Please use command "ulimit -n 999999" before use this in linux**

**1 Threads =  1 connection, 100~300 connections can down a normal website in 10s(specially apache server LOL)**
 
**This is golang and threads are just goroutines so you set more higher threads like 1000-5000 is fine.**

**Why can it run over 1000 threads(goroutines)? [Read this](http://tleyden.github.io/blog/2014/10/30/goroutines-vs-threads/)**

## INFO

 - [x] HTTP Get Flood
 - [x] HTTP Post Flood
 - [x] Random url(http get flood)
 - [x] Self edit header(You can use "nil" to use default header)
 - [x] Improved threading control
 - [x] More powerful flood
 - [x] Auto get ip form domain(golang inbuilt function)
 - [x] More format for random url(http get flood)
 - [x] Fixed for 386 systems
 -----------------------------------------------------
 Default header setting:
 - [x] Random user-agents
 - [x] Random data(http post flood) 
 - [x] Random Accpetall
 - [x] Random Referer(only for http get flood)


## Download
***Please download the F\*cking golang at first.***

Then:

    git clone https://github.com/Leeon123/golang-httpflood.git

Header.txt format:

    Accept: text/html
    User-agent: Wget
    Referer: http://google.com

Or anything else of http header. If you don't have any idea of this please just use "nil" for using default random header.
## Usage

    cd golang-httpflood
    go build httpflood.go
    ./httpflood  <url> <threads> <get/post> <seconds> <header.txt/nil>
