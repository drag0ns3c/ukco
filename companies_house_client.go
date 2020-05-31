package ukco

const (
	API_BASE_URL = "https://api.companieshouse.gov.uk"
)

type CompaniesHouseClient struct {
	Client *Client
}

func (c *CompaniesHouseClient) CompanyProfile(number string) *CompanyProfile {
	return nil
}
