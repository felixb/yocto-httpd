yocto-httpd
===========

[![Build Status](https://travis-ci.org/felixb/yocto-httpd.svg?branch=master)][1]

yhttp serves 200 responses to any GET request.

Building
========

Build the binary from source:

    go build yhttpd.go

Installing
==========

Copy the binary to a directory in $PATH:

    sudo cp yhttpd /usr/local/bin

Running
=======

Run yhttpd (inside your docker container) on any port like this:

    yhttpd --port 8081

[1]: https://travis-ci.org/felixb/yocto-httpd
