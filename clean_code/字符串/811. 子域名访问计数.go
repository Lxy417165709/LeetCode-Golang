package 字符串

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	timesOfDomain := make(map[string]int)
	for i := 0; i < len(cpdomains); i++ {
		record(timesOfDomain, cpdomains[i])
	}
	return getFormatRecords(timesOfDomain)
}

func record(timesOfDomain map[string]int, cpdomain string) {
	times, domains := getTimesAndDomains(cpdomain)
	for t := 0; t < len(domains); t++ {
		timesOfDomain[domains[t]] += times
	}
}

func getFormatRecords(timesOfDomain map[string]int) []string {
	result := make([]string, 0)
	for domain, times := range timesOfDomain {
		result = append(result, fmt.Sprintf("%d %s", times, domain))
	}
	return result
}

func getTimesAndDomains(cpdomain string) (int, []string) {
	parts := strings.Split(cpdomain, " ")
	times, _ := strconv.Atoi(parts[0])
	domains := getDomains(parts[1])
	return times, domains
}

func getDomains(domain string) []string {
	domains := make([]string, 0)
	domains = append(domains, domain)
	for i := 0; i < len(domain); i++ {
		if domain[i] == '.' {
			domains = append(domains, domain[i+1:])
		}
	}
	return domains
}

/*
	题目链接: https://leetcode-cn.com/problems/subdomain-visit-count/
*/
