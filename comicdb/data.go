package comicdb

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

//ComicDB ...
type ComicDB struct {
	Date   time.Time
	Titles []*ComicTitle
}

//ComicTitle ...
type ComicTitle struct {
	Title  string
	Volume int
	Year   int
	Issues []*ComicIssue
}

//ComicIssue ...
type ComicIssue struct {
	IssueNumber int
	VanityTitle string
	Box         int
}

//NewDB ...
func NewDB() *ComicDB {
	var db ComicDB
	db.Date = time.Now()
	return &db
}

//LoadComicDBBytes ...
func LoadComicDBBytes(b []byte) (db *ComicDB, err error) {
	err = json.Unmarshal(b, db)
	return db, err
}

//LoadComicDBFile ...
func LoadComicDBFile(f string) (db *ComicDB, err error) {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		return db, err
	}
	return LoadComicDBBytes(b)
}

//String ...
func (db *ComicDB) String() string {
	b, err := json.Marshal(db)
	if err != nil {
		return err.Error()
	}

	return string(b)
}

//AddComicTitle ...
func (db *ComicDB) AddComicTitle(t *ComicTitle) {
	db.Titles = append(db.Titles, t)
}

//AddComicIssue ...
func (t *ComicTitle) AddComicIssue(i *ComicIssue) {
	t.Issues = append(t.Issues, i)
}

//NewComicTitle ...
func NewComicTitle(t string, v int, y int) *ComicTitle {
	var ct ComicTitle
	ct.Title = t
	ct.Volume = v
	ct.Year = y
	return &ct

}

//NewComicIssue ...
func NewComicIssue(i int, vt string, b int) *ComicIssue {
	var ci ComicIssue
	ci.IssueNumber = i
	ci.VanityTitle = vt
	ci.Box = b
	return &ci
}
