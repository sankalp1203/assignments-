package services

import (
	"reflect"
	"testing"
	"viewcount/platform"
)

func TestGetVideos(t *testing.T) {
	tests := []struct {
		name string
		want []platform.Video
	}{
		{
			name: "test1",
			want: []platform.Video{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVideos(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVideos() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPostVideo(t *testing.T) {
	type args struct {
		r *platform.Video
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{r: &platform.Video{Title: "title1"}},
		},
		{
			name: "test2",
			args: args{r: &platform.Video{Title: "title2"}},
		},
	}
	size := 1
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = PostVideo(tt.args.r)
			videos := GetVideos()
			if len(videos) != size {
				t.Errorf("size of container should be %v but it is %v", size, len(videos))
			}
			size++
		})
	}
}

func TestGetViews(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name:  "test1",
			args:  args{id: PostVideo(&platform.Video{Title: "title1"})},
			want:  0,
			want1: true,
		},
		{
			name:  "test2",
			args:  args{id: "123"},
			want:  0,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetViews(tt.args.id)
			if got != tt.want {
				t.Errorf("GetViews() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetViews() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestIncViews(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{id: PostVideo(&platform.Video{Title: "title1"})},
			want: true,
		},
		{
			name: "test2",
			args: args{id: "123"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IncViews(tt.args.id); got != tt.want {
				if got {
					views, _ := GetViews(tt.args.id)
					if views != 1 {
						t.Errorf("view not incremented")
					}
				}
				t.Errorf("IncViews() = %v, want %v", got, tt.want)
			}
		})
	}
}
