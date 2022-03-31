package entity

type MethodPolicy struct {
	Name     string `json:"name,omitempty"`
	Category string `json:"category,omitempty"`
	Allow    bool   `json:"allow" bson:"allow"`
}
