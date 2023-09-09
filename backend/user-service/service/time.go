package service

import "time"

type RealTimeProvider struct{}

func NewRealTimeProvider() *RealTimeProvider {
	return &RealTimeProvider{}
}

func (r *RealTimeProvider) Now() time.Time {
	return time.Now()
}
