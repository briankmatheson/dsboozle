
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

type Track struct {
	gorm.Model
	Name       string
	MasterFile TrackFile
	TrackFiles []TrackFile
	Artist     string
	Composer   string
	Genre      string
	Year       int
	Lyrics     string
	Comment    string
}

type Collection struct {
	gorm.Model
	Name   string
	Type   string
	Tracks []Tracks
	Artist string
	Genre  string
	Year   string
	Comment
}