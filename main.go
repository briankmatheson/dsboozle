package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type FileType struct {
	gorm.Model
	Encoding string
	Bitrate  string
	Lossless bool
}

func CreateFileType(db *gorm.DB, filetype *FileType) error {
	result := db.Create(&filetype)
	return result.Error
}
func ReadFileType(db *gorm.DB, id uint) (FileType, error) {
	var filetype FileType
	result := db.First(&filetype, id)
	return filetype, result.Error
}
func UpdateFileType(db *gorm.DB, filetype *FileType) error {
	result := db.Save(&filetype)
	return result.Error
}
func DeleteFileType(db *gorm.DB, id uint) error {
	result := db.Delete(&FileType{}, id)
	return result.Error
}

type TrackFile struct {
	gorm.Model
	Url        string
	FileTypeID int
	FileType   FileType `gorm:"foreignKey:FileTypeID"`
	Size       int
	Checksum   string
}

func CreateTrackFile(db *gorm.DB, trackfile *TrackFile) error {
	result := db.Create(&trackfile)
	return result.Error
}
func ReadTrackFile(db *gorm.DB, id uint) (TrackFile, error) {
	var trackfile TrackFile
	result := db.First(&trackfile, id)
	return trackfile, result.Error
}
func UpdateTrackFile(db *gorm.DB, trackfile *TrackFile) error {
	result := db.Save(&trackfile)
	return result.Error
}
func DeleteTrackFile(db *gorm.DB, id uint) error {
	result := db.Delete(&TrackFile{}, id)
	return result.Error
}

type Bank struct {
	gorm.Model
	Name string
}

func CreateBank(db *gorm.DB, bank *Bank) error {
	result := db.Create(&bank)
	return result.Error
}
func ReadBank(db *gorm.DB, id uint) (Bank, error) {
	var bank Bank
	result := db.First(&bank, id)
	return bank, result.Error
}
func UpdateBank(db *gorm.DB, bank *Bank) error {
	result := db.Save(&bank)
	return result.Error
}
func DeleteBank(db *gorm.DB, id uint) error {
	result := db.Delete(&Bank{}, id)
	return result.Error
}

type User struct {
	gorm.Model
	Email string
	BuyID int
	Buy   Bank `gorm:"foreignKey:BuyID"`
}

func CreateUser(db *gorm.DB, user *User) error {
	result := db.Create(&user)
	return result.Error
}
func ReadUser(db *gorm.DB, id uint) (User, error) {
	var user User
	result := db.First(&user, id)
	return user, result.Error
}
func UpdateUser(db *gorm.DB, user *User) error {
	result := db.Save(&user)
	return result.Error
}
func DeleteUser(db *gorm.DB, id uint) error {
	result := db.Delete(&User{}, id)
	return result.Error
}

type Artist struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserID"`
	SellID int
	Sell   Bank `"gorm:foreignKey:SellID"`
}

func CreateArtist(db *gorm.DB, artist *Artist) error {
	result := db.Create(&artist)
	return result.Error
}
func ReadArtist(db *gorm.DB, id uint) (Artist, error) {
	var artist Artist
	result := db.First(&artist, id)
	return artist, result.Error
}
func UpdateArtist(db *gorm.DB, artist *Artist) error {
	result := db.Save(&artist)
	return result.Error
}
func DeleteArtist(db *gorm.DB, id uint) error {
	result := db.Delete(&Artist{}, id)
	return result.Error
}

type Contract struct {
	gorm.Model
	Parties      []User `gorm:"many2many:contract_parties;"`
	MasterFileID int
	MasterFile   TrackFile `"gorm:foreignKey:MasterFileID"`
}

func CreateContract(db *gorm.DB, contract *Contract) error {
	result := db.Create(&contract)
	return result.Error
}
func ReadContract(db *gorm.DB, id uint) (Contract, error) {
	var contract Contract
	result := db.First(&contract, id)
	return contract, result.Error
}
func UpdateContract(db *gorm.DB, contract *Contract) error {
	result := db.Save(&contract)
	return result.Error
}
func DeleteContract(db *gorm.DB, id uint) error {
	result := db.Delete(&Contract{}, id)
	return result.Error
}

type Distribution struct {
	gorm.Model
	Name      string
	UploadURL string
	UploadID  int
	Status    Upload `"gorm:foreignKey:UploadID"`
}

func CreateDistribution(db *gorm.DB, distribution *Distribution) error {
	result := db.Create(&distribution)
	return result.Error
}
func ReadDistribution(db *gorm.DB, id uint) (Distribution, error) {
	var distribution Distribution
	result := db.First(&distribution, id)
	return distribution, result.Error
}
func UpdateDistribution(db *gorm.DB, distribution *Distribution) error {
	result := db.Save(&distribution)
	return result.Error
}
func DeleteDistribution(db *gorm.DB, id uint) error {
	result := db.Delete(&Distribution{}, id)
	return result.Error
}

type Upload struct {
	PercentComplete int
}

func CreateUpload(db *gorm.DB, upload *Upload) error {
	result := db.Create(&upload)
	return result.Error
}
func ReadUpload(db *gorm.DB, id uint) (Upload, error) {
	var upload Upload
	result := db.First(&upload, id)
	return upload, result.Error
}
func UpdateUpload(db *gorm.DB, upload *Upload) error {
	result := db.Save(&upload)
	return result.Error
}
func DeleteUpload(db *gorm.DB, id uint) error {
	result := db.Delete(&Upload{}, id)
	return result.Error
}

type Track struct {
	gorm.Model
	Name         string
	UploadFileID int
	UploadFile   TrackFile `"gorm:"foreignKey:UploadFileID"`
	MasterFileID int
	MasterFile   TrackFile   `gorm:"foreignKey:MasterFileID"`
	TrackFiles   []TrackFile `gorm:"many2many:track_trackfiles;"`
	ArtistID     int
	Artist       Artist `gorm:"foreignKey:ArtistID"`
	ComposerID   int
	Composer     Artist `gorm:"foreignKey:ComposerID"`
	ContractID   int
	Contract     Contract `gorm:"foreignKey:ContractID"`
	Genre        string
	Year         int
	Lyrics       string
	Comment      string
}

func CreateTrack(db *gorm.DB, track *Track) error {
	result := db.Create(&track)
	return result.Error
}
func ReadTrack(db *gorm.DB, id uint) (Track, error) {
	var track Track
	result := db.First(&track, id)
	return track, result.Error
}
func UpdateTrack(db *gorm.DB, track *Track) error {
	result := db.Save(&track)
	return result.Error
}
func DeleteTrack(db *gorm.DB, id uint) error {
	result := db.Delete(&Track{}, id)
	return result.Error
}

type Collection struct {
	gorm.Model
	Name    string
	Type    string
	Tracks  []Track `gorm:"many2many:collection_tracks;"`
	Artist  string
	Genre   string
	Year    string
	Comment string
}

type Entitlement struct {
	gorm.Model
	CollectionID int
	Collection   Collection `gorm:"foreignKey=CollectionID"`
	Consumer     User
	Owner        Artist
}

func main() {
	db, err := gorm.Open(
		postgres.New(postgres.Config{DSN: "host=192.168.1.31 user=dsboozle password=gr8passwd dbname=dsboozle port=5432 sslmode=disable TimeZone=America/New_York"}),
		&gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Track{})

	// Create
	track := Track{Name: "New Track", Genre: "Pop"}
	CreateTrack(db, &track)

	// Read
	readTrack, err := ReadTrack(db, track.ID)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Track:", readTrack)
	}

	// Update
	track.Genre = "Rock"
	UpdateTrack(db, &track)

	// Delete
	DeleteTrack(db, track.ID)
}
