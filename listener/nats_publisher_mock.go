package listener

import "github.com/stretchr/testify/mock"

type publisherMock struct {
	mock.Mock
}

func (p *publisherMock) Publish(subject string, event Event) error {
	args := p.Called(subject, event)
	return args.Error(0)
}
