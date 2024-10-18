package services

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/bxcodec/faker/v3"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

type ListTestCases struct {
	Name         string
	Response     []domain.Newsletter
	CountResp    int64
	CountRespMsg string
	ErrResp      *domain.ApiError
	ErrRespMsg   string
}

func TestListNewsletters(t *testing.T) {
	assert := assert.New(t)

	// Dummies
	ctx := context.Background()
	logger := zerolog.New(os.Stderr)

	// Repo
	repo := new(repositories.MockINewsletterRepository)
	opts := ports.FindManyOpts{
		Skip: 0,
		Take: 100,
	}

	// Service
	evt := new(ports.MockEventMessaging)
	service := NewNewsletterService(ctx, &logger, repo, evt)

	dummyRes := domain.Newsletter{}
	_ = faker.FakeData(&dummyRes)
	testCases := []ListTestCases{
		{
			Name:         "1 newsletter",
			Response:     []domain.Newsletter{},
			CountResp:    int64(0),
			CountRespMsg: "Should return 0 newsletter",
			ErrRespMsg:   "Should not return any error",
		},
		{
			Name:         "Empty newsletters",
			Response:     []domain.Newsletter{},
			CountResp:    int64(0),
			CountRespMsg: "Should return empty newsletters",
			ErrResp:      domain.ErrFetchNewsletters,
			ErrRespMsg:   "Should return error",
		},
	}

	// CASE 1 SUCCESS:
	repo.On("FindMany", ctx, opts, &testCases[0].Response, true).Return(&testCases[0].CountResp, nil).Once()
	paging := domain.PaginationsParams{
		Page:     1,
		PageSize: 100,
	}

	result, apiErr := service.ListNewsletters(ctx, &paging)

	assert.Nil(apiErr, testCases[0].ErrRespMsg)
	passed := assert.NotNil(result, "Result not nil")
	if passed {
		assert.Equal(result.Metadata.Length, int64(len(testCases[0].Response)), testCases[0].CountRespMsg)
	}

	// CASE 2 ERROR:
	repo.On("FindMany", ctx, opts, &testCases[1].Response, true).Return(&testCases[1].CountResp, testCases[1].ErrResp).Once()

	result2, apiErr2 := service.ListNewsletters(ctx, &paging)

	assert.Nil(result2, testCases[1].CountRespMsg)
	passed = assert.NotNil(apiErr2, testCases[1].ErrRespMsg)
	if passed {
		assert.Equal(apiErr2.Code, testCases[1].ErrResp.Code)
	}

	repo.AssertExpectations(t)
}

func TestSendNewsletters(t *testing.T) {
	assert := assert.New(t)

	// Dummies
	ctx := context.Background()
	logger := zerolog.New(os.Stderr)

	// Repo
	repo := new(repositories.MockINewsletterRepository)

	// Service
	evt := new(ports.MockEventMessaging)
	service := NewNewsletterService(ctx, &logger, repo, evt)

	newsletter := domain.Newsletter{}
	dummyDto := domain.SendNewsletterDTO{NewsletterId: "TEST"}
	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"_id": "TEST",
		},
	}
	updOpts := ports.UpdateOpts{
		Filter: map[string]interface{}{
			"_id": "TEST",
		},
		Payload: &map[string]interface{}{
			"sentCount": newsletter.SentCount + 1,
		},
	}
	evtOpt := domain.MessageEvent{
		EventTopic: domain.SendEmailTopic,
		Data: map[string]interface{}{
			"newsletter": &newsletter,
		},
	}

	// CASE 1 SUCCESS:
	repo.On("FindOne", ctx, opts, &newsletter).Return(nil).Once()
	repo.On("UpdateOne", ctx, updOpts).Return(nil, nil).Once()
	evt.On("SendMessage", ctx, &evtOpt).Return().Once()
	apiErr := service.SendNewsletter(ctx, &dummyDto, &newsletter)

	assert.Nil(apiErr, "No api error returned")

	// CASE 2 UPDATE ERROR (BYPASS)
	repo.On("FindOne", ctx, opts, &newsletter).Return(nil).Once()
	repo.On("UpdateOne", ctx, updOpts).Return(nil, errors.New("ERROR")).Once()
	evt.On("SendMessage", ctx, &evtOpt).Return().Once()
	apiErr = service.SendNewsletter(ctx, &dummyDto, &newsletter)

	assert.Nil(apiErr, "No api error returned bypassing update error")

	// CASE 3 ERROR:
	repo.On("FindOne", ctx, opts, &newsletter).Return(domain.ErrNewsletterNotFound).Once()
	apiErr = service.SendNewsletter(ctx, &dummyDto, &newsletter)

	passed := assert.NotNil(apiErr, "Error expected")
	if passed {
		assert.Equal(apiErr.Code, domain.ErrNewsletterNotFound.Code, "Expected not found error")
	}

	repo.AssertExpectations(t)
	evt.AssertExpectations(t)
}
