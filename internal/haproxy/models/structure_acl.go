package models

type ACL struct {
	ACLName   string `json:"acl_name"`
	Criterion string `json:"criterion"`
	Index     int    `json:"index"`
	Value     string `json:"value"`
}
