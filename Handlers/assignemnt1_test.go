package handlers

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock_querybuilder "github.com/niroopreddym/caymanislandsquerybuilder/mocks"
	"github.com/stretchr/testify/assert"
)

func Test_Assignment1_returnsSuccess(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	conditionalQueryMock := mock_querybuilder.NewMockIQueryBuilder(controller)
	joinQueryMock := mock_querybuilder.NewMockIQueryBuilder(controller)

	conditionalQueryMock.EXPECT().GetQueryPattern(gomock.Any()).AnyTimes().Return("hello")
	handler := AssignemntHandler{
		ConditionalQueryBuilder: conditionalQueryMock,
		JoinBuilder:             joinQueryMock,
		FileName:                "./testdata/input.json",
	}

	str := handler.Assignment1()
	assert.Equal(t, str, "hello")
}
