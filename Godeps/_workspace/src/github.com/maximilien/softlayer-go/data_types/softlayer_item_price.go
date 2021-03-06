package data_types

type SoftLayer_Item_Price struct {
	Id              int        `json:"id"`
	LocationGroupId int        `json:"locationGroupId"`
	Categories      []Category `json:"categories,omitempty"`
	Item            *Item      `json:"item,omitempty"`
}

type Item struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Capacity    string `json:"capacity"`
}

type Category struct {
	CategoryCode string `json:"categoryCode"`
}
