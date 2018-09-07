# github-prometheus-exporter
Exporter to export github stats to Prometheus.

## Hacknig
Dependencies:
```
dep ensure
```

Tests:
```
go test
```

Run development server:
```
./start-dev
```


## Alternatives
* [github-exporter](https://github.com/infinityworks/github-exporter) is
probably a much better piece of software, but that would remove my opportunity
to learn how to write a Prometheus exporter.
