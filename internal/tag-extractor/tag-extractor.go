package tag_extractor

import (
	"regexp"
	"strings"
)

type Extractor interface {
	Extract(s string) []string
}

type Simple struct{}

func NewSimple() *Simple {
	return &Simple{}
}

func (e *Simple) Extract(s string) []string {
	// remove any special char
	r := regexp.MustCompile(`[^a-zA-Z\d\s:]`)
	s = r.ReplaceAllString(s, " ")

	// remove multiple contiguous blank spaces
	r = regexp.MustCompile(`\s\s+`)
	s = r.ReplaceAllString(s, " ")

	if s == "" || s == " " {
		return []string{}
	}

	chunks := strings.Split(s, " ")

	tags := chunks
	chunkCnt := len(chunks)

	for i := chunkCnt - 1; i >= 1; i-- {
		tags = append(tags, strings.Join(chunks[0:i+1], " "))
	}

	for i := 1; i < chunkCnt-1; i++ {
		tags = append(tags, strings.Join(chunks[i:chunkCnt], " "))
	}

	return tags
}
