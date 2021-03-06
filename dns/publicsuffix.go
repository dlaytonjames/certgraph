package dns

import (
	"net/http"
	"time"

	"github.com/weppos/publicsuffix-go/publicsuffix"
)

var (
	suffixListFindOptions = &publicsuffix.FindOptions{
		IgnorePrivate: true,
		DefaultRule:   publicsuffix.DefaultRule,
	}
	suffixListURL = "https://publicsuffix.org/list/public_suffix_list.dat"
	suffixList    = publicsuffix.DefaultList
	nsCache       = make(map[string]bool)
)

// UpdatePublicSuffixList gets a new copy of the public suffix list from the internat and updates the built in copy with the new rules
func UpdatePublicSuffixList(timeout time.Duration) error {
	suffixListParseOptions := &publicsuffix.ParserOption{
		PrivateDomains: !suffixListFindOptions.IgnorePrivate,
	}
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(suffixListURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	newSuffixList := publicsuffix.NewList()
	newSuffixList.Load(resp.Body, suffixListParseOptions)
	suffixList = newSuffixList
	return err
}

// TLDPlus1 returns TLD+1 of domain
func TLDPlus1(domain string) (string, error) {
	return publicsuffix.DomainFromListWithOptions(suffixList, domain, suffixListFindOptions)
}
