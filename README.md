# Golang-httpflood ![](https://img.shields.io/badge/Version-1.7-brightgreen.svg) ![](https://img.shields.io/badge/license-MIT-blue.svg)
Using Golang(net/socket) to httpflood

**Warning: Please use command "ulimit -n 999999" before use this in linux**

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
 - [x] Auto get ip form domain(golang inbuilt function)
 - [x] More format for random url(http get flood)
 - [x] Fixed for 386 systems
 -----------------------------------------------------
 Default header setting:
 - [x] Random user-agents
 - [x] Random data(http post flood) 
 - [x] Random Accpet

## Sth need to talk

I'll update a socks5 httpflood(golang) base later, I'h already done with that project and

down lots of site. But I'll not release my version,  becuase it's little dangerous.

I used it to take down some gov website so i need to change the random url format( LOL)

The release will not have any function, just send simple http request through socks5 with socket

You can copy this source's fucntion to that source.

Base: https://github.com/Leeon123/Golang_CC_Base

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
