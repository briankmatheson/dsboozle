package main

import (
	"gorm.io/gorm"
)

type FileType struct {
	encoding string
	bitrate  string
	lossless bool
}

type TrackFile struct {
	url      string
	filetype FileType
	size     int
	checksum string
}

type User struct {
	gorm.Model
	email string
	Buy   bank
}

type Artist struct {
	gorm.Model
	User
	Sell bank
}

type Contract struct {
	gorm.Model
	Parties    []User
	MasterFile TrackFile
}

type Distribution struct {
	gorm.Model
	Name      string
	UploadURL string
	Status    Upload
}

type Upload struct {
	PercentComplete int
}

type Track struct {
	gorm.Model
	Name       string
	UploadFile TrackFile
	MasterFile TrackFile
	TrackFiles []TrackFile
	Artist     Artist
	Composer   Artist
	Contract   Contract
	Genre      string
	Year       int
	Lyrics     string
	Comment    string
}

type Collection struct {
	gorm.Model
	Name    string
	Type    string
	Tracks  []Tracks
	Artist  string
	Genre   string
	Year    string
	Comment string
}

type Entitlement struct {
	gorm.Model
	Track    Track
	Consumer User
	Owner    User
}
