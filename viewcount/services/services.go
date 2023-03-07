package services

import (
	"sync"
	"viewcount/platform"
)

var videos = platform.NewContainer().(*platform.Container)

func GetVideos() []platform.Video {
	return videos.GetAll()
}

func PostVideo(r *platform.Video) string {
	video := platform.Video{
		Id:    RandStringRunes(10),
		Title: r.Title,
		Views: 0,
	}

	videos.AddVideo(video)

	return video.Id
}

var mut = &sync.RWMutex{}

func GetViews(id string) (int, bool) {
	result := 0
	mut.RLock()
	video, ok := videos.GetVideo(id)
	if ok {
		result = video.Views
	}
	mut.RUnlock()
	return result, ok
}

func IncViews(id string) bool {
	mut.RLock()
	video, ok := videos.GetVideo(id)
	mut.RUnlock()

	if ok {
		mut.Lock()
		video.Views++
		mut.Unlock()
	}

	return ok
}
