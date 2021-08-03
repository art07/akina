package db

import (
	"art/bots/akina/internal/datalab"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// https://github.com/mattn/go-sqlite3

//goland:noinspection GoUnhandledErrorResult
func InitDbJob(mode int) {
	akinaDb, err := sql.Open("sqlite3", "./akina.db")
	if err != nil {
		log.Fatal(err)
	}
	defer akinaDb.Close()

	dbRows, err := akinaDb.Query("SELECT * FROM last_videos")
	if err != nil {
		log.Fatal(err)
	}
	defer dbRows.Close()

	watchedVideos := make([]*lastWatchedVideo, 0, 5)
	for dbRows.Next() {
		var video lastWatchedVideo
		dbRows.Scan(&video.Id, &video.Category, &video.VideoId)
		watchedVideos = append(watchedVideos, &video)
	}

	switch mode {
	case 0:
		setLastVideos(watchedVideos)
	case 1:
		for _, item := range watchedVideos {
			log.Printf("\nDB INFO:\nId > %d | Categoty > %8s | VideoId > %s\n", item.Id, item.Category, item.VideoId)
		}
	}
}

func setLastVideos(watchedVideos []*lastWatchedVideo) {
	for i, video := range watchedVideos {
		log.Printf("Data from DB <%s> <%s>\n", (*video).Category, (*video).VideoId)
		log.Printf("Data in dl <%s> <%s>\n", datalab.GetDl().Youtube.Categories[i].Name, datalab.GetDl().Youtube.Categories[i].LastWatchedVideo)
		datalab.GetDl().Youtube.Categories[i].LastWatchedVideo = (*video).VideoId
		log.Printf("NEW data in dl <%s> <%s>\n\n", datalab.GetDl().Youtube.Categories[i].Name, datalab.GetDl().Youtube.Categories[i].LastWatchedVideo)
	}
}

//goland:noinspection GoUnhandledErrorResult
func UpdateDbRecord(c *datalab.Category) error {
	akinaDb, err := sql.Open("sqlite3", "akina.db")
	if err != nil {
		return err
	}
	defer akinaDb.Close()

	statement, err := akinaDb.Prepare("UPDATE last_videos SET video_id=? where category=?")
	if err != nil {
		return err
	}

	_, err = statement.Exec(c.LastWatchedVideo, c.Name)
	if err != nil {
		return err
	}

	return nil
}
