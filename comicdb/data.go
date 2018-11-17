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

//AddComicTitle ...
func (db *ComicDB) AddComicTitle(t *ComicTitle) {
	db.Titles = append(db.Titles, t)
}

//AddComicIssue ...
func (t *ComicTitle) AddComicIssue(i *ComicIssue) {
	t.Issues = append(t.Issues, i)
}

//NewComicTitle ...
func NewComicTitle(t string, v int, y int) (ct *ComicTitle) {
	ct.Title = t
	ct.Volume = v
	ct.Year = y
	return ct

}

//NewComicIssue ...
func NewComicIssue(i int, vt string, b int) (ci *ComicIssue) {
	ci.IssueNumber = i
	ci.VanityTitle = vt
	ci.Box = b
	return ci
}
