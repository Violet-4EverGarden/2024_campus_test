package solution

import "fmt"

type Url struct {
	// http or https
	scheme string
	host   string
	port   int
	// path to the resource
	path string
	// Like: key1=value1&key2=value2
	queryParams string
}

func (c *Url) Build() string {
	port := ""
	if c.port != 0 {
		port = fmt.Sprintf(":%d", c.port)
	}
	qp := ""
	if c.queryParams != "" {
		qp = "?" + c.queryParams
	}

	return fmt.Sprintf("%s://%s%s%s%s",
		c.scheme,
		c.host,
		port,
		c.path,
		qp,
	)
}

type IUrlBuilder interface {
	// Set scheme to "https", otherwise leave http
	Https() IUrlBuilder
	// Set host
	Host(string) IUrlBuilder
	// Set port number, or leave as 0
	Port(int) IUrlBuilder
	// Set path
	Path(string) IUrlBuilder
	// Set params as a map of strings
	QueryParams(map[string]string) IUrlBuilder
	// Returns Url object with defined fields
	GetUrl() *Url
}

type UrlBuilder struct {
	url *Url
}

// Returns builder instance
func NewUrlBuilder() IUrlBuilder {
	return &UrlBuilder{url: &Url{scheme: "http"}}
}

func (b *UrlBuilder) Https() IUrlBuilder {
	b.url.scheme = "https"
	return b
}

func (b *UrlBuilder) Host(host string) IUrlBuilder {
	b.url.host = host
	return b
}

func (b *UrlBuilder) Port(port int) IUrlBuilder {
	b.url.port = port
	return b
}

func (b *UrlBuilder) Path(path string) IUrlBuilder {
	b.url.path = path
	return b
}

func (b *UrlBuilder) QueryParams(params map[string]string) IUrlBuilder {
	query := ""
	for key, value := range params {
		if query != "" {
			query += "&"
		}
		query += fmt.Sprintf("%s=%s", key, value)
	}
	b.url.queryParams = query
	return b
}

func (b *UrlBuilder) GetUrl() *Url {
	return b.url
}
