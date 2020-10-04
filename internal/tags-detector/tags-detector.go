package detector

import (
	"strings"
)

type Tag struct {
	Name  string
	Count uint16
}

type TagsDetector struct {
}

func New() *TagsDetector {
	return &TagsDetector{}
}

func (d *TagsDetector) Detect(elems []string) []Tag {
	return nil
}

func (d *TagsDetector) disassembleString(s string) []Tag {
	chunks := strings.Split(s, " ")

	tags := []Tag{}

	for i := len(chunks) - 1; i >= 0; i-- {
		tags = append(tags, Tag{
			Name:  strings.Join(chunks[0:i+1], " "),
			Count: 0,
		})
	}

	return tags
}
