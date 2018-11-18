# FROM golang:alpine
# MAINTAINER Alamin Mahamud <alamin.ineedahelp@gmail.com>
# RUN apk update && apk add git
# COPY . /go/src/app
# WORKDIR /go/src/app
# RUN go get -d -v ./...
# RUN go install -v ./...
# CMD ["app"]
FROM alpine
COPY awesome-go /bin/
EXPOSE 5001
ENTRYPOINT ["bin/awesome-go"]
