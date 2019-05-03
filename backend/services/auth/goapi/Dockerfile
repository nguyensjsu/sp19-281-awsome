FROM golang:latest
EXPOSE 9000
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GOPATH /app
RUN cd /app ; go get auth; go install auth
CMD ["/app/bin/auth"]