# github-prometheus-exporter
Exporter to export github stats to Prometheus, written by a guy who doesn't know
golang.

## Limitations
Currently only one metric available: The number of lines in `README.md` in the
root of each (hard coded) repository. The data is provided by the `readme` GitHub API ([https://developer.github.com/v3/repos/contents/#get-the-readme](https://developer.github.com/v3/repos/contents/#get-the-readme)).

## Usage
Build:
```
go build
```

Run:
```
./github-prometheus-exporter --github-token <token> -r <repo 1> ... -r <repo n>
```
where `<token>` is the GitHub application token and `<repo x>` is a list of
GitHub repositories (i.e. `hawkaa/CloakedMailman).


## Hacking
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
