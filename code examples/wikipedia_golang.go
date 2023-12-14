package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    const Username = "YOUR_USERNAME"
    const Password = "YOUR_PASSWORD"

    payload := map[string]interface{}{
	"source": "universal",
	"url": "https://en.wikipedia.org/wiki/Oxylabs",
	"context": []map[string]interface{}{
             {"key": "follow_redirects", value: true},
         }
    }

    jsonValue, _ := json.Marshal(payload)

    client := &http.Client{}
    request, _ := http.NewRequest("POST",
        "https://realtime.oxylabs.io/v1/queries",
        bytes.NewBuffer(jsonValue),
    )

    request.SetBasicAuth(Username, Password)
    response, _ := client.Do(request)

    responseText, _ := ioutil.ReadAll(response.Body)
    fmt.Println(string(responseText))
}
