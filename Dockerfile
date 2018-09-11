FROM golang:1.11.0 AS build
WORKDIR /go/src/github.com/hawkaa/github-prometheus-exporter/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /bin/github-prometheus-exporter
FROM scratch
COPY --from=build /bin/github-prometheus-exporter /bin/github-prometheus-exporter
EXPOSE 8080
CMD ["/bin/github-prometheus-exporter"]
