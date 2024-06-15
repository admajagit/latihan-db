package golangmysql

import (
	"context" // import paket context untuk pengelolaan konteks dalam operasi database
	"fmt"     // import paket fmt untuk fungsi output seperti Println
	"strconv"
	"testing" // import paket testing untuk menyediakan dukungan test

	
)

// digunakan untuk mengeksekusi pernyataan SQL dalam konteks tertentu, dan untuk menangani kesalahan yang mungkin terjadi selama eksekusi tersebut

// printah sql dari golang
// seperti ini menambah/insert data sql dari golang

func TestElek(t *testing.T)  {
	// karena kita sudah buat koneksi sebelumnya kita bisa terhubung ke mysql
	
	db := GetConnect() // memanggil fungsi GetConnect untuk mendapatkan koneksi database

	// jangan lupa pakai close karena kita di test

	//buat context kosong
	contox := context.Background() // membuat context kosong yang tidak memiliki deadline, cancel, atau nilai tambahan

	// buat script untuk insert data
	scriptku := "INSERT INTO latihan(id, nama) VALUES(3,'alok')" // string SQL untuk insert data

	// eksekusi script SQL dengan context yang telah dibuat
	_, err := db.ExecContext(contox, scriptku) // menjalankan perintah SQL dengan context

	if err != nil {
		panic(err) // jika terjadi error, hentikan provvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvvgram dengan panic
	}

	// menutup koneksi database setelah semua operasi selesai
	defer db.Close() 

	fmt.Println("insert data selesai") // mencetak pesan bahwa proses insert data telah selesai

}

func TestEko(t *testing.T)  {
	
	
	db := GetConnect() 

	
	username := "amain"
	password := "adminkun"

	contox := context.Background()

	scriptku := "INSERT INTO login(username, password) VALUES(?,?)" 

	_, err := db.ExecContext(contox, scriptku,username,password) 

	if err != nil {
		panic(err) 
	}
	defer db.Close() 

	fmt.Println("insert data selesai")

}


// func TestPrepare(t *testing.T)  {
	
	
// 	db := GetConnect()

// 	contox := context.Background()

// 	scriptku := "INSERT INTO latihan_factorial(email, koment) VALUES(?,?)" 

// 	for i := 0; i < 10; i++ {
// 		email := "ajixs" + strconv.Itoa(i) + "@gmail.com"
// 		komen := "komentar ke = " + strconv.Itoa(i)

// 		result,err := statement.ExecContext(contox,scriptku)

// 		if err!= nil {
// 			panic(err)
// 		}

// 		fmt.Println("coment id = " ,email)
// 	}


// 	defer db.Close()
// }

func TestPreparee(t *testing.T) {
	db := GetConnect()
	

	contox := context.Background()

	scriptku := "INSERT INTO latihan_factorial(email, koment) VALUES(?,?)"

	// Menyiapkan prepared statement
	statement, err := db.PrepareContext(contox, scriptku)
	if err != nil {
		panic(err)
	}
	

	// Menggunakan prepared statement untuk memasukkan data
	for i := 0; i < 10; i++ {
		email := "ajixs" + strconv.Itoa(i) + "@gmail.com"
		komen := "komentar ke = " + strconv.Itoa(i)

		_, err := statement.ExecContext(contox, email, komen)
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id = ", email)
	}

	defer statement.Close()
	defer db.Close()
}

// repared statement adalah fitur dalam database yang memungkinkan kita untuk menyiapkan suatu pernyataan SQL sekali saja
//  kemudian mengeksekusinya berkali-kali dengan parameter yang berbeda.
func TestPrepareen(t *testing.T) {
    // Membuat koneksi ke database
    db := GetConnect()
    
    // Membuat context untuk mengatur eksekusi query
    contexto := context.Background()

    // Query SQL untuk memasukkan data ke tabel latihan_factorial
    scriptku := "INSERT INTO latihan_factorial(email, koment) VALUES(?,?)"

    // Menyiapkan prepared statement dengan menggunakan query di atas
    statement, err := db.PrepareContext(contexto, scriptku)
    if err != nil {
        // Jika ada kesalahan saat menyiapkan prepared statement, hentikan eksekusi dan tampilkan error
        panic(err)
    }
    
    // Menggunakan prepared statement untuk memasukkan data ke dalam tabel
    for i := 0; i < 10; i++ {
        // Membuat email dan komentar untuk setiap iterasi
        email := "ajil" + strconv.Itoa(i) + "@gmail.com"
        komen := "komentar ke = " + strconv.Itoa(i)

        // Mengeksekusi prepared statement dengan parameter email dan komentar
        _, err := statement.ExecContext(contexto, email, komen)
        if err != nil {
            // Jika ada kesalahan saat eksekusi, hentikan eksekusi dan tampilkan error
            panic(err)
        }

        // Menampilkan email yang dimasukkan
        fmt.Println("comment email = ", email)
    }

    // Menutup prepared statement setelah selesai
    defer statement.Close()
    
    // Menutup koneksi database setelah selesai
    defer db.Close()
}
