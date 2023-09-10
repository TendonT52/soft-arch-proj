package mock

import "time"

type MockTimeProvider struct{}

func NewMockTimeProvider() *MockTimeProvider {
	return &MockTimeProvider{}
}

func (r *MockTimeProvider) Now() time.Time {
	return time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
}
