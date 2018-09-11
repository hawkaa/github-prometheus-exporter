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

type readmeResponse struct {
	Content string `json:"content"`
}

// NewExporter returns an Exporter with the repositories injected
func NewExporter(repos []string, githubToken string) Exporter {
	var e Exporter
	e.repos = repos
	e.githubToken = githubToken
	return e
}

// Describe forwards the Readme metric description
func (exporter *Exporter) Describe(ch chan<- *prometheus.Desc) {
	// We need to send a description. What do we send?
	ch <- readme
}

// Collect loops the available repositories and returns a readme metric gauge per repo
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
	r := readmeResponse{}
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
