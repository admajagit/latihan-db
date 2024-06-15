package golangmysql

import (
	"database/sql"
	"time"
)

// jangan pakai db.close karena itu hanya untuk testing kalau dsini akan langsung ke close koneksi ke db


// GetConnect membuka koneksi ke database MySQL dan mengembalikan objek *sql.DB
// yang dapat digunakan untuk berinteraksi dengan database.
func GetConnect() *sql.DB {
	// Membuka koneksi ke database MySQL.
	// Format DSN (Data Source Name): username:password@protocol(address)/dbname
	// Dalam contoh ini:
	// - username: "root"
	// - password: (kosong, tidak ada password)
	// - protocol: "tcp"
	// - address: "localhost:3306" (localhost di port 3306)
	// - dbname: "golang_database"
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_database")

	// Jika ada kesalahan saat membuka koneksi, hentikan eksekusi dan tampilkan kesalahan.
	if err != nil {
		panic(err)
	}
	  
db.SetMaxIdleConns(10)
db.SetMaxOpenConns(100)
db.SetConnMaxIdleTime(5 *time.Minute)
db.SetConnMaxLifetime(1 *time.Minute)

	// Mengembalikan objek *sql.DB untuk digunakan di tempat lain dalam aplikasi.
	// Jangan menutup koneksi di sini karena koneksi akan digunakan oleh pemanggil fungsi ini.
	return db
}
