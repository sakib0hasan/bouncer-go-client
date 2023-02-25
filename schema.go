package bouncer

type GetCreditResponse struct {
	Credits int `json:"credits"`
}

type VerifyEmailResponse struct {
	Email  string `json:"email"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
	Domain struct {
		Name       string `json:"name"`
		AcceptAll  string `json:"acceptAll"`
		Disposable string `json:"disposable"`
		Free       string `json:"free"`
	} `json:"domain"`
	Account struct {
		Role        string `json:"role"`
		Disabled    string `json:"disabled"`
		FullMailbox string `json:"fullMailbox"`
	} `json:"account"`
	DNS struct {
		Type   string `json:"type"`
		Record string `json:"record"`
	} `json:"dns"`
	Provider string `json:"provider"`
}
