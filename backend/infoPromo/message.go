package infoPromo

type InfoRequest struct {
	Judul     string `json:"judul" binding:"required"`
	Kategori  string `json:"kategori" binding:"required"`
	Startdate string `json:"startdate" binding:"required"`
	Enddate   string `json:"enddate" binding:"required"`
	Kodepromo string `json:"kodepromo" binding:"required"`
	Foto      string `json:"foto" binding:"required"`
	Deskripsi string `json:"deskripsi" binding:"required"`
	Syarat    string `json:"syarat" binding:"required"`
}