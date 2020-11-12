package models

// GroupLevel struct for GroupLevel
type GroupLevel struct {
	Id          int64  `json:"id,omitempty"`
	ParentId    int64  `json:"parentId,omitempty"`
	SuperParent bool   `json:"superParent,omitempty"`
	LevelName   string `json:"levelName,omitempty"`
	Recursable  bool   `json:"recursable,omitempty"`
	Group       bool   `json:"group,omitempty"`
	Center      bool   `json:"center,omitempty"`
	New         bool   `json:"new,omitempty"`
}
