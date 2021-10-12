#!/bin/sh

if [ ! -d certs/ ]; then
    mkdir certs/
fi
apk add openssl
cd certs
/usr/bin/openssl genrsa -out server.key 2048
/usr/bin/openssl req -nodes -new -x509 -sha256 -days 1825 -config cert.conf -extensions 'req_ext' -key server.key -out server.crt