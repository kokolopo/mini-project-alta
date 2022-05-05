package menu

type InputNewMenu struct {
	Nama      string  `json:"nama" binding:"required"`
	Deskripsi string  `json:"deskripsi" binding:"required"`
	Harga     float64 `json:"harga" binding:"required"`
	Kategori  string  `json:"kategori" binding:"required"`
}
