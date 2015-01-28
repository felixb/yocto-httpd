FROM centos:6
MAINTAINER Felix Bechstein <felix.bechstein@otto.de>
EXPOSE 8080

CMD ["/usr/bin/yhttpd", "--port", "8080"]
ADD https://github.com/otto-de/yocto-httpd/releases/download/v0.1.0/yhttpd /usr/bin/yhttpd
RUN chmod 555 /usr/bin/yhttpd
