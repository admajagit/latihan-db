package golangmysql


import (
	"database/sql"
	"testing"

	_"github.com/go-sql-driver/mysql"
)

func TestKon(t *testing.T) {
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

    // Pastikan untuk menutup koneksi ke database setelah fungsi selesai dijalankan.
    defer db.Close()
}