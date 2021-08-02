package youtube

import (
	"art/bots/akina/internal/datalab"
	"art/bots/akina/internal/db"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const maxRes = 5

func CheckTheBestVideo(category *datalab.Category) {
	var bestVideosFromAllCategoryChannels []*video

	// 1) Для каждого канала нахожу лучшее видео.
	for _, channel := range category.BestChannelsArr {
		bestVideosFromAllCategoryChannels = append(bestVideosFromAllCategoryChannels, getBestVideo(channel))
	}

	// 2) Нахожу лучшее видео для всех каналов этой категории.
	vdo := getMostViewedVideo(&bestVideosFromAllCategoryChannels)

	// 3) Если это видео еще не было самым популярным, сохранить в dl и в db.
	if vdo.ID.VideoID != category.LastWatchedVideo {
		category.LastWatchedVideo = vdo.ID.VideoID
		datalab.GetDl().Akina.SendMsg(datalab.ToOurGroup, fmt.Sprintf("%s%s", datalab.GetDl().Youtube.MainPartOfYbUrl, category.LastWatchedVideo))
		err := db.UpdateDbRecord(category)
		if err != nil {
			log.Println(err)
		}
	}

}

func getBestVideo(chanId string) *video {
	httpRequest, err := makeRequest(chanId)
	if err != nil {
		return nil
	}

	allVideos, err := getVideos(httpRequest.URL.String())
	if err != nil {
		return nil
	}

	var videoIdsForNewRequest [maxRes]string
	for i, v := range allVideos.VideosArr {
		videoIdsForNewRequest[i] = v.ID.VideoID
	}

	httpRequest, err = makeRequestForAdditionalInfo(&videoIdsForNewRequest)
	if err != nil {
		return nil
	}

	videosAddInfo, err := getAdditionalInfoForVideos(httpRequest.URL.String())
	if err != nil {
		return nil
	}

	for idx := range allVideos.VideosArr {
		v := allVideos.VideosArr[idx]
		v.AddInfo = videosAddInfo.VideoInfosArr[idx]
	}

	return getMostViewedVideo(&allVideos.VideosArr)
}

func getMostViewedVideo(videoArr *[]*video) (rVideo *video) {
	rVideo = (*videoArr)[0]
	for idx := range *videoArr {
		v := (*videoArr)[idx]
		if v.AddInfo.ViewCount() > rVideo.AddInfo.ViewCount() {
			rVideo = v
		}
	}
	return
}

func getAdditionalInfoForVideos(url string) (*videoInfos, error) {
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var VideoInfos videoInfos
	_ = json.Unmarshal(body, &VideoInfos)

	return &VideoInfos, nil
}

func getVideos(url string) (*videos, error) {
	httpResponse, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer httpResponse.Body.Close()

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}

	var Videos videos
	_ = json.Unmarshal(body, &Videos)

	return &Videos, nil
}

func makeRequest(chanId string) (*http.Request, error) {
	request, err := http.NewRequest("GET", datalab.GetDl().Youtube.MainPartOfHttpUrl, nil)
	if err != nil {
		return nil, err
	}

	/*https://www.googleapis.com/youtube/v3/search
	?part=id&channelId=123&maxResults=5&order=date&type=video&videoType=any&key=123*/
	query := request.URL.Query()
	query.Add("part", "id")
	query.Add("channelId", chanId)
	query.Add("maxResults", strconv.Itoa(maxRes))
	query.Add("order", "date")
	query.Add("type", "video")
	query.Add("videoType", "any")
	query.Add("key", datalab.GetDl().Youtube.Token)
	request.URL.RawQuery = query.Encode()

	return request, nil
}

func makeRequestForAdditionalInfo(urls *[maxRes]string) (*http.Request, error) {
	request, err := http.NewRequest("GET", datalab.GetDl().Youtube.MainPartOfHttpUrl2, nil)
	if err != nil {
		return nil, err
	}

	/*https://www.googleapis.com/youtube/v3/videos
	?part=statistics&id=HDKAakn1RuI%2CiJfoP0kE_18%2CHDKAakn1RuI&key=123*/
	query := request.URL.Query()
	query.Add("part", "statistics")
	for i := 0; i < maxRes; i++ {
		query.Add("id", urls[i])
	}
	query.Add("key", datalab.GetDl().Youtube.Token)

	request.URL.RawQuery = query.Encode()
	log.Println(request.URL.String())

	return request, nil
}
