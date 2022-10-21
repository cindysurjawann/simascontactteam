package zoomhistory

type ZoomHistoryRequest struct {
	Nama       string `json:"nama" binding:"required"`
	Email      string `json:"email" binding:"required"`
	Kategori   string `json:"kategori" binding:"required"`
	Keterangan string `json:"keterangan" binding:"required"`
	Lokasi     string `json:"lokasi"`
}
