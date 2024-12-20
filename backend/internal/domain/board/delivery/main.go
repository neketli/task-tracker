package delivery

import (
	"projectly-server/internal/domain/board/usecase"

	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	boardUseCase usecase.BoardUseCase
}

func New(router *echo.Group, b usecase.BoardUseCase) {
	handler := &BoardHandler{boardUseCase: b}

	board := router.Group("/board")
	{
		board.POST("/create", handler.CreateBoard)
		board.PATCH("/:id", handler.UpdateBoard)
		board.DELETE("/:id", handler.DeleteBoard)
		board.GET("/:id", handler.GetBoard)

		board.GET("/list", handler.GetBoardList)
		board.GET("/list-user", handler.GetUserBoardList)
	}
}
