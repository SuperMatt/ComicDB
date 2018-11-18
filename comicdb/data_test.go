package comicdb

import (
	"testing"
	"time"
)

func TestComicDB(t *testing.T) {
	var db ComicDB
	db.Date = time.Now()
	title := NewComicTitle("The Batman Who Laughs", 1, 2018)
	issue := NewComicIssue(1, "", 1)

	title.AddComicIssue(issue)
	db.AddComicTitle(title)

	p := db.String()
	t.Log(p)
}
