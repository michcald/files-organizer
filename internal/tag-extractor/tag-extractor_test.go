package tag_extractor_test

import (
	"testing"

	tag_extractor "github.com/michcald/files-organizer/internal/tag-extractor"
)

func TestDisassebleString(t *testing.T) {
	tests := []struct {
		s            string
		expectedTags []string
	}{
		{
			"hello! there",
			[]string{
				"hello there",
				"hello",
				"there",
			},
		},
		{
			"hello  there", // multiple blank spaces
			[]string{
				"hello there",
				"hello",
				"there",
			},
		},
		{
			"hello!  there", // multiple blank spaces with special char
			[]string{
				"hello there",
				"hello",
				"there",
			},
		},
		{
			",.$",
			[]string{},
		},
		{
			"hello world my name is matt",
			[]string{
				"world",
				"my",
				"name",
				"is",
				"hello world my name is matt",
				"hello world my name is",
				"hello world my name",
				"hello world my",
				"hello world",
				"hello",
				"matt",
				"is matt",
				"name is matt",
				"my name is matt",
				"world my name is matt",
			},
		},
	}

	e := tag_extractor.NewSimple()

	for i, tt := range tests {
		tags := e.Extract(tt.s)

		if len(tags) != len(tt.expectedTags) {
			t.Errorf("test %d: expected len %d. got=%d", i, len(tt.expectedTags), len(tags))
		}

		for i, expectedTag := range tt.expectedTags {
			found := false

			for _, tag := range tags {
				if expectedTag == tag {
					found = true
					break
				}
			}

			if !found {
				t.Errorf("test %d: expected tag name `%s`. got=nothing", i, expectedTag)
			}
		}
	}
}
