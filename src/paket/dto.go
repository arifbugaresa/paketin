package paket

type PaketRequest struct {
	Produk        string `json:"produk"`
	BeratPaket    int64  `json:"berat_paket"`
	NamaPengirim  string `json:"nama_pengirim"`
	NamaPenerima  string `json:"nama_penerima"`
	NomorPenerima string `json:"nomor_penerima"`
	AlamatTujuan  string `json:"alamat_tujuan"`
}
