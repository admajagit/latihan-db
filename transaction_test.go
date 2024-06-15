package golangmysql

import (
	"context"
	"fmt"
	"testing"
)

// database transaction merupakan fitur andalan di database adalah transaction

// secara default semua perintah sql yang kita kirimkan menggunkan golang akan otomatis di commit atau autocoommit
// namun kita bisa menggunkan fitur transaksi sehingga sql yang kita kirim tidak secraa otomatis di commit di database
// untuk memulai transaksi,kita bisa menggunkan function(DB)Begin(),dimana akan menghasilkan struct TX yang merupakan representasi Transaction
// struct Tx ini yang kita gunakan sebagai pengganti  db untuk melakukan transaksi dimana hampir semua function di db ada TX,Exec,Query,Prepare
// setelah proses transaksi selesai,kita bisa gunkan function (Tx) commit()untuk melakukan commit atau rollback()

// Commit: Saat Anda melakukan commit, semua perubahan yang dilakukan selama transaksi tersebut diterapkan permanen ke database. Setelah commit, perubahan tersebut tidak bisa dibatalkan menggunakan rollback.
// Rollback: Rollback hanya bisa dilakukan selama transaksi masih berlangsung dan belum di-commit. Rollback mengembalikan database ke keadaan sebelum transaksi dimulai, membatalkan semua perubahan yang dilakukan selama transaksi.
func TestCommit(t *testing.T) {
	db := GetConnect()

	ctix := context.Background()

	tx, error := db.Begin()

	if error != nil {
		panic(error)
	}


	scriptku := "INSERT INTO latihan(id, nama) VALUES(?,?)" // string SQL untuk insert data
	idne := 11
	namane := "ki hubner joko"

	_, err := tx.ExecContext(ctix, scriptku, idne, namane)
	if err != nil {
		panic(err)
	}


	sekrip := "INSERT INTO latihan(id, nama) VALUES('10','Aminkun')" // string SQL untuk insert data
	_, errorcuy := tx.ExecContext(ctix, sekrip)
	if errorcuy != nil {
		panic(errorcuy)
	}

	fmt.Println("berhasil insert")


	defer db.Close()
	tx.Commit()

	
}

