package service

import "learningGin/entity"

// VideoService interface
type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

// videoServiceImpl concrete Impl
type videoServiceImpl struct {
	videos []entity.Video
}

/* This method is used to crete a new video service impl
The return type is VideoService which is our Interface
In the body of the method we return a pointer of our video service implmentation
This tells the compiler that video service impl implements the video service
If video service does not implement the interface then we get a compilation error
*/
func New() VideoService {
	return &videoServiceImpl{
		videos: []entity.Video{},
	}
}

// Save We use the receiver functions below to implement the video service interface methods
func (v *videoServiceImpl) Save(video entity.Video) entity.Video {
	v.videos = append(v.videos, video)
	return video
}

func (v *videoServiceImpl) FindAll() []entity.Video {
	return v.videos
}
