FROM scratch
MAINTAINER Felix Bechstein <f@ub0r.de>
EXPOSE 8080
ENTRYPOINT ["/yhttpd"]
ADD yhttpd /yhttpd
