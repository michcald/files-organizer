package detector

import "testing"

func TestDisassebleString(t *testing.T) {
	tests := []struct {
		s            string
		expectedTags []Tag
	}{
		{
			"hello world my name is matt",
			[]Tag{
				{
					Name:  "hello world my name is matt",
					Count: 0,
				},
				{
					Name:  "hello world my name is",
					Count: 0,
				},
				{
					Name:  "hello world my name",
					Count: 0,
				},
				{
					Name:  "hello world my",
					Count: 0,
				},
				{
					Name:  "hello world",
					Count: 0,
				},
				{
					Name:  "hello",
					Count: 0,
				},
			},
		},
	}
	// todo tags must only be a string (no counter as not necessary here)

	d := New()

	for i, tt := range tests {
		tags := d.disassembleString(tt.s)

		if len(tags) != len(tt.expectedTags) {
			t.Errorf("test %d: expected len %d. got=%d", i, len(tags), len(tt.expectedTags))
		}

		for i, tag := range tt.expectedTags {
			if tags[i].Name != tag.Name {
				t.Errorf("test %d: expected tag name `%s`. got=`%s`", i, tag.Name, tags[i].Name)
			}
		}
	}
}
