package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type wordpressSite struct {
	hostname string
	version string
}

var wpVersionRegexList = []string{
	`wp\-embed\.min\.js\?ver=([\d\.]+)`,
	`comment\-reply\.min\.js\?ver=([\d\.]+)`,
	`wp\-emoji\-release\.min\.js\?ver=([\d\.]+)`,
	`content="WordPress ([\d\.]+)`,
}

var domainListPath = flag.String("domain-list", "domains.txt", "Path to the list of domains to be surveyed.")

var scanStream = make(chan *wordpressSite)

func main() {
	flag.Parse()
	domainList, err := ioutil.ReadFile(*domainListPath)
	if err != nil {
		fmt.Println("error", err)
	}
	domains := regexp.MustCompile(`\n+`).Split(string(domainList), -1)
	n := 0
	for _, d := range domains {
		if len(d) > 0 && d[0] != '#' {
			w := &wordpressSite{string(d), ""}
			go getWPVersion(w)
			n++
		}
	}

	for i := n; i != 0; i-- {
		site := <-scanStream
		fmt.Println("Site", site.hostname, "is on version", site.version)
	}
}


func getWPVersion(s *wordpressSite) {
	resp, err := http.Get(fmt.Sprintf("http://%s", s.hostname))
	if err != nil {
		fmt.Println("Error getting version for", s.hostname)
		s.version = "unknown"
		scanStream <- s
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	for _, r := range wpVersionRegexList {
		match := regexp.MustCompile(r).FindStringSubmatch(string(body))
		if len(match) > 1 {
			s.version = match[1]
		}
	}
	scanStream <- s
}
