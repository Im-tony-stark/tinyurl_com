package shorteningo

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
  FullLink   string `json:"fullLink"`
  ShortLink  string `json:"shortLink"`
  Facebook   int    `json:"facebook"`
  Twitter    int    `json:"twitter"`
  Pinterest  int    `json:"pinterest"`
  Instagram  int    `json:"instagram"`
  GooglePlus int    `json:"googlePlus"`
  Linkedin   int    `json:"linkedin"`
  Rest       int    `json:"rest"`
  Bots       string `json:"bots"`
}

type TinyUrlCreate struct {
  Data   DataCreate `json:"data"`
  Code   int        `json:"code"`
  Errors []string   `json:"errors"`
}

// Deleted -> Expires_at: nil
type DataCreate struct {
  Domain     string          `json:"domain"`
  Alias      string          `json:"alias"`
  Deleted    bool            `json:"deleted"`
  Archived   bool            `json:"archived"`
  Analytics  AnalyticsCreate `json:"analytics"`
  Tags       []string        `json:"tags"`
  Created_at string          `json:"created_at"`
  Tiny_url   string          `json:"tiny_url"`
  Url        string          `json:"url"`
}

type AnalyticsCreate struct {
  Enable bool `json:"enable"`
  Public bool `json:"public"`
}

type TinyUrlUpdate struct {
  Data   DataUpdate `json:"data"`
  Code   int        `json:"code"`
  Errors []string   `json:"errors"`
}

type DataUpdate struct {
  Domain    string          `json:"domain"`
  Alias     string          `json:"alias"`
  Deleted   bool            `json:"deleted"`
  Archived  bool            `json:"archived"`
  Tags      []string        `json:"tags"`
  Analytics AnalyticsCreate `json:"analytics"`
  Tiny_url  string          `json:"tiny_url"`
}

type TinyUrlChange struct {
  Data   DataChange `json:"data"`
  Code   int        `json:"code"`
  Errors []string   `json:"errors"`
}

type DataChange struct {
  Url string `json:"url"`
}

type TinyUrlReceiveInformation struct {
  Data   DataReceiveInformation `json:"data"`
  Code   int                    `json:"code"`
  Errors []string               `json:"errors"`
}

type DataReceiveInformation struct {
  Domain    string                 `json:"domain"`
  Alias     string                 `json:"alias"`
  Deleted   bool                   `json:"deleted"`
  Archived  bool                   `json:"archived"`
  Tags      []string               `json:"tags"`
  Analytics AnalyticsCreate        `json:"analytics"`
  Tiny_url  string                 `json:"tiny_url"`
  User      UserReceiveInformation `json:"user"`
  Team      TeamReceiveInformation `json"team"`
  Hits      int                    `json:"hits"`
  Url       string                 `json:"url"`
}

type UserReceiveInformation struct {
  Name  string `json:"name"`
  Email string `json:"email"`
}

type TeamReceiveInformation struct {
  Name string `json:"name"`
  Slug string `json:"slug"`
}

type TinyUrlDelete struct {
  Data   DataDelete `json:"data"`
  Code   int        `json:"code"`
  Errors []string   `json:"errors"`
}

type DataDelete struct {
  Domain     string          `json:"domain"`
  Alias      string          `json:"alias"`
  Deleted    bool            `json:"deleted"`
  Archived   bool            `json:"archived"`
  Tags       []string        `json:"tags"`
  Analytics  AnalyticsCreate `json:"analytics"`
  Tiny_url   string          `json:"tiny_url"`
  Url        string          `json:"url"`
}