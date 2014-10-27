yocto-httpd
===========

yhttp serves 200 responses to any GET request.
We use it for satisfying centurion's need for a http service running inside every container.

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
