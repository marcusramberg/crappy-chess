package main

import (
	"fmt"
	"image/color"
)

/*
 *
 */
func selectTile() {
	if flag, x, y := click(); flag == true {
		for a, row := range board {
			for b, _ := range row {
				tile := &board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) && tile.piece != (Piece{}) && tile.piece.color == playersTurn {
					if selectedTile != nil {
						selectedTile.color = previousColor
					}
					previousColor = tile.color
					tile.color = color.RGBA{250, 100, 50, 200}
					selectedTile = tile
				}
			}
		}
	}
}

func moveTile() {
	if flag, x, y := click(); flag == true {
		for a, row := range board {
			for b := range row {
				tile := &board[a][b]
				if (tile.xleft <= x && x < tile.xright) && (tile.ytop <= y && y < tile.ybot) {
					if selectedTile == tile {
						selectedTile.color = previousColor
						selectedTile = nil
					} else {
						if canMove(selectedTile, tile) {
							tile.piece = selectedTile.piece
							selectedTile.piece = Piece{}
							selectedTile.color = previousColor
							selectedTile = nil
							if isCheck() {
								fmt.Println("check")
							}
							changeTurn()
						} else {

						}
					}
				}
			}
		}
	}

}

func changeTurn() {
	if playersTurn == 0 {
		playersTurn = 1
	} else {
		playersTurn = 0
	}
}

func isCheck() bool {
	var x, y int
	for a, row := range board {
		for b, tile := range row {
			if tile.piece.rank == 5 && tile.piece.color != playersTurn {
				x, y = b, a
				break
			}
		}
	}
	for _, row := range board {
		for _, tile := range row {
			if tile.piece != (Piece{}) && tile.piece.color == playersTurn {
				if canMove(&tile, &board[y][x]) {
					return true
				}
			}
		}
	}
	return false
}

//Checks if the selected piece is allowed to move to the new tile that you click on.
//returns true/false
func canMove(tileFrom *Tile, tileTo *Tile) bool {
	var boo bool
	switch tileFrom.piece.rank {
	case 0:
		boo = movePawn(tileFrom, tileTo)
	case 1:
		boo = moveKnight(tileFrom, tileTo)
	case 2:
		boo = moveBishop(tileFrom, tileTo)
	case 3:
		boo = moveRook(tileFrom, tileTo)
	case 4:
		boo = moveQueen(tileFrom, tileTo)
	case 5:
		boo = moveKing(tileFrom, tileTo)
	}
	return boo
}

func movePawn(tileFrom *Tile, tileTo *Tile) bool {
	xold, yold := tileFrom.b, tileFrom.a
	xnew, ynew := tileTo.b, tileTo.a
	c := tileFrom.piece.color
	switch c {
	case 0:
		if xnew == xold && ynew == yold-1 && tileTo.piece == (Piece{}) {
			return true
		} else if xnew == xold && ynew == yold-2 && tileTo.piece == (Piece{}) && board[ynew+1][xold].piece == (Piece{}) && yold == 6 {
			return true
		} else if xnew == xold+1 && ynew == yold-1 && tileTo.piece != (Piece{}) && tileTo.piece.color != playersTurn {
			return true
		} else if xnew == xold-1 && ynew == yold-1 && tileTo.piece != (Piece{}) && tileTo.piece.color != playersTurn {
			return true
		}
	case 1:
		if xnew == xold && ynew == yold+1 && tileTo.piece == (Piece{}) {
			return true
		} else if xnew == xold && ynew == yold+2 && tileTo.piece == (Piece{}) && board[ynew-1][xold].piece == (Piece{}) && yold == 1 {
			return true
		} else if xnew == xold+1 && ynew == yold+1 && tileTo.piece != (Piece{}) && tileTo.piece.color != playersTurn {
			return true
		} else if xnew == xold-1 && ynew == yold+1 && tileTo.piece != (Piece{}) && tileTo.piece.color != playersTurn {
			return true
		}

	}
	return false
}

func moveKnight(tileFrom *Tile, tileTo *Tile) bool {

	xold, yold := tileFrom.b, tileFrom.a
	xnew, ynew := tileTo.b, tileTo.a

	if (xnew == xold-1 || xnew == xold+1) && (ynew == yold+2 || ynew == yold-2) && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		return true
	} else if (xnew == xold+2 || xnew == xold-2) && (ynew == yold-1 || ynew == yold+1) && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		return true
	}
	return false
}

func moveBishop(tileFrom *Tile, tileTo *Tile) bool {
	xold, yold := tileFrom.b, tileFrom.a
	xnew, ynew := tileTo.b, tileTo.a
	if tileTo.piece.color == playersTurn && tileTo.piece != (Piece{}) {
		return false
	}
	for i := 1; i < 8; i++ {
		switch {
		case xnew == xold+i && ynew == yold+i:
			for j := 1; j < i; j++ {
				if board[yold+j][xold+j].piece != (Piece{}) {
					return false
				}
			}
			return true
		case xnew == xold-i && ynew == yold-i:
			for j := 1; j < i; j++ {
				if board[yold-j][xold-j].piece != (Piece{}) {
					return false
				}
			}
			return true
		case xnew == xold+i && ynew == yold-i:
			for j := 1; j < i; j++ {
				if board[yold-j][xold+j].piece != (Piece{}) {
					return false
				}
			}
			return true
		case xnew == xold-i && ynew == yold+i:
			for j := 1; j < i; j++ {
				if board[yold+j][xold-j].piece != (Piece{}) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func moveRook(tileFrom *Tile, tileTo *Tile) bool {
	xold, yold := tileFrom.b, tileFrom.a
	xnew, ynew := tileTo.b, tileTo.a

	if xold == xnew && ynew > yold && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		for i := yold + 1; i < ynew; i++ {
			if board[i][xnew].piece != (Piece{}) {
				return false
			}
		}
		return true
	} else if xold == xnew && ynew < yold && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		for i := yold - 1; i > ynew; i-- {
			if board[i][xnew].piece != (Piece{}) {
				return false
			}
		}
		return true
	} else if yold == ynew && xnew > xold && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		for i := xold + 1; i < xnew; i++ {
			if board[ynew][i].piece != (Piece{}) {
				return false
			}
		}
		return true
	} else if yold == ynew && xnew < xold && (tileTo.piece == (Piece{}) || tileTo.piece.color != playersTurn) {
		for i := xold - 1; i > xnew; i-- {
			if board[ynew][i].piece != (Piece{}) {
				return false
			}
		}
		return true
	}

	return false
}

func moveQueen(tileFrom *Tile, tileTo *Tile) bool {
	if moveRook(tileFrom, tileTo) || moveBishop(tileFrom, tileTo) {
		return true
	}
	return false
}

func moveKing(tileFrom *Tile, tileTo *Tile) bool {
	xold, yold := tileFrom.b, tileFrom.a
	xnew, ynew := tileTo.b, tileTo.a
	if tileTo.piece.color == playersTurn && tileTo.piece != (Piece{}) {
		return false
	}
	
	return false
}
