FROM golang:1

COPY . ./src/client
WORKDIR ./src/client

RUN CGO_ENABLED=0 go build .

FROM scratch

WORKDIR /

COPY --from=0 /go/src/client/* /

CMD [ "/client" ]

