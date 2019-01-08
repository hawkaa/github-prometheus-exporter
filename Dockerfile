FROM golang:1.11.2 AS build
RUN mkdir /github-prometheus-exporter
WORKDIR /github-prometheus-exporter
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix \
    cgo -o /bin/github-prometheus-exporter
FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /bin/github-prometheus-exporter /bin/github-prometheus-exporter
EXPOSE 8080
ENTRYPOINT ["/bin/github-prometheus-exporter"]
