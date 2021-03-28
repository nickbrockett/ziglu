package main

import (
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"sort"
	"time"
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

type Rss struct {
	Channel Channel `xml:"channel"`
}

//ChannelReader visits source website and returns a list of news items.
type ChannelReader func(s string) Items

type Reader struct {
	sources       []string
	channelReader ChannelReader
	items         Items
}

func NewReader(sources []string, channelReader ChannelReader) Reader {
	return Reader{
		sources:       sources,
		channelReader: channelReader,
		items:         Items{},
	}
}

func (r *Reader) Read() {

	for _, v := range r.sources {
		r.items = append(r.items, r.channelReader(v)...)
	}
	// additional processing to ensure date order, and a unique retrievable reference.
	r.items.dateOrderedItems()
	r.items.hashItemKey()
}

// RSSReader (the default implementation) will visit the provided URL and decode XML into a RSS Channel Struct.
// then convert to a slice of actual news items.
func RSSReader(address string) Items {

	var allItems Items
	rss := Rss{}

	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return nil
	}
	defer resp.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return nil
	}
	for k := range rss.Channel.Items {
		allItems = append(allItems, &rss.Channel.Items[k])
	}
	return allItems
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

//hashItemKey converts news item unique reference to a hash.
func (items Items) hashItemKey() {

	for _, v := range items {
		shaBytes := sha256.Sum256([]byte(v.Key))
		v.Key = b64.URLEncoding.EncodeToString(shaBytes[:])
	}

}
