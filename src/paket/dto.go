package paket

type PaketRequest struct {
	Produk        string `json:"produk"`
	BeratPaket    int64  `json:"berat_paket"`
	NamaPengirim  string `json:"nama_pengirim"`
	NamaPenerima  string `json:"nama_penerima"`
	NomorPenerima string `json:"nomor_penerima"`
	AlamatTujuan  string `json:"alamat_tujuan"`
}

type PaketResponse struct {
	ID        int64  `json:"id"`
	NomorResi string `json:"nomor_resi"`
	Produk    string `json:"produk"`
}

type PaketDetaiResponse struct {
	ID            int64  `json:"id"`
	NomorResi     string `json:"nomor_resi"`
	Produk        string `json:"produk"`
	BeratPaket    int64  `json:"berat_paket"`
	NamaPengirim  string `json:"nama_pengirim"`
	NamaPenerima  string `json:"nama_penerima"`
	NomorPenerima string `json:"nomor_penerima"`
	AlamatTujuan  string `json:"alamat_tujuan"`
	Status        string `json:"status"`
}