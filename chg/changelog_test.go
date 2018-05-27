package chg

import (
	"bytes"
	"testing"
)

func TestChangelogVersion(t *testing.T) {
	unreleased := &Version{Name: "Unreleased"}
	v123 := &Version{Name: "1.2.3"}

	c := Changelog{}
	c.Versions = append(c.Versions, unreleased)
	c.Versions = append(c.Versions, v123)

	t.Run("version=unreleased", func(t *testing.T) {
		result := c.Version("unreleased")
		if result != unreleased {
			t.Error("Test comparing 'unreleased' version failed")
		}
	})

	t.Run("version=1.2.3", func(t *testing.T) {
		result := c.Version("1.2.3")
		if result != v123 {
			t.Error("Test comparing '1.2.3' version failed")
		}
	})

	t.Run("version=unknown", func(t *testing.T) {
		result := c.Version("unknown")
		if result != nil {
			t.Error("Test comparing 'unknown' version failed")
		}
	})
}

func TestChangelogRenderLinks(t *testing.T) {
	unreleased := &Version{Name: "Unreleased", Link: "http://example.com/unreleased"}
	v123 := &Version{Name: "1.2.3", Link: "http://example.com/1.2.3"}
	v456 := &Version{Name: "4.5.6"}

	c := Changelog{}
	c.Versions = append(c.Versions, unreleased)
	c.Versions = append(c.Versions, v123)
	c.Versions = append(c.Versions, v456)

	expected := "[Unreleased]: http://example.com/unreleased\n[1.2.3]: http://example.com/1.2.3\n"

	var buf bytes.Buffer
	c.RenderLinks(&buf)
	result := string(buf.Bytes())

	if result != expected {
		t.Errorf("TestChangelogRenderLinks fail, expecting %s found %s", expected, result)
	}
}

func TestChangelogRender(t *testing.T) {
	c := Changelog{
		Preamble: "Any paragraph\nto be inserted.\n",
	}

	t.Run("empty-versions", func(t *testing.T) {
		expected := "# Changelog\n\nAny paragraph\nto be inserted.\n"
		var buf bytes.Buffer
		c.Render(&buf)
		result := buf.String()
		if result != expected {
			t.Errorf("TestChangelogRender with empty versions fails, got %s expected %s", result, expected)
		}
	})

	t.Run("with-versions", func(t *testing.T) {
		c.Versions = []*Version{
			&Version{Name: "1.0.0"},
			&Version{Name: "2.0.0"},
		}

		expected := "# Changelog\n\nAny paragraph\nto be inserted.\n\n## 1.0.0\n\n## 2.0.0\n"
		var buf bytes.Buffer
		c.Render(&buf)
		result := buf.String()
		if result != expected {
			t.Errorf("TestChangelogRender with versions fails, got =%s= expected =%s=", result, expected)
		}
	})
}