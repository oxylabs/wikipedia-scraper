curl --user USERNAME:PASSWORD \
'https://realtime.oxylabs.io/v1/queries' \
-H "Content-Type: application/json" \
-d '{"source": "universal", "url":
"https://en.wikipedia.org/wiki/Oxylabs",
"context": [{ "key": "follow_redirects", "value": true}]}'
