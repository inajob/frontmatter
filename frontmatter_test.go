package frontmatter

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestFrontMatter(t *testing.T) {
	bodyData, err := ioutil.ReadFile("testdata/frontmatter/body.md")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile("testdata/frontmatter/normal.md")
	if err != nil {
		t.Error(err)
	}
	front, body, err := ParseFrontMatter(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if _, ok := front["frontmatter"]; !ok {
		t.Error("expected frontmatter to contain frontmatter got nil instead")
	}
	if value, _ := front["frontmatter"]; value != "test" {
		t.Errorf("expected frontmatter to contain frontmatter value  got %s instead", value)
	}
	if body != string(bodyData) {
		t.Errorf("expected %s got %s", string(bodyData), body)
	}
}

func TestFrontMatterNoFrontMatter(t *testing.T) {
	bodyData, err := ioutil.ReadFile("testdata/frontmatter/body.md")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile("testdata/frontmatter/nofrontmatter.md")
	if err != nil {
		t.Error(err)
	}
	front, body, err := ParseFrontMatter(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if l := len(front); l != 0 {
		t.Errorf("expected frontmatter size is 0 got %d instead", l)
	}
	if body != string(bodyData) {
		t.Errorf("expected %s got %s", string(bodyData), body)
	}
}

func TestFrontMatterContainDelim(t *testing.T) {
	bodyData, err := ioutil.ReadFile("testdata/frontmatter/bodycontaindelim.md")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile("testdata/frontmatter/containdelim.md")
	if err != nil {
		t.Error(err)
	}
	front, body, err := ParseFrontMatter(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if _, ok := front["frontmatter"]; !ok {
		t.Error("expected frontmatter to contain frontmatter got nil instead")
	}
	if value, _ := front["frontmatter"]; value != "test" {
		t.Errorf("expected frontmatter to contain frontmatter value  got %s instead", value)
	}
	if body != string(bodyData) {
		t.Errorf("expected %s got %s", string(bodyData), body)
	}
}

func TestFrontMatterInvalidFront(t *testing.T) {
	bodyData, err := ioutil.ReadFile("testdata/frontmatter/bodyinvalidfront.md")
	if err != nil {
		t.Error(err)
	}

	data, err := ioutil.ReadFile("testdata/frontmatter/invalidfront.md")
	if err != nil {
		t.Error(err)
	}
	front, body, err := ParseFrontMatter(bytes.NewReader(data))
	if err != nil {
		t.Error(err)
	}
	if l := len(front); l != 0 {
		t.Errorf("expected frontmatter size is 0 got %d instead", l)
	}
	if body != string(bodyData) {
		t.Errorf("expected %s got %s", string(bodyData), body)
	}
}
