package exporter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

var readme = prometheus.NewDesc(
	prometheus.BuildFQName("github", "doc", "readme_lines"),
	"The number of lines in the README.md file in the repository",
	[]string{"repo"},
	nil,
)

// Exporter class that implements the Collector interface
type Exporter struct {
	repos       []string
	githubToken string
}

type ReadmeResponse struct {
	Content string `json:"content"`
}

// NewExporter returns an Exporter with the repositories injected
func NewExporter(repos []string, githubToken string) Exporter {
	var e Exporter
	e.repos = repos
	e.githubToken = githubToken
	return e
}

// Describe sends the super-set of all possible descriptors of metrics
// collected by this Collector to the provided channel and returns once
// the last descriptor has been sent. The sent descriptors fulfill the
// consistency and uniqueness requirements described in the Desc
// documentation.
//
// It is valid if one and the same Collector sends duplicate
// descriptors. Those duplicates are simply ignored. However, two
// different Collectors must not send duplicate descriptors.
//
// Sending no descriptor at all marks the Collector as “unchecked”,
// i.e. no checks will be performed at registration time, and the
// Collector may yield any Metric it sees fit in its Collect method.
//
// This method idempotently sends the same descriptors throughout the
// lifetime of the Collector.
//
// If a Collector encounters an error while executing this method, it
// must send an invalid descriptor (created with NewInvalidDesc) to
// signal the error to the registry.
func (exporter *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// We need to send a description. What do we send?
	ch <- readme
}

// Collect is called by the Prometheus registry when collecting
// metrics. The implementation sends each collected metric via the
// provided channel and returns once the last metric has been sent. The
// descriptor of each sent metric is one of those returned by Describe
// (unless the Collector is unchecked, see above). Returned metrics that
// share the same descriptor must differ in their variable label
// values.
//
// This method may be called concurrently and must therefore be
// implemented in a concurrency safe way. Blocking occurs at the expense
// of total performance of rendering all registered metrics. Ideally,
// Collector implementations support concurrent readers.
func (exporter *Exporter) Collect(ch chan<- prometheus.Metric) {
	for _, repo := range exporter.repos {
		url := fmt.Sprintf("https://%s@api.github.com/repos/%s/readme", exporter.githubToken, repo)
		length, err := getRepoReadmeLength(url)
		if err != nil {
			log.WithError(err).Warningln("Could not get readme length of repository")
			continue
		}
		ch <- prometheus.MustNewConstMetric(
			readme,
			prometheus.GaugeValue,
			float64(length),
			repo,
		)
	}
}

func getRepoReadmeLength(url string) (int, error) {
	log.Infof("Querying %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return 0, errors.New("Unable to query API")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return 0, errors.New("Did not receive 200 OK")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, errors.New("Could not read response")
	}
	r := ReadmeResponse{}
	jsonErr := json.Unmarshal(body, &r)
	if jsonErr != nil {
		return 0, errors.New("Unable to read JSON response")
	}
	decoded, err := base64.StdEncoding.DecodeString(r.Content)
	if err != nil {
		return 0, errors.New("Unable to decode base 64 string")
	}
	lineSep := []byte{'\n'}
	// The two hard problems in computer science: Cache invalidation, naming, and off-by-one errors
	return bytes.Count(decoded, lineSep) + 1, nil
}
