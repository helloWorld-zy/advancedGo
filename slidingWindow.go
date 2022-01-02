package slidingWindows

import "time"

type SlidingWindow struct {
	Size            int64
	WindowPerBucket int64
	CurrentBucket   int
	Bucket          []int
	StartTime       int64
}

func NewSlidingWindow(size int64, windowPerBucket int64) *SlidingWindow {
	return &SlidingWindow{
		Size:            size,
		WindowPerBucket: windowPerBucket,
		CurrentBucket:   0,
		Bucket:          make([]int, windowPerBucket),
		StartTime:       time.Now().UnixMilli(),
	}
}
func (s *SlidingWindow) WindowsCount() int {
	count := 0
	currentTime := time.Now().UnixMilli()
	time := currentTime - s.Size - s.StartTime
	if time < 0 {
		time = 0
	}
	windowsNeeded := time / s.Size / s.WindowPerBucket
	s.SlideWindow(windowsNeeded)
	for i := 0; i < int(s.WindowPerBucket); i++ {
		count += s.Bucket[i]
	}
	return count
}

func (s *SlidingWindow) SlideWindow(windowsNeeded int64) {
	slideNum := s.WindowPerBucket
	if windowsNeeded == 0 {
		return
	}
	if windowsNeeded < s.WindowPerBucket {
		slideNum = windowsNeeded
	}
	for i := 0; i < int(slideNum); i++ {
		s.CurrentBucket = (s.CurrentBucket + 1) % int(s.WindowPerBucket)
		s.Bucket[s.CurrentBucket] = 0
	}
	s.StartTime = s.StartTime + windowsNeeded*(s.Size/s.WindowPerBucket)
}
