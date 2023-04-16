package entity

type StatsRequestInfo struct {
	Method    string
	Endpoint  string
	UserAgent string
}

type StatsResult struct {
	Count           int `json:"count"`
	UniqueUserAgent int `json:"unique_user_agent"`
}
