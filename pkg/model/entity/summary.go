package entity

type ResourceSummary struct {
	TotalProjects     int `json:"totalProjects"`
	TotalTeams        int `json:"totalTeams"`
	TotalInstances    int `json:"totalInstances"`
	TotalAPIKeys      int `json:"totalAPIKeys"`
	TotalVolumeClaims int `json:"totalVolumeClaims"`
}
