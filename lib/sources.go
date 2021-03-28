package lib

// TODO ideally need to be stored externally.
var providers *Providers

//Category contains the category name and address.
type Category struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

//Provider contains the provider name and list of categories.
type Provider struct {
	Name       string     `json:"name"`
	Categories []Category `json:"categories"`
}

//Providers contains a list of Provider.
type Providers []Provider

//GetProviders returns all currently available providers of feeds.
func GetProviders() *Providers {
	if providers == nil {
		return DefaultProviders()
	}
	return providers
}

//SetProvider sets the currently available providers of feeds.
func SetProviders(p *Providers)  {
	providers = p
}

//Default Providers initiates the list of providers.
func DefaultProviders() *Providers {

	providers = &Providers{
		Provider{Name: "bbc", Categories: []Category{
			{Name: "news", Address: "http://feeds.bbci.co.uk/news/uk/rss.xml"},
			{Name: "technology", Address: "http://feeds.bbci.co.uk/news/technology/rss.xml"},
		}},
		Provider{Name: "sky", Categories: []Category{
			{Name: "news", Address: "http://feeds.skynews.com/feeds/rss/uk.xml"},
			{Name: "technology", Address: "http://feeds.skynews.com/feeds/rss/technology.xml"},
		}},
	}

	return providers
}

// FilterByProvider will return all providers if none specified or only those providers matching given filter.
func (p Providers) FilterByProvider(provider string) Providers {

	if provider == "" {
		return p
	}

	result := Providers{}

	for _, v := range p {
		if provider == v.Name || provider == "" {
			result = append(result, v)
		}
	}
	return result

}

// FilterAddressesByCategory will return all addresses if no category specified or only those matching given filter.
func (p Providers) FilterAddressesByCategory(category string) (result []string) {

	for _, v := range p {
		for _, c := range v.Categories {
			if category == c.Name || category == "" {
				result = append(result, c.Address)
			}
		}
	}
	return result

}

// GetNewsSources returns addresses available to search.
func GetNewsSources(provider, category string) []string {

	// apply filters to provide a refined list of addresses.
	feedProviders := GetProviders().FilterByProvider(provider)
	feedAddresses := feedProviders.FilterAddressesByCategory(category)

	return feedAddresses
}