package meli

type Client struct {
	OAuth      *oauthEndpoints
	Sites      *sitesEndpoints
	Categories *categoriesEndpoints
}

func New(c Config) (*Client, error) {
	//If userCode is not provided, then a generic client is returned.
	//This client can be used only to access public API
	// if strings.Compare(userCode, "") == 0 {
	// 	return nil
	// }
	// if c.UserCode == "" {
	// 	return nil, errors.New("USER_CODE_IS_MISSING")
	// }

	if c.BaseURL == "" {
		c.BaseURL = "https://api.mercadolibre.com"
	}

	client := &Client{
		OAuth:      newOAuthEndpoints(c),
		Sites:      newSitesEndpoints(c),
		Categories: newCategoriesEndpoints(c),
	}

	return client, nil
}
