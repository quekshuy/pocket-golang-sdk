/*

Takes care of OAuth to the Pocket service.
To use, need to set Pocket API consumer key as environment variable POCKET_API_KEY.

*/
package auth

import (
    "log"
    "os"
    "net/http"
    "net/url"
    "io/ioutil"
)

const (
    ENV_API_KEY = "POCKET_API_KEY"
)

func apiCredentials() string {
    return os.Getenv(ENV_API_KEY)
}

func responseBodyAsValues(r *http.Response) (url.Values, error) {

    body, err := ioutil.ReadAll(r.Body)
    defer r.Body.Close()

    if err != nil {
        return url.Values{}, err
    }

    return url.ParseQuery(string(body))
}

// GetPocketRequestToken will get the first request token, kicks off the authentication
// process.
func GetPocketRequestToken(callbackUrl string) (string) {
    apiKey := apiCredentials()
    resp, err := http.PostForm(
        "https://getpocket.com/v3/oauth/request",
        url.Values{"consumer_key": {apiKey}, "redirect_uri": {callbackUrl}},
    )

    if err != nil {
        log.Fatalf("Error getting code from Pocket: %v", err)
    }
    values, err := responseBodyAsValues(resp)
    return values.Get("code")
}

func GetPocketAccessToken(code string)( string,  string) {
    apiKey := apiCredentials()
    resp, err := http.PostForm(
        "https://getpocket.com/v3/oauth/authorize",
        url.Values{"consumer_key": {apiKey}, "code": {code} },
    )

    if err != nil {
        log.Fatalf("Error getting code from Pocket: %v", err)
    }
    values, err := responseBodyAsValues(resp)
    return values.Get("username"), values.Get("access_token")
}
