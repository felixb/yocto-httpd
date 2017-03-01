.PHONY: docker

yhttpd: yhttpd.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o yhttpd yhttpd.go

docker: yhttpd
	docker build -t felixb/yocto-httpd .
