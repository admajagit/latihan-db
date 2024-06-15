package golangmysql

import (
	"database/sql"
	
)

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")

	// Jika ada kesalahan saat membuka koneksi, hentikan eksekusi dan tampilkan kesalahan.
	if err != nil {
		panic(err)
		// Mengembalikan objek *sql.DB untuk digunakan di tempat lain dalam aplikasi.
		// Jangan menutup koneksi di sini karena koneksi akan digunakan oleh pemanggil fungsi ini.

	}
	return db
}
