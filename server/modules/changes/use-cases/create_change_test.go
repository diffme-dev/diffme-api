package use_cases

import (
	//"context"
	//"diffme.dev/diffme-api/tests/mocks"
	//"errors"
	//"github.com/stretchr/testify/assert"
	//"github.com/stretchr/testify/mock"
	"testing"
	//"time"
)

func TestFetch(t *testing.T) {
	//mockArticleRepo := new(mocks.ArticleRepository)
	//mockArticle := domain.Article{
	//	Title:   "Hello",
	//	Content: "Content",
	//}

	//mockListArtilce := make([]domain.Article, 0)
	//mockListArtilce = append(mockListArtilce, mockArticle)

	t.Run("success", func(t *testing.T) {
		//mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
		//	mock.AnythingOfType("int64")).Return(mockListArtilce, "next-cursor", nil).Once()
		//mockAuthor := domain.Author{
		//	ID:   1,
		//	Name: "Iman Tumorang",
		//}
		//mockAuthorrepo := new(mocks.AuthorRepository)
		//mockAuthorrepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockAuthor, nil)
		//u := ucase.NewArticleUsecase(mockArticleRepo, mockAuthorrepo, time.Second*2)
		//num := int64(1)
		//cursor := "12"
		//list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		//cursorExpected := "next-cursor"
		//assert.Equal(t, cursorExpected, nextCursor)
		//assert.NotEmpty(t, nextCursor)
		//assert.NoError(t, err)
		//assert.Len(t, list, len(mockListArtilce))
		//
		//mockArticleRepo.AssertExpectations(t)
		//mockAuthorrepo.AssertExpectations(t)
	})

}
