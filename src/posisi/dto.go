package posisi

type PosisiRequest struct {
	NomorResi         string `json:"nomor_resi"`
	LokasiSekarang    string `json:"lokasi_sekarang"`
	LokasiSelanjutnya string `json:"lokasi_selanjutnya"`
}
