package platform

type videoContainer interface {
	AddVideo(video Video)
	GetAll() []Video
}

type Video struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Views int    `json:"views"`
}

type Container struct {
	videos []Video
	mp     map[string]*Video
}

func NewContainer() videoContainer {
	return &Container{
		videos: []Video{},
		mp:     make(map[string]*Video),
	}
}

func (c *Container) AddVideo(video Video) {
	c.videos = append(c.videos, video)
	c.mp[video.Id] = &video
}

func (c *Container) GetAll() []Video {
	return c.videos
}

func(c *Container) GetVideo(id string) (*Video,bool) {
	video,ok:= c.mp[id]
	return video,ok
}
