package yaml_test

import (
	"testing"

	"github.com/emosbaugh/yaml"
)

func TestComments(t *testing.T) {
	cases := []struct {
		name     string
		data     string
		expected string
	}{
		{
			name:     "basic",
			data:     "key1: value1",
			expected: "key1: value1\n",
		},
		{
			name: "missing values",
			data: `key1:
  # a wild comment
  key2:

`,
			expected: `key1:
  # a wild comment
  key2: null
`,
		},
		{
			name: "comment indent",
			data: `# comment 0

# comment 000
zkey1: zvalue1
# comment 1
  # comment 2
ykey2:
  # comment 3
  # comment 4
  xkey3: xvalue3
  wkey4:
    # comment 41
    - item1
    # comment 42
    - item2
    # comment 43
  # comment 5
    # comment 6
wkey5: |
  555

  555
# comment 7`,
			expected: `# comment 0

# comment 000
zkey1: zvalue1
# comment 1
  # comment 2
ykey2:
  # comment 3
  # comment 4
  xkey3: xvalue3
  wkey4:
    # comment 41
  - item1
    # comment 42
  - item2
    # comment 43
  # comment 5
    # comment 6
wkey5: |
  555

  555

# comment 7
`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var i yaml.MapSlice
			var unmarshaler yaml.CommentUnmarshaler
			err := unmarshaler.Unmarshal([]byte(c.data), &i)
			if err != nil {
				t.Errorf("CommentUnmarshaler.Unmarshal() error = %v", err)
			}
			// b, _ := json.MarshalIndent(i, "", "  ")
			// fmt.Printf("+++ %s \n", b)
			got, err := yaml.Marshal(i)
			if err != nil {
				t.Errorf("Marshal() error = %v", err)
			}
			if string(got) != c.expected {
				t.Errorf("CommentUnmarshaler.Unmarshal() = %v, want %v", string(got), c.expected)
			}
		})
	}
}
