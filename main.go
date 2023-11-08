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

type TrackFile struct {
	gorm.Model
	Url        string
	FileTypeID int
	FileType   FileType `gorm:"foreignKey:FileTypeID"`
	Size       int
	Checksum   string
}

type Bank struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Email string
	BuyID int
	Buy   Bank `gorm:"foreignKey:BuyID"`
}

type Artist struct {
	gorm.Model
	UserID int
	User   User `gorm:"foreignKey:UserID"`
	SellID int
	Sell   Bank `"gorm:foreignKey:SellID"`
}

type Contract struct {
	gorm.Model
	Parties      []User `gorm:"many2many:contract_parties;"`
	MasterFileID int
	MasterFile   TrackFile `"gorm:foreignKey:MasterFileIDf"`
}

type Distribution struct {
	gorm.Model
	Name      string
	UploadURL string
	UploadID  int
	Status    Upload `"gorm:foreignKey:UploadID"`
}

type Upload struct {
	PercentComplete int
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
	TrackID  int
	Track    Track `gorm:"foreignkey=TrackID"`
	Consumer User
	Owner    Artist
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
