package golangmysql

import (
	"context"   // Mengimpor paket context untuk pengelolaan konteks dalam operasi database
	"fmt"       // Mengimpor paket fmt untuk fungsi output seperti Println
	"log"       // Mengimpor paket log untuk logging error dan informasi lainnya
	"testing"   // Mengimpor paket testing untuk menyediakan dukungan test
)

func TestAmbil(t *testing.T) {
	db := GetConnect() // Memanggil fungsi GetConnect untuk mendapatkan koneksi database
	defer db.Close() // Menutup koneksi database setelah semua operasi selesai

	contox := context.Background() // Membuat context kosong yang tidak memiliki deadline, cancel, atau nilai tambahan
	scriptku := "SELECT id, nama FROM latihan" // String SQL untuk query data

	row, err := db.QueryContext(contox, scriptku) // Menjalankan query SQL dengan context
	if err != nil {
		log.Panic(err) // Jika terjadi error saat query, log error dan hentikan program
	}
	defer row.Close() // Menutup rows setelah semua operasi selesai

	// Iterasi melalui setiap baris yang dihasilkan oleh query
	for row.Next() {
		var id int
		var nama string

		error := row.Scan(&id, &nama) // Mengambil data dari setiap baris dan menampilkannya ke variabel id dan nama
		if err != nil {
			log.Fatal(error) // Jika terjadi error saat scanning, log error dan hentikan program
		}
		fmt.Println(id, nama) // Mencetak id dan nama dalam satu baris
	}
}