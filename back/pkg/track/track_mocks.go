package track

type mockRepo struct {
}

func NewMockRepo() Repository {
	return &mockRepo{}
}

func (mockRepo *mockRepo) GetTracks(term string) (*DataApi, error) {
	return nil, nil
}
