# Wikipedia Scraper

[![Oxylabs promo code](https://user-images.githubusercontent.com/129506779/250792357-8289e25e-9c36-4dc0-a5e2-2706db797bb5.png)](https://oxylabs.go2cloud.org/aff_c?offer_id=7&aff_id=877&url_id=112)

With [Wikipedia Scraper](https://oxylabs.io/products/scraper-api/web/wikipedia), you can gather public data, such as article content, its edit history, images, profile pages, and much more, from Wikipedia pages on a large scale and without interruptions. 

Follow this guide to scrape Wikipedia using our [Scraper APIs](https://oxylabs.io/products/scraper-api). 

### How it works

You can get Wikipedia results by providing your own URLs to our service. We can return the HTML of any Wikipedia page you like.

#### Python code example

The example below shows how to get a Wikipedia page result in HTML format.

```python
import requests
from pprint import pprint

# Structure payload.
payload = {
   'source': 'universal',
   'url': 'https://en.wikipedia.org/wiki/Oxylabs',
   'context': [
        {'key': 'follow_redirects', "value": True}
      ],
}

# Get response.
response = requests.request(
    'POST',
    'https://realtime.oxylabs.io/v1/queries',
    auth=('YOUR_USERNAME', 'YOUR_PASSWORD'), #Your credentials go here
    json=payload,
)

# Instead of response with job status and results url, this will return the
# JSON response with results.
pprint(response.json())
```

Find code examples for other programming languages [**here**](), and see Oxylabs documentation for more details.

#### Output example

```json
{
  "results": [
    {
      "content": "<!doctype html>\n<html lang=\"en\">\n<head>
      ...
      </script></body>\n</html>\n",
      "created_at": "2023-06-28 07:56:42",
      "updated_at": "2023-06-28 07:56:43",
      "page": 1,
      "url": "https://en.wikipedia.org/wiki/Oxylabs",
      "job_id": "7079729310709324801",
      "status_code": 200
    }
  ]
}
```

Also, check [this tutorial](https://oxylabs.io/blog/how-to-scrape-wikipedia) for a detailed walkthrough.

Using Oxylabs’ Wikipedia Scraper, you can leave infrastructure handling and web unblocking to us and focus on the most important things – data and its analysis.

If you have questions, please contact us at hello@oxylabs.io or via live chat on our website.
