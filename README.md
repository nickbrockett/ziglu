# ziglu

This service exposes the following endpoints.
1) [GET /sources] A list of RSS News Feed sources can be retrieved. (Comprising a Provider and a list of their News Categories with URI)
    e.g. http://localhost:8080/sources which demonstates a default set of providers with associated categories and links.
    
   [
    {
        "name": "bbc",
        "categories": [
            {
                "name": "news",
                "address": "http://feeds.bbci.co.uk/news/uk/rss.xml"
            },
            {
                "name": "technology",
                "address": "http://feeds.bbci.co.uk/news/technology/rss.xml"
            }
        ]
    },
    {
        "name": "sky",
        "categories": [
            {
                "name": "news",
                "address": "http://feeds.skynews.com/feeds/rss/uk.xml"
            },
            {
                "name": "technology",
                "address": "http://feeds.skynews.com/feeds/rss/technology.xml"
            }
        ]
    }
]
    
2) [POST /sources] The list of RSS News Feed sources can be edited/changed. i.e. the above list can be reformulated.
3) [GET /feeds{?provider,category}] All RSS News items that optionally satisfy a provider/and or category filter (e.g. http://localhost:8080/feeds?provider=bbc&category=news) will support the Client to display a scollable set of related news items. Note. The individual news item is given its KEY identifier from the service. Short example output given here 
4) [
    {
        "Title": "Covid in Wales: Lockdown review 'will give hospitality certainty'",
        "Link": "https://www.bbc.co.uk/news/uk-wales-56551184",
        "Desc": "First minister to give hospitality sector \"the certainty it's looking for\" in lockdown rules review.",
        "PubDate": "Sun, 28 Mar 2021 17:21:03 GMT",
        "Key": "zLs-V3JgSZuCmy6jM3EYCasnSFo_GnSsuW1G752wcHQ="
    },
    {
        "Title": "Coronavirus: UK vaccine offer to Ireland 'a runner', says Arlene Foster",
        "Link": "https://www.bbc.co.uk/news/uk-northern-ireland-56556125",
        "Desc": "Arlene Foster says she will again ask Boris Johnson to offer vaccines to the Republic of Ireland.",
        "PubDate": "Sun, 28 Mar 2021 17:07:04 GMT",
        "Key": "WgjjXFyVAbI8dGYg2sx9pQ3q0wbMitIIROpd7RDw26o="
    }
    ]
    
4) [GET /feeds/{feed_id}] A single RSS News Item, retrieved by unique hash value. e.g. http://localhost:8080/feeds/WgjjXFyVAbI8dGYg2sx9pQ3q0wbMitIIROpd7RDw26o= will produce a HTML output for that item, with the response cached for performance for future and repeated access. Example output below 
5) <html>

<head>

	<title>News Item</title>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<meta charset="UTF-8">


	<link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/css/bootstrap.min.css"
		integrity="sha384-1q8mTJOASx8j1Au+a5WDVnPi2lkFfwwEAa8hDDdjZlpLegxhjVME1fgjWPGmkzs7" crossorigin="anonymous">
	<script async src="https://maxcdn.bootstrapcdn.com/bootstrap/3.3.6/js/bootstrap.min.js"
		integrity="sha384-0mSbJDEHialfmuBBQP6A4Qrprq5OVfW37PRR3j5ELqxss1yVqOtnepnHVP9aJ7xS" crossorigin="anonymous">
	</script>
</head>

<body class="container">


	<h2>News Item</h2>

	<a href="https://www.bbc.co.uk/news/uk-northern-ireland-56556125">

		<h3>Coronavirus: UK vaccine offer to Ireland &#39;a runner&#39;, says Arlene Foster</h3>
	</a>
	<p>Arlene Foster says she will again ask Boris Johnson to offer vaccines to the Republic of Ireland.</p>
	<p>Sun, 28 Mar 2021 17:07:04 GMT</p>




</body>

</html>
