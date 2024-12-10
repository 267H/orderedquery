package orderedquery

import (
	"net/url"
	"strings"
)

type Query struct {
	pairs [][2]string
}

func NewQuery(capacity int) *Query {
	return &Query{
		pairs: make([][2]string, 0, capacity),
	}
}

func (q *Query) Add(key, value string) {
	decodedKey, _ := url.QueryUnescape(key)
	decodedValue, _ := url.QueryUnescape(value)
	q.pairs = append(q.pairs, [2]string{decodedKey, decodedValue})
}

func (q *Query) Get(key string) string {
	for _, pair := range q.pairs {
		if pair[0] == key {
			return pair[1]
		}
	}
	return ""
}

func (q *Query) Encode() string {
	var b strings.Builder
	for i, p := range q.pairs {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(url.QueryEscape(p[0]))
		b.WriteByte('=')
		b.WriteString(url.QueryEscape(p[1]))
	}
	return b.String()
}

func ParseQuery(queryString string) (*Query, error) {
	if len(queryString) > 0 && queryString[0] == '?' {
		queryString = queryString[1:]
	}

	if queryString == "" {
		return NewQuery(0), nil
	}

	pairs := strings.Split(queryString, "&")
	q := NewQuery(len(pairs))

	for _, pair := range pairs {
		if pair == "" {
			continue
		}

		kv := strings.SplitN(pair, "=", 2)
		if len(kv) == 1 {
			q.Add(kv[0], "")
		} else {
			q.Add(kv[0], kv[1])
		}
	}

	return q, nil
}
