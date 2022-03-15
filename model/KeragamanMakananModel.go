package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type KeragamanMakanan struct {
	IdKeragaman		uint		`gorm:"primary_key"`
	IdUser			uint
	IdSubjek		uint
	NamaSubjek		string
	Tanggal			time.Time
	Jawaban1		bool
	Jawaban2		bool
	Jawaban3		bool
	Jawaban4		bool
	Jawaban5		bool
	Jawaban6		bool
	Jawaban7		bool
	Jawaban8		bool
	Jawaban9		bool

	Nasi			bool
	Roti			bool
	Sereal			bool
	Jagung			bool
	Ubi				bool
	Singkong		bool
	Kentang			bool
	Ketela			bool
	Sagu			bool
	Mi				bool

	Bayam			bool
	Kangkung		bool
	Brokoli			bool
	Sawi			bool
	Selada			bool
	Buncis			bool
	KacangPanjang	bool
	Pakcoy			bool

	Wortel			bool
	Pepaya			bool
	Semangka		bool
	Tomat			bool
	Paprika			bool
	Labu			bool
	Jambu			bool

	Kubis			bool
	Lobak			bool
	Timun			bool
	Terong			bool
	Asparagus		bool
	Taoge			bool

	Apel			bool
	Jeruk			bool
	Pisang			bool
	Anggur			bool
	Nanas			bool
	Alpukat			bool
	Manggis			bool
	Belimbing		bool
	Mangga			bool
	Durian			bool
	Rambutan		bool
	Sirsak			bool
	Salak			bool

	Sapi			bool
	Kambing			bool
	Ayam			bool
	Bebek			bool
	Burung			bool
	Ikan			bool
	Cumi			bool
	Udang			bool
	Gurita			bool
	Kepiting		bool

	TelurAyam		bool
	TelurBebek		bool
	TelurPuyuh		bool

	Tahu			bool
	Tempe			bool
	KacangKedelai	bool
	KacangTanah		bool
	KacangHijau		bool
	KacangMete		bool
	KacangKenari	bool
	KacangAlmond	bool
	KacangPolong	bool
	KacangMerah		bool

	Susu			bool
	Keju			bool
	Yoghurt			bool
}

func InitTableKeragamanMakanan(DBCon *gorm.DB) {
	DBCon.AutoMigrate(&KeragamanMakanan{})
	DBCon.Model(&KeragamanMakanan{}).AddForeignKey("id_subjek", "data_diris(id_data_diri)", "CASCADE",
		"CASCADE")
}

func CreateKeragaman(keragaman KeragamanMakanan, db *gorm.DB) KeragamanMakanan {
	err := db.Create(&keragaman).Error
	if err != nil{
		panic(err)
	}

	return keragaman
}

func GetKeragaman(idUser uint, db *gorm.DB) []KeragamanMakanan {
	var daftarKeragaman []KeragamanMakanan

	err := db.Where("id_user = ?", idUser).Find(&daftarKeragaman).Error
	if err != nil{
		panic(err)
	}

	return daftarKeragaman
}

func GetKeragamanBySubjek(idUser uint, idSubjek uint, db *gorm.DB) []KeragamanMakanan {
	var daftarKeragaman []KeragamanMakanan

	err := db.Where("id_user = ? AND id_subjek = ?", idUser, idSubjek).Find(&daftarKeragaman).Error
	if err != nil{
		panic(err)
	}

	return daftarKeragaman
}