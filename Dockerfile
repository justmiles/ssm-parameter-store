FROM golang:1.12 AS builder
WORKDIR /go/src/github.com/justmiles/ssm-parameter-store/
RUN adduser --system tooluser
COPY . /go/src/github.com/justmiles/ssm-parameter-store/
RUN go get -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /bin/ssm-parameter-store

FROM scratch
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /bin/ssm-parameter-store /bin/ssm-parameter-store
USER tooluser
ENTRYPOINT ["/bin/ssm-parameter-store"]
