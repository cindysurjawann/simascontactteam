package managelink

type GetLinkRequest struct {
	LinkType string `json:"linktype"`
}

type UpdateLinkRequest struct {
	LinkType  string `json:"linktype"`
	LinkValue string `json:"linkvalue" binding:"required"`
	UpdatedBy string `json:"updatedby" binding:"required"`
}
