package models

// PageSummary represents a summary of HTML content.
type PageSummary struct {
	Version             string         // Version represents the HTML version.
	Title               string         // Title represents the HTML document title.
	HeaderCount         map[string]int // HeaderCount stores the count of each header type.
	ExternalLinks       []string       // ExternalLinks stores external links found in the HTML.
	InternalLinks       []string       // InternalLinks stores internal links found in the HTML.
	AccessibleLinkCount int            // AccessibleLinkCount stores the count of accessible links.
	HasLoginPage        bool           // HasLoginPage indicates if the HTML contains a login page.
}

// NewPageSummary creates a new instance of PageSummary with initialized fields.
func NewPageSummary() *PageSummary {
	return &PageSummary{
		HeaderCount:   make(map[string]int),
		ExternalLinks: make([]string, 0),
		InternalLinks: make([]string, 0),
	}
}

// WithVersion sets the HTML version of the PageSummary.
func (h *PageSummary) WithVersion(version string) {
	h.Version = version
}

// WithTitle sets the HTML document title of the PageSummary.
func (h *PageSummary) WithTitle(title string) {
	h.Title = title
}

// CountHeader increments the count for the specified header type in the PageSummary.
func (h *PageSummary) CountHeader(header string) {
	h.HeaderCount[header]++
}

// AddExternalLink adds an external link to the ExternalLinks slice of the PageSummary.
func (h *PageSummary) AddExternalLink(externalLink string) {
	h.ExternalLinks = append(h.ExternalLinks, externalLink)
}

// AddInternalLink adds an internal link to the InternalLinks slice of the PageSummary.
func (h *PageSummary) AddInternalLink(internalLink string) {
	h.InternalLinks = append(h.InternalLinks, internalLink)
}

// CountLinkAsAccessible increments the count of accessible links in the PageSummary.
func (h *PageSummary) CountLinkAsAccessible() {
	h.AccessibleLinkCount++
}