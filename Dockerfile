FROM golang:alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -v

# some thing
