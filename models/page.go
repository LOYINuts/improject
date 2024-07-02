package models

type BasePage struct {
	PageIndex int `form:"page_index"`
	PageSize  int `form:"page_size"`
}
