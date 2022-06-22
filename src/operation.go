package main

import (
    _ "fmt"
    "strconv"
)

/**
 * When some copy operation will be done, the first cell to
 * be called must be saved because if there is a loop
 * its position in the table let know actually if there is a loop.
 * **/
var no_looprow int = -1
var no_loopcol int = -1

func op_getcoords_cell (strcont string) (int, int) {
    /**
     * Given some =A0 this function returns the coordinates
     * of that cell.
     * =C31
     * ROW: 31
     * COL: 2
     *
     * 65 : Ascii position of 'A', A is the first column
     * in any table.
     * **/
     var col int = int(strcont[1]) - 65
     row, _ := strconv.Atoi(strcont[2:])
     return row, col
}

func op_getsingle_cell (coord string) *CELL {
    /**
     * When the argument of some function is the coordinate
     * of some cell this funcion will be called to get
     * that specific cell
     * **/
    row_arg, col_arg := op_getcoords_cell(coord)
    var cellagr *CELL = &Table[row_arg][col_arg]
    return cellagr
}

func op_getargument (cellcont string) string {
    return cellcont[5:len(cellcont) - 1]
}

func Op_setnoloops (row int, col int) {
    no_looprow = row
    no_loopcol = col
}

func Op_copy (row int, col int) {
    if no_looprow == -1 {
        Op_setnoloops(row, col)
    }
    var thsCell *CELL = &Table[row][col]

    row_cpy, col_cpy := op_getcoords_cell(thsCell.content)
    if row_cpy == no_looprow && col_cpy == no_loopcol {
        thsCell.content = "!N/A!"
        thsCell.celltype = ERROR
        return
    }

    var cpyCell *CELL = &Table[row_cpy][col_cpy]
    if cpyCell.celltype == COPY_OP {
        Op_copy(row_cpy, col_cpy)
    } else if cpyCell.celltype == ABS_OP {
        Op_abs(row_cpy, col_cpy)
    } else if cpyCell.celltype == BIN_OP {
        Op_bin(row_cpy, col_cpy)
    }

    thsCell.content = cpyCell.content
    thsCell.celltype = cpyCell.celltype
}

func Op_abs (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    var argument string = op_getargument(thsCell.content)
    var cpyCell *CELL = op_getsingle_cell(argument)

    /**
     * Abs operation is only to cells that already has some value.
     * It means abs function won't call another function to fill the value
     * in the current cell.
     * **/
    if cpyCell.celltype == INTEGER || cpyCell.celltype == FLOAT {
        if cpyCell.content[0] == '-' {
            thsCell.content = cpyCell.content[1:]
        } else {
            thsCell.content = cpyCell.content
        }
        thsCell.celltype = cpyCell.celltype
    } else {
        thsCell.content = "!REF!"
        thsCell.celltype = ERROR
    }
}

func Op_bin (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    var argument string = op_getargument(thsCell.content)

    if argument[0] == '=' {
        var cpyCell *CELL = op_getsingle_cell(argument)
        if cpyCell.celltype != INTEGER {
            thsCell.content = "!NUM!"
            thsCell.celltype = ERROR
            return
        }
        numint, _ := strconv.Atoi(cpyCell.content)
        thsCell.content = strconv.FormatInt( int64(numint), 2 )
    } else {
        numint, _ := strconv.Atoi(argument)
        thsCell.content = strconv.FormatInt( int64(numint), 2 )
    }
    thsCell.celltype = BINARY_NUM
}
