package youtube

import "strconv"

type video struct {
	Kind string `json:"kind"`
	Etag string `json:"etag"`
	ID   struct {
		Kind    string `json:"kind"`
		VideoID string `json:"videoId"`
	} `json:"id"`
	AddInfo videoInfo
}

type videos struct {
	VideosArr []*video `json:"items"`
}

/*----------------------------------------------------------------------------------------------*/

type videoInfo struct {
	Kind       string `json:"kind"`
	Etag       string `json:"etag"`
	ID         string `json:"id"`
	Statistics struct {
		ViewCount     string `json:"viewCount"`
		LikeCount     string `json:"likeCount"`
		DislikeCount  string `json:"dislikeCount"`
		FavoriteCount string `json:"favoriteCount"`
		CommentCount  string `json:"commentCount"`
	} `json:"statistics"`
}

func (a *videoInfo) ViewCount() int {
	i, err := strconv.Atoi(a.Statistics.ViewCount)
	if err != nil {
		return 0
	}
	return i
}

type videoInfos struct {
	VideoInfosArr [maxRes]videoInfo `json:"items"`
}
