# Soal Teori

### Soal
Rancangkanlah diagram database untuk aplikasi rumah makan. Jelaskan teknologi yang akan dipakai untuk aplikasi ini dan mengapa anda memilih teknologi tersebut. Kebutuhan:
1. Aplikasi ini bisa memasukkan pesanan-pesanan makanan pelanggan
2. Aplikasi ini bisa mengeluarkan struk pembelian
3. Aplikasi ini bisa mengeluarkan laporan penghasilan mingguan dan bulanan
4. Aplikasi ini bisa mengeluarkan laporan stok

Selain kebutuhan pokok di atas, silahkan tambahkan ide original anda untuk membuat aplikasi lebih baik.

### Jawaban

#### Diagram: 
![Diagram Database Aplikasi Rumah Makan](https://github.com/kobulwidodo/qoin-techincal-test/assets/53964878/f60047a7-a687-439d-8713-f794a8b8e4dc)

#### Teknologi:
1. **MySQL**: 
   - MySQL adalah RDBMS open source yang sangat cocok untuk pengembangan aplikasi. 
   - Dengan dukungan komunitas yang besar, MySQL menjadi pilihan yang tepat karena keamanan dan skalabilitasnya.

2. **Redis**:
   - Redis dapat meningkatkan efisiensi pengambilan data, terutama untuk data yang jarang berubah seperti menu.
   - Implementasi Redis dapat mengurangi beban pada sistem, terutama dalam proses query database.

#### Ide Tambahan:
1. **Dashboard Analitik Penjualan**:
   - Memberikan analisis penjualan untuk membantu pengambilan keputusan bisnis.
   - Menampilkan data seperti menu mana yang penjualannya meningkat dan pola waktu pesanan masuk.

2. **Grafik Penjualan**:
   - Menyediakan visualisasi jumlah pesanan berdasarkan kategori harian.
   - Membantu pelaku bisnis memahami perkembangan penjualan.
