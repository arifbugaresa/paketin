package user

type UserRequest struct {
	ID           int64  `json:"id"`
	Password     string `json:"password"`
	NamaKantor   string `json:"nama_kantor"`
	NamaAdmin    string `json:"nama_admin"`
	NomorKantor  string `json:"nomor_kantor"`
	AlamatKantor string `json:"alamat_kantor"`
}
