FROM golang:1

COPY . ./src/server
WORKDIR ./src/server

RUN CGO_ENABLED=0 go build .

FROM scratch

WORKDIR /

COPY --from=0 /go/src/server/* /

CMD [ "/server" ]

