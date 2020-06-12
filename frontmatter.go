package frontmatter

import (
	"bufio"
	"bytes"
	"io"

	"gopkg.in/yaml.v2"
)

// ParseFrontMatter is split frontmatter and document body
func ParseFrontMatter(input io.Reader) (front map[string]interface{}, body string, err error) {
	bufsize := 1024 * 1024
	reader := bufio.NewReaderSize(input, bufsize)
	var buf = make([]byte, 0, 100)
	var delim = []byte{'-', '-', '-'}
	isBody := false
	front = make(map[string]interface{})
ROOT:
	for {
		isPrefix := true
		for isPrefix {
			var line []byte
			line, isPrefix, err = reader.ReadLine()
			if err == io.EOF {
				break ROOT
			} else if err != nil {
				return nil, "", err
			}
			if !isBody {
				if bytes.Equal(line, delim) {
					// End of frontmatter
					isBody = true
					err := yaml.Unmarshal(buf, front)
					if err == nil {
						buf = buf[:0]
					} else {
						buf = append(buf, line...)
						buf = append(buf, '\n')
					}
				} else {
					buf = append(buf, line...)
					buf = append(buf, '\n')
				}
			} else {
				buf = append(buf, line...)
				buf = append(buf, '\n')
			}
		}
	}
	return front, string(buf), nil
}
