package model

import (
	"github.com/jinzhu/gorm"
)

type Makanan struct {
	IdMakanan		uint `gorm:"primary_key"`
	NamaMakanan		string
	Jenis			int
	Src				string
}

func InitTableMakanan(DBCon *gorm.DB) {
	if !DBCon.HasTable(&Makanan{}) {
		DBCon.AutoMigrate(&Makanan{})
		var daftarMakanan []Makanan

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Biskuit", Jenis: 1, Src: "makanan/Biskuit.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Jagung", Jenis: 1, Src: "makanan/Jagung.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kentang", Jenis: 1, Src: "makanan/Kentang.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Mie", Jenis: 1, Src: "makanan/Mie.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Nasi Putih", Jenis: 1, Src: "makanan/Nasi Putih.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Roti", Jenis: 1, Src: "makanan/Roti.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Sereal", Jenis: 1, Src: "makanan/Sereal.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Singkong", Jenis: 1, Src: "makanan/Singkong.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Ubi", Jenis: 1, Src: "makanan/Ubi.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Bayam", Jenis: 2, Src: "makanan/Bayam.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Brokoli", Jenis: 2, Src: "makanan/Brokoli.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Buncis", Jenis: 2, Src: "makanan/Buncis.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kacang Panjang", Jenis: 2, Src: "makanan/Kacang Panjang.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kangkung", Jenis: 2, Src: "makanan/Kangkung.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Pakcoy", Jenis: 2, Src: "makanan/Pakcoy.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Sawi", Jenis: 2, Src: "makanan/Sawi.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Selada", Jenis: 2, Src: "makanan/Selada.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Pepaya", Jenis: 3, Src: "makanan/Pepaya.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Semangka", Jenis: 3, Src: "makanan/Semangka.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Tomat", Jenis: 3, Src: "makanan/Tomat.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Wortel", Jenis: 3, Src: "makanan/Wortel.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kubis", Jenis: 4, Src: "makanan/Kubis.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Taoge", Jenis: 4, Src: "makanan/Taoge.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Terong", Jenis: 4, Src: "makanan/Terong.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Timun", Jenis: 4, Src: "makanan/Timun.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Alpukat", Jenis: 5, Src: "makanan/Alpukat.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Anggur", Jenis: 5, Src: "makanan/Anggur.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Apel", Jenis: 5, Src: "makanan/Apel.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Belimbing", Jenis: 5, Src: "makanan/Belimbing.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Jeruk", Jenis: 5, Src: "makanan/Jeruk.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Manggis", Jenis: 5, Src: "makanan/Manggis.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Nanas", Jenis: 5, Src: "makanan/Nanas.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Pisang", Jenis: 5, Src: "makanan/Pisang.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Ayam", Jenis: 6, Src: "makanan/Ayam.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Cumi", Jenis: 6, Src: "makanan/Cumi.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Daging", Jenis: 6, Src: "makanan/Daging.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Ikan", Jenis: 6, Src: "makanan/Ikan.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Udang", Jenis: 6, Src: "makanan/Udang.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Telur Asin", Jenis: 7, Src: "makanan/Telur Asin.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Telur Mata Sapi", Jenis: 7, Src: "makanan/Telur Mata Sapi.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kacang Kedelai", Jenis: 8, Src: "makanan/Kacang Kedelai.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kacang Merah", Jenis: 8, Src: "makanan/Kacang Merah.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kacang Polong", Jenis: 8, Src: "makanan/Kacang Polong.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Kacang Tanah", Jenis: 8, Src: "makanan/Kacang Tanah.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Tahu", Jenis: 8, Src: "makanan/Tahu.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Tempe", Jenis: 8, Src: "makanan/Tempe.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Keju", Jenis: 9, Src: "makanan/Keju.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Susu", Jenis: 9, Src: "makanan/Susu.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Yoghurt", Jenis: 9, Src: "makanan/Yoghurt.jpg"})

		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Cakwe", Jenis: 10, Src: "makanan/Cakwe.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Chiki", Jenis: 10, Src: "makanan/Chiki.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Cilok", Jenis: 10, Src: "makanan/Cilok.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Cilor", Jenis: 10, Src: "makanan/Cilor.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Donat", Jenis: 10, Src: "makanan/Donat.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Gulali", Jenis: 10, Src: "makanan/Gulali.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Lolipop", Jenis: 10, Src: "makanan/Lolipop.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Makaroni", Jenis: 10, Src: "makanan/Makaroni.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Permen", Jenis: 10, Src: "makanan/Permen.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Piscok", Jenis: 10, Src: "makanan/Piscok.jpg"})
		daftarMakanan = append(daftarMakanan, Makanan{NamaMakanan: "Risol", Jenis: 10, Src: "makanan/Risol.jpg"})

		for _, makanan := range daftarMakanan {
			DBCon.Create(&makanan)
		}
	} else {
		DBCon.AutoMigrate(&Makanan{})
	}
}


func GetMakananList(db *gorm.DB) []Makanan {
	var daftarMakanan []Makanan

	err := db.Find(&daftarMakanan).Error
	if err != nil {
		panic(err)
	}

	return daftarMakanan
}