package shorteningo

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "net/url"
  "math/rand"
  "time"
)

type Options struct {
  Domain, Alias string
}

type TinyResult interface {
  OptionsTiny() string
}

func (o *Options) OptionsTiny() string {
  if o.Domain == "" {
    o.Domain = "tiny.one"
  }

  if o.Alias == "" {
    var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
    var lettres = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    b := make([]rune, 8)
    for i := range b {
      b[i] = lettres[seededRand.Intn(len(lettres))]
    }
    o.Alias = string(b)
  }

  v := url.Values{}
  v.Add("domain", o.Domain)
  v.Add("alias", o.Alias)
  return v.Encode()
}

func (c *Configuration) TinyUrlNew(url string, t TinyResult) (TinyurlCreate, error) {
  var tc TinyurlCreate

  address := fmt.Sprintf(
    "https://api.tinyurl.com/create?api_token=%s&url=%s&%s",
    c.Key, url, t.OptionsTiny(),
  )
  req, err := http.NewRequest(http.MethodPost, address, nil)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return tc, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return tc, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return tc, unauthorized
  case 405: return tc, method_not_allowed
  case 422: return tc, invalid_url
  default: return tc, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return tc, err
  }

  if err = json.Unmarshal(data, &tc); err != nil {
    return tc, err
  }

  return tc, nil
}