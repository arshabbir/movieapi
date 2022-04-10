FROM golang:bullseye

ENV GO111MODULE=on

RUN mkdir /build

COPY . /build

WORKDIR /build

RUN go build -o testapi

EXPOSE 8080

ENTRYPOINT ["./testapi"]
