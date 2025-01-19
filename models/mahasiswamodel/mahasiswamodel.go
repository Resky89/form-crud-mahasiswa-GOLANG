package mahasiswamodel

import (
	"database/sql"

	"github.com/jeypc/go_form_mahasiswa_uts/config"
	"github.com/jeypc/go_form_mahasiswa_uts/entities"
)

type MahasiswaModel struct {
	db *sql.DB
}

func New() *MahasiswaModel{
	db, err := config.DBConnection()
	if err != nil{
		panic(err)
	}
	return &MahasiswaModel{db: db}
}

func (m *MahasiswaModel) FindAll(mahasiswa *[]entities.Mahasiswa) error {
	rows, err := m.db.Query("select * from mahasiswa")
	if err != nil{
		return err
	}

	defer rows.Close()

	for rows.Next(){
		var data entities.Mahasiswa
		rows.Scan(
			&data.NPM,
			&data.Nama,
			&data.Prodi,
			&data.Kelas)

		*mahasiswa = append(*mahasiswa, data)
	}
	return nil
}

func (m *MahasiswaModel) Create(mahasiswa *entities.Mahasiswa) error {
	result, err := m.db.Exec("insert into mahasiswa ( nama, prodi, kelas) values(?,?,?)",
		 mahasiswa.Nama, mahasiswa.Prodi, mahasiswa.Kelas)

	if err != nil {
		return err
	}

	lastInsertId, _ := result.LastInsertId()
	mahasiswa.NPM = lastInsertId
	return nil
}

func (m *MahasiswaModel) Find(npm int64, mahasiswa *entities.Mahasiswa) error {
	return m.db.QueryRow("select * from mahasiswa where npm = ?", npm).Scan(
		&mahasiswa.NPM,
		&mahasiswa.Nama,
		&mahasiswa.Prodi,
		&mahasiswa.Kelas)
}

func (m *MahasiswaModel) Update(mahasiswa entities.Mahasiswa) error {

	_, err := m.db.Exec("update mahasiswa set nama = ?, prodi = ?, kelas = ? where npm = ?",
		mahasiswa.Nama, mahasiswa.Prodi, mahasiswa.Kelas,  mahasiswa.NPM)

	if err != nil {
		return err
	}

	return nil
}

func (m *MahasiswaModel) Delete(npm int64) error {
	_, err := m.db.Exec("delete from mahasiswa where npm = ?", npm)
	if err != nil {
		return err
	}
	return nil
}

func (m *MahasiswaModel) SearchByName(name string, mahasiswa *[]entities.Mahasiswa) error {
    rows, err := m.db.Query("SELECT * FROM mahasiswa WHERE nama LIKE ?", "%"+name+"%")
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var data entities.Mahasiswa
        rows.Scan(
            &data.NPM,
            &data.Nama,
            &data.Prodi,
            &data.Kelas,
        )
        *mahasiswa = append(*mahasiswa, data)
    }
    return nil
}