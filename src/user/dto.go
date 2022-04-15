package user

// DTO in
type UserRequest struct {
	ID           int64  `json:"id"`
	Password     string `json:"password"`
	NamaKantor   string `json:"nama_kantor"`
	NamaAdmin    string `json:"nama_admin"`
	NomorKantor  string `json:"nomor_kantor"`
	AlamatKantor string `json:"alamat_kantor"`
}

//DTO Out
type UserResponse struct {
	NamaKantor   string `json:"nama_kantor"`
	NamaAdmin    string `json:"nama_admin"`
	AlamatKantor string `json:"alamat_kantor"`
	NomorKantor  string `json:"nomor_kantor"`
}
