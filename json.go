package shortening

type CuttlyCreate struct {
  Url StatusCreate `json:"url"`
}

type StatusCreate struct {
  Status    int8   `json:"status"`
  FullLink  string `json:"fullLink"`
  Date      string `json:"date"`
  ShortLink string `json:"shortLink"`
  Title     string `json:"title"`
}

type CuttlyStats struct {
  Stats StatusStats `json:"stats"`
}

type StatusStats struct {
  Status     int    `json:"status"`
  Clicks     int    `json:"clicks"`
  Date       string `json:"date"`
  Title      string `json:"title"`
  FillLink   string `json:"fullLink"`
  ShortLink  string `json:"shortLink"`
  Facebook   int    `json:"facebook"`
  Twitter    int    `json:"twitter"`
  Pinterest  int    `json:"pinterest"`
  Instagram  int    `json:"instagram"`
  GooglePlus int    `json:"googlePlus"`
  Linkedin   int    `json:"linkedin"`
  Rest       int    `json"rest"`
  Bots       string `json:"bots"`
}