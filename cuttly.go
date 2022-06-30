package shorteningo

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func (c *Configuration) CuttlyNew(url string) (CuttlyCreate, error) {
  var cc CuttlyCreate

  address := fmt.Sprintf("https://cutt.ly/api/api.php?key=%s&short=%s", c.Key, url)
  req, err := http.NewRequest(http.MethodGet, address, nil)
  if err != nil {
    return cc, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return cc, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return cc, err
  }

  if err := json.Unmarshal(data, &cc); err != nil {
    return cc, err
  }

  switch cc.Url.Status {
  case 1: cc.Url.Title = LinkShortenig[1]
  case 2: cc.Url.Title = LinkShortenig[2]
  case 3: cc.Url.Title = LinkShortenig[3]
  case 4: cc.Url.Title = LinkShortenig[4]
  case 5: cc.Url.Title = LinkShortenig[5]
  case 6: cc.Url.Title = LinkShortenig[6]
  case 7: // OK
  }

  return cc, nil
}

func (c *Configuration) CuttlyAnalytics(short string) (CuttlyStats, error) {
  var cs CuttlyStats

  address := fmt.Sprintf("https://cutt.ly/api/api.php?key=%s&stats=%s", c.Key, short)
  req, err := http.NewRequest(http.MethodGet, address, nil)
  if err != nil {
    return cs, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return cs, err
  }
  defer resp.Body.Close()

  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return cs, err
  }

  if err := json.Unmarshal(data, &cs); err != nil {
    return cs, err
  }

  return cs, nil
}