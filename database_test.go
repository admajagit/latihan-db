package golangmysql

import (
	"database/sql"
	"testing"

	_"github.com/go-sql-driver/mysql"
)


// membuat koneksi databse atau kita sebut database interface 
// untuk melakaukan koneksi ke database di golang kita bisa membuat objek sql.DB menggunkan function sql.Open(drive,data source name)
// untuk driver bisa diisi mysql
// sedangkan untuk datasourcename tiap database punya cara oenulisan maisng masing
// jika objek sql.db tidak digunkan maka sebaknya kita tutup dengan close()

// mengkoneksikan db dengan golang

// GetConnect membuka koneksi ke database MySQL dan mengembalikan objek *sql.DB
// yang dapat digunakan untuk berinteraksi dengan database.
// TestConnection adalah fungsi untuk menguji koneksi ke database MySQL.
func TestConnection(t *testing.T) {
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


// dengan database pooling kita bisa menentukan jumlah dan makskimal kenokesi yang dibuat oleh golang sehingga tidak
// membanjjiri koneksi ke dattabase karena ada batas maksimal koneksi yang bisa ditangani oleh database 
// idle = koneksi yang tdiak digunkan

// pengaturan database pooling 

// (db)setMaxldleConns(10) =maksimal koneksi 
 // jika koneksi yang tidak digunkan suudah 10  maka akan tertolak koneksi baru dan menunggu salah satu dari 10 aktif kebali agar bisa menambah
// koneksi baru
// (db)setMaxOpenConns(100) =jadi jika kita punya koneksi maksimal 100 maka yang koneksi lebih dari itu  maka akan terrolak 
// db.SetConnMaxIdleTime(5 *time.Minute) koneksi yang tidak digunkan akan terhapus setelah 5 menit
// (db)setConnmaxlifetime(50 *time.minute) = jadi baats waktu penggunaan koneksi jika lebih dari itu maka semua koneksi akan terputus dari db


// import (
// 	"database/sql"
// 	"testing"
// 	_ "github.com/go-sql-driver/mysql" // Import driver MySQL, diperlukan untuk inisialisasi tetapi tidak langsung digunakan dalam kode
// )

// func TestConnection(t *testing.T) {
// 	// Membuka koneksi ke database MySQL
// 	// Format koneksi: "user:password@tcp(host:port)/dbname"
// 	// Dalam hal ini, usernya adalah "root", tanpa password, hostnya "localhost", port "3306", dan nama database "belajar-golang-database"
// 	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar-golang-database")

// 	// Jika ada error saat membuka koneksi, maka program akan panic dan menampilkan error tersebut
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Menutup koneksi ke database ketika fungsi ini selesai dijalankan
// 	// defer memastikan db.Close() dipanggil terakhir setelah semua operasi dalam fungsi ini selesai
// 	defer db.Close()
// }