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
        thsCell.content = "!ERR!"
        thsCell.celltype = ERROR
        return
    }

    var cpyCell *CELL = &Table[row_cpy][col_cpy]
    if cpyCell.celltype == COPY_OP {
        Op_copy(row_cpy, col_cpy)
    } else if cpyCell.celltype == ABS_OP {
        Op_abs(row_cpy, col_cpy)
    }

    thsCell.content = cpyCell.content
    thsCell.celltype = cpyCell.celltype
}

func Op_abs (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    var argument string = thsCell.content[5:len(thsCell.content) - 1]

    row_cpy, col_cpy := op_getcoords_cell(argument)
    var cpyCell *CELL = &Table[row_cpy][col_cpy]

    /**
     * Abs operation is only to cells that already has some value.
     * It means abs function won't call another function to fill the value
     * in the current cell.
     * **/
    if cpyCell.celltype == NUMBER {
        if cpyCell.content[0] == '-' {
            thsCell.content = cpyCell.content[1:]
        } else {
            thsCell.content = cpyCell.content
        }
        thsCell.celltype = cpyCell.celltype
    } else {
        thsCell.content = "!ERR!"
        thsCell.celltype = ERROR
    }
}
