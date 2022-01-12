package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"translateapp/internal/translateapp"
)

type LibreTranslator struct {
	mock.Mock
}

func (l *LibreTranslator) GetLanguages(ctx context.Context) (*[]translateapp.Language, error) {
	args := l.Called(ctx)
	return args.Get(0).(*[]translateapp.Language), args.Error(1)
}

func (l *LibreTranslator) Translate(ctx context.Context, word translateapp.Word) (*translateapp.Word, error) {
	args := l.Called(ctx)
	return args.Get(0).(*translateapp.Word), args.Error(1)

}
