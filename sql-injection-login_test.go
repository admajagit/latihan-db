package golangmysql

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// contoh sql injection

func TestAji(t *testing.T) {
	// Buka koneksi ke database
	db := GetConnect()

	defer db.Close()

	// buat context
	ctx := context.Background()

	// buat data palsu
	username := "joko"
	password := "admin"

	// Query menggunakan prepared statement untuk mengambil username saja
    scriptkun := "SELECT username FROM login WHERE username= '" + username + "' AND password= '" + password

	roes, err := db.QueryContext(ctx, scriptkun, username, password)
	if err != nil {
		panic(err)
	}
	defer roes.Close()

	if roes.Next() {
		var username string
		err := roes.Scan(&username)

		if err != nil {
			panic(err)
		}
		fmt.Println("berhasil login", username)
	} else {
		fmt.Println("Gagal login")
	}
}


// jadi sql injeection terjadi ketika inputan langsung masuk ke query tanpa context
func TestAfuh(t *testing.T) {
    // Buka koneksi ke database
    db := GetConnect()
    defer db.Close()

    // buat context
    ctx := context.Background()

	// 
    // input data palsu yang bisa menyebabkan SQL Injection yaitu dengan 1=1 akan bernilai benar karena 1 true 
	// jadi semua yang kita input akan benar hasilnya meskipun berbeda dengan data di datbase  
    username := "ahmad' OR '1'='1"
    password := "adkin' OR '1'='1"

    // Query menggunakan string query rentan
    scriptkun := "SELECT password  FROM login WHERE username= '" + username + "' AND password= '" + password + "'"
    roes, err := db.QueryContext(ctx, scriptkun)
    if err != nil {
        panic(err)
    }
    defer roes.Close()

    if roes.Next() {
        var dbUsername string
        err := roes.Scan(&dbUsername)

        if err != nil {
            panic(err)
        }
        fmt.Println("berhasil login", dbUsername)
    } else {
        fmt.Println("Gagal login")
    }
}


// solusinya sql injection yaitu menmasukan inputan ke context dulu baru ke query
// untuk menandai jika sebuah query membutuhkan data kita bisa gunkan tanda ?
func TestAhuji(t *testing.T) {
    // Buka koneksi ke database
    db := GetConnect()
    defer db.Close()

    // buat context
    ctx := context.Background()

    username := "amoi"
    password := "lindhu"

    // buat query dan gunkan ? sebagai parameter 
    scriptkun := "SELECT username FROM login WHERE username = ? AND password = ? LIMIT 1"

	// buat context  untuk memasukan ke parameter query
    row := db.QueryRowContext(ctx, scriptkun, username, password)

    var dbUsername string
	
    err := row.Scan(&dbUsername)
    if err != nil {
        panic(err)
    } else {
        fmt.Println("berhasil login", dbUsername)
    }
}
// import (
//     "context"
//     "database/sql"
//     "fmt"
//     "log"
//     "testing"
//     _ "github.com/go-sql-driver/mysql"
// )

// // Fungsi untuk mendapatkan koneksi database
// func GetConnect() *sql.DB {
//     connStr := "youruser:yourpassword@tcp(127.0.0.1:3306)/golang_database"
//     db, err := sql.Open("mysql", connStr)
//     if err != nil {
//         log.Fatal(err)
//     }
//     return db
// }

// func TestAhuji(t *testing.T) {
//     // Buka koneksi ke database menggunakan fungsi GetConnect
//     db := GetConnect()
//     // Menutup koneksi ke database setelah fungsi selesai dijalankan
//     defer db.Close()

//     // Membuat context untuk eksekusi query
//     ctx := context.Background()

//     // Data login palsu yang akan diuji
//     username := "amoi"
//     password := "lindhu"

//     // Buat query menggunakan prepared statement untuk menghindari SQL injection
//     // Menggunakan ? sebagai parameter placeholder
//     scriptkun := "SELECT username FROM login WHERE username = ? AND password = ? LIMIT 1"
    
//     // Menjalankan query dengan context dan parameter yang diberikan
//     row := db.QueryRowContext(ctx, scriptkun, username, password)

//     // Variabel untuk menampung hasil query
//     var dbUsername string
    
//     // Memindai hasil query ke variabel dbUsername
//     err := row.Scan(&dbUsername)
//     if err != nil {
//         // Jika terjadi error saat scan, panic untuk menunjukkan error
//         panic(err)
//     } else {
//         // Jika tidak ada error, berarti login berhasil
//         fmt.Println("berhasil login", dbUsername)
//     }
// }

// func main() {
//     // Memanggil fungsi test
//     TestAhuji(nil)
// }


// prepare statement
// query atau exec dengan parameter 
// saat kita menggunkan function query atau exec yang menggunkan parameter,sebenarnya implementasi dibawahnya menggunkan prepare statement 
// jadi tahapan pertama statement nya disiapkan terlebih dahulu,setelah itu baru isi dengan parameeter 
// kadang ada kasus kita ingin melakukan beberapa hal yanng sama sekaligus,hanya berbeda parameternya,misal insert data langsung banyak 
// pembuatan prepare statement bisa dilakukan denengan manual,tanpa harus menggunkan query parameeter atau exec parameter 


// 
func TestLoka(t *testing.T)  {
    
}