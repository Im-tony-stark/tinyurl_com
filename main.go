package main

import (
  "fmt"
  sh "github.com/vexilology/shorteningo"
)

func conf() *sh.Configuration { return sh.Add("PRIVATE_KEY") }

func main() {
  c := conf()

  // Example 1 - Create
  // empty Domain: tiny.one
  // empty Alias:  generate random string
  e := &sh.Options{Domain: "", Alias: ""}
  tn, status_code, err := c.TinyNew("YOUR_URL", e)
  if err != nil {
    fmt.Println(status_code, err)
    return
  }
  fmt.Println(
    tn.Data.Domain, tn.Data.Alias, tn.Data.Tiny_url,
    tn.Data.Deleted, tn.Data.Archived,
    tn.Data.Analytics.Enable, tn.Data.Analytics.Public,
    tn.Data.Tags, tn.Data.Created_at,
    tn.Data.Url, tn.Code, tn.Errors,
  )
