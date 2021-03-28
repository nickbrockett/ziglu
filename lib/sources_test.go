package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testProviders Providers

func initTestProviders() {

	testProviders = Providers{
		Provider{Name: "provider1", Categories: []Category{
			{Name: "category1", Address: "provider1Category1Address"},
			{Name: "category2", Address: "provider1Category2Address"},
		}},
		Provider{Name: "provider2", Categories: []Category{
			{Name: "category1", Address: "provider2Category1Address"},
			{Name: "category2", Address: "provider2Category2Address"},
		}},
	}
}

func TestGetProviders(t *testing.T) {

	initTestProviders()
	providers = &testProviders
	assert.Equal(t, testProviders, *GetProviders())
}

func TestSetProviders(t *testing.T) {

	initTestProviders()
	providers = nil                               // set providers to nil
	GetProviders()                                // GetProviders sets to our default.
	assert.NotEqual(t, testProviders, *providers) // Confirm non-equality.
	SetProviders(&testProviders)
	assert.Equal(t, testProviders, *providers)

}

func TestProviders_FilterByProviderAndCategory(t *testing.T) {

	initTestProviders()

	for _, test := range []struct {
		name     string
		provider string
		category string
		expected []string
	}{
		{name: "nothing filtered",
			provider: "",
			category: "",
			expected: []string{"provider1Category1Address",
				"provider1Category2Address",
				"provider2Category1Address",
				"provider2Category2Address"},
		},
		{name: "only provider 1",
			provider: "provider1",
			category: "",
			expected: []string{"provider1Category1Address",
				"provider1Category2Address"},
		},
		{name: "only category 2",
			provider: "",
			category: "category2",
			expected: []string{"provider1Category2Address",
				"provider2Category2Address"},
		},
		{name: "only provider 1 and category 2",
			provider: "provider1",
			category: "category2",
			expected: []string{"provider1Category2Address"},
		},
	} {
		t.Run(test.name, func(t *testing.T) {

			result := testProviders.FilterByProvider(test.provider).FilterAddressesByCategory(test.category)
			assert.Equal(t, test.expected, result)
		})
	}

}
