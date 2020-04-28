package responses

type Link struct {
	Name string
	URL  string
}

func Links(supplied ...Link) map[string]string {

	links := make(map[string]string, 0)
	for _, link := range defaultLinks() {
		links[link.Name] = link.URL
	}

	for _, link := range supplied {
		links[link.Name] = link.URL
	}

	return links
}

func defaultLinks() []Link {
	return []Link{
		{
			Name: "api",
			URL:  Conf.AddressAPI,
		},
		{
			Name: "site",
			URL:  Conf.AddressWebsite,
		},
	}
}
