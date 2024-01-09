package purl

type purlListData struct {
	Request struct {
		StatusCode int  `json:"status_code"`
		Success    bool `json:"success"`
	} `json:"request"`
	Response struct {
		Message string `json:"message"`
		Purls   []struct {
			Name    string `json:"name"`
			URL     string `json:"url"`
			Counter int    `json:"counter"`
		} `json:"purls"`
	} `json:"response"`
}
