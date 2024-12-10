package orderedquery

import (
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
	q.pairs = append(q.pairs, [2]string{key, value})
}

func (q *Query) Encode() string {
	var b strings.Builder
	for i, p := range q.pairs {
		if i > 0 {
			b.WriteByte('&')
		}
		b.WriteString(p[0])
		b.WriteByte('=')
		b.WriteString(p[1])
	}
	return b.String()
}
