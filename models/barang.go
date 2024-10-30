package models

type Barang struct {
	ID_barang   int    `json:"id_barang"`
	Nama_barang string `json:"nama_barang"`
	Stok        int    `json:"stok"`
	ID_jenis    int    `json:"id_jenis"`
}

type GetAll struct {
	NO                int    `json:"no"`
	Nama_barang       string `json:"nama_barang"`
	Jenis_barang      string `json:"jenis_barang"`
	Jumlah_terjual    string `json:"jumlah_terjual"`
	Tanggal_transaksi string `json:"tanggal_transaksi"`
}
