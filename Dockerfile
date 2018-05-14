FROM golang:alpine as build
ENV CGO_ENABLED 0
ENV GOOS linux
COPY . /go/src/github.com/concourse-sonarqube-notifier

RUN mkdir -p /assets \
 && go test -v github.com/concourse-sonarqube-notifier/assets/in/main \
 && go test -v github.com/concourse-sonarqube-notifier/assets/out/main \
 && go test -v github.com/concourse-sonarqube-notifier/assets/check/main \
 && go build -a -installsuffix cgo -o /assets/in github.com/concourse-sonarqube-notifier/assets/in/main \
 && go build -a -installsuffix cgo -o /assets/out github.com/concourse-sonarqube-notifier/assets/out/main \
 && go build -a -installsuffix cgo -o /assets/check github.com/concourse-sonarqube-notifier/assets/check/main

FROM scratch
COPY --from=build /etc/ssl/certs/ /etc/ssl/certs
COPY --from=build assets/ /opt/resource/