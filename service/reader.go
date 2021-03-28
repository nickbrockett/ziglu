package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"sort"
	"time"

	"priva.te/ziglu/lib"
)

//Item contains a news item
type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	Desc    string `xml:"description"`
	PubDate string `xml:"pubDate"`
	Key     string `xml:"guid"`
}

//Items contains a slice of pointers to a news item.
type Items []*Item

//Channel contains an RSS feed in its entirety.
type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

//Rss is a RSS Feed Channel
type Rss struct {
	Channel Channel `xml:"channel"`
}

// rssReader will visit the provided URL and decode XML into a RSS Channel Struct.
func rssReader(address string) Rss {

	rss := Rss{}

	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return rss
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return rss
	}
	return rss
}

// dateOrderedItems returns a date desc slice of news items.
func (items Items) dateOrderedItems() {

	dateFormat := "Mon, 02 Jan 2006 15:04:05 GMT"
	alternativeFormat := "Mon, 02 Jan 2006 15:04:05 Z0700"

	sort.Slice(items, func(i, j int) bool {

		iTime, err := time.Parse(dateFormat, items[i].PubDate)

		if err != nil {
			iTime, _ = time.Parse(alternativeFormat, items[i].PubDate)
		}

		jTime, err := time.Parse(dateFormat, items[j].PubDate)

		if err != nil {
			jTime, _ = time.Parse(alternativeFormat, items[j].PubDate)
		}

		return iTime.After(jTime)
	})

}

//getFeedItems returns a slice of news items.
func getFeedItems(provider, category string) Items {

	var allItems Items

	// apply filters to provide a refined list of addresses.
	feedProviders := lib.GetProviders().FilterByProvider(provider)
	feedAddresses := feedProviders.FilterAddressesByCategory(category)

	for _, address := range feedAddresses {
		rss := rssReader(address)
		//retain pointers to the rss items returned.
		for k := range rss.Channel.Items {
			allItems = append(allItems, &rss.Channel.Items[k])
		}
	}

	if allItems != nil {
		allItems.dateOrderedItems()
	}

	hashItemKey(allItems)

	return allItems
}

//hashItemKey converts news item unique reference to a hash.
func hashItemKey(items Items) {

	for _, v := range items {
		shaBytes := sha256.Sum256([]byte(v.Key))
		v.Key = b64.URLEncoding.EncodeToString(shaBytes[:])
	}

}
