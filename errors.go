package shorteningo

import "errors"

var LinkShortenig = map[int]string{
	1: "The shortened link comes from the domain that shortens the link, i.e. the link has already been shortened",
	2: "The entered link is not a link",
	3: "The preferred link name / alias is already taken",
	4: "Invalid API key",
	5: "The link has not passed the validation. Includes invalid characters",
	6: "The link provided is from a blocked domain",
}

var (
	unauthorized       = errors.New("Unauthorized")
	method_not_allowed = errors.New("Method not allowed")
	invalid_url        = errors.New("Invalid URL or Alias field is too long/small (minimum 6 characters)")
	something_wrong    = errors.New("Something went wrong")
)