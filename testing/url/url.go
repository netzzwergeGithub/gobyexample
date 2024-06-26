package url

import (
	"errors"
	"strings"
)

// A URL represents a parsed URL.
type URL struct{ Scheme, Host, Path string }

// Parse parses a raw url into a URL structure.
func Parse(rawURL string) (*URL, error) {
	scheme, rest, ok := strings.Cut(rawURL, ":")
	if !ok || scheme == "" {
		return nil, errors.New("missing scheme")
	}

	if !strings.HasPrefix(rest, "//") {
		return &URL{Scheme: scheme}, nil
	}

	host, path, _ := strings.Cut(rest[2:], "/")

	u := &URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}

	return u, nil
}

// String reassembles the URL into a URL string.
func (u *URL) String() string {
	if u == nil {
		return ""
	}

	var s strings.Builder

	lenURL := len(u.Scheme) + len("://") +
		len(u.Host) + len("/") +
		len(u.Path)
	s.Grow(lenURL)

	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteByte('/')
		s.WriteString(p)
	}

	return s.String()
}
