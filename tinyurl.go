package shorteningo

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "net/url"
  "math/rand"
  "time"
  "bytes"
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
    var random = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    var lettres = []rune(random)
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

func (c *Configuration) TinyNew(url string, t TinyResult) (TinyUrlCreate, int, error) {
  var tc TinyUrlCreate

  address := fmt.Sprintf(
    "https://api.tinyurl.com/create?api_token=%s&url=%s&%s",
    c.Key, url, t.OptionsTiny(),
  )
  req, err := http.NewRequest(http.MethodPost, address, nil)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return tc, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return tc, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return tc, resp.StatusCode, unauthorized
  case 405: return tc, resp.StatusCode, method_not_allowed
  case 422: return tc, resp.StatusCode, invalid_url
  default: return tc, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return tc, 1, err
  }

  if err = json.Unmarshal(data, &tc); err != nil {
    return tc, 1, err
  }

  return tc, 0, nil
}

// NOTES: new_domain not working, need to subscribe and check
func (c *Configuration) TinyUpdate(domain, alias, new_alias string) (TinyUrlUpdate, int, error) {
  var tu TinyUrlUpdate

  settings, err := json.Marshal(map[string]interface{} {
    "domain": domain,
    "alias": alias,
    "new_alias": new_alias,
    "new_stats": true,
  })
  if err != nil {
    return tu, 1, err
  }

  address := fmt.Sprintf("https://api.tinyurl.com/update?api_token=%s", c.Key)
  req, err := http.NewRequest(http.MethodPatch, address, bytes.NewBuffer(settings))
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return tu, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return tu, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return tu, resp.StatusCode, unauthorized
  case 405: return tu, resp.StatusCode, method_not_allowed
  case 422: return tu, resp.StatusCode, invalid_url
  default: return tu, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return tu, 1, err
  }

  if err = json.Unmarshal(data, &tu); err != nil {
    return tu, 1, err
  }

  return tu, 0, nil
}

func (c *Configuration) TinyChange(url, domain, alias string) (TinyUrlChange, int, error) {
  var tc TinyUrlChange

  settings, err := json.Marshal(map[string]interface{} {
    "url": url,
    "domain": domain,
    "alias": alias,
  })
  if err != nil {
    return tc, 1, err
  }

  address := fmt.Sprintf("https://api.tinyurl.com/change?api_token=%s", c.Key)
  req, err := http.NewRequest(http.MethodPatch, address, bytes.NewBuffer(settings))
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return tc, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return tc, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return tc, resp.StatusCode, unauthorized
  case 405: return tc, resp.StatusCode, method_not_allowed
  case 422: return tc, resp.StatusCode, invalid_url
  default: return tc, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return tc, 1, err
  }

  if err = json.Unmarshal(data, &tc); err != nil {
    return tc, 1, err
  }

  return tc, 0, nil
}

func (c *Configuration) TinyInformation(domain, alias string) (TinyUrlReceiveInformation, int, error) {
  var tri TinyUrlReceiveInformation

  address := fmt.Sprintf("https://api.tinyurl.com/alias/%s/%s?api_token=%s", domain, alias, c.Key)
  req, err := http.NewRequest(http.MethodGet, address, nil)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return tri, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return tri, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return tri, resp.StatusCode, unauthorized
  case 405: return tri, resp.StatusCode, method_not_allowed
  case 422: return tri, resp.StatusCode, invalid_url
  default: return tri, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return tri, 1, err
  }

  if err = json.Unmarshal(data, &tri); err != nil {
    return tri, 1, err
  }

  return tri, 0, nil
}

// NOTES: delete method not working, need to subscribe and check
func (c *Configuration) TinyDelete(domain, alias string) (TinyUrlDelete, int, error) {
  var td TinyUrlDelete

  address := fmt.Sprintf("https://api.tinyurl.com/alias/%s/%s?api_token=%s", domain, alias, c.Key)
  req, err := http.NewRequest(http.MethodDelete, address, nil)
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return td, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return td, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return td, resp.StatusCode, unauthorized
  case 405: return td, resp.StatusCode, method_not_allowed
  case 422: return td, resp.StatusCode, invalid_url
  default: return td, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return td, 1, err
  }

  if err = json.Unmarshal(data, &td); err != nil {
    return td, 1, err
  }

  return td, 0, nil
}

// NOTES: archive method not working, need to subscribe and check
func (c *Configuration) TinyArchive(domain, alias string) (TinyUrlUpdate, int, error) {
  var ta TinyUrlUpdate

  settings, err := json.Marshal(map[string]interface{} {
    "domain": domain,
    "alias": alias,
  })
  if err != nil {
    return ta, 1, err
  }

  address := fmt.Sprintf("https://api.tinyurl.com/archive?api_token=%s", c.Key)
  req, err := http.NewRequest(http.MethodPatch, address, bytes.NewBuffer(settings))
  req.Header.Set("Content-Type", "application/json")
  if err != nil {
    return ta, 1, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return ta, 1, err
  }
  defer resp.Body.Close()

  switch resp.StatusCode {
  case 200: // OK
  case 401: return ta, resp.StatusCode, unauthorized
  case 405: return ta, resp.StatusCode, method_not_allowed
  case 422: return ta, resp.StatusCode, invalid_url
  default: return ta, resp.StatusCode, something_wrong
  }

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return ta, 1, err
  }

  if err = json.Unmarshal(data, &ta); err != nil {
    return ta, 1, err
  }

  return ta, 0, nil
}