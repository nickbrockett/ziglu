# ziglu

This service exposes the following endpoints.
1) [GET] A list of RSS News Feed sources can be retrieved. (Comprising a Provider and a list of their News Categories with URI)
    e.g. http://localhost:8080/sources
    [
    {
        "name": "bbc",
        "categories": [
            {
                "name": "news",
                "category": "http://feeds.bbci.co.uk/news/uk/rss.xml"
            },
            {
                "name": "technology",
                "category": "http://feeds.bbci.co.uk/news/technology/rss.xml"
            }
        ]
    },
    {
        "name": "sky",
        "categories": [
            {
                "name": "news",
                "category": "http://feeds.skynews.com/feeds/rss/uk.xml"
            },
            {
                "name": "technology",
                "category": "http://feeds.skynews.com/feeds/rss/technology.xml"
            }
        ]
    }
]
    
2) [POST] The list of RSS News Feed sources can be edited/changed.
3) [GET] All RSS News items that optionally satisfy a provider/and or category filter (e.g. http://localhost:8080/feeds?provider=bbc&category=news)
4) [GET] A single RSS News Item, retrieved by unique hash value. e.g. http://localhost:8080/feeds/WgjjXFyVAbI8dGYg2sx9pQ3q0wbMitIIROpd7RDw26o=
