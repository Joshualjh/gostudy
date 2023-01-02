package db

import (
	"main/pkg/model"

	"github.com/pkg/errors"
)

func (h *DBHandler) CreateBoard(board *model.Board) (*model.Board, error) {
	result := h.gDB.Create(board)
	return board, errors.Wrap(result.Error, "db handler error1")
}

func (h *DBHandler) GetBoradList() ([]*model.Board, error) {
	boardList := []*model.Board{}
	result := h.gDB.Find(&boardList)

	return boardList, errors.Wrap(result.Error, "db handler error2")
}

func (h *DBHandler) GetBoradListById(id string) (*model.Board, error) {
	board := &model.Board{}
	result := h.gDB.First(board, id)

	return board, errors.Wrap(result.Error, "db handler error2")
}

func (h *DBHandler) DeleteBoard(id string) error {
	result := h.gDB.Delete(&model.Board{}, id)
	return errors.Wrap(result.Error, "db handler error3")
}

func (h *DBHandler) UpdateBoard(id uint, newboard *model.Board) (*model.Board, error) {
	oldBoard := &model.Board{}
	result := h.gDB.Model(oldBoard).Where(id).Updates(newboard)
	return oldBoard, errors.Wrap(result.Error, "db handler error4")
}
