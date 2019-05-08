FROM golang:1.11.0-alpine3.8

## required stuff
#LABEL LOCATION="git@github.comcast.com:/Entitlements/user-rights-compute-service.git"
#LABEL DESCRIPTION=""

COPY ./.out/entarchdemo /usr/local/bin
EXPOSE 8080

CMD /usr/local/bin/entarchdemo
