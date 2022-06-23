package main
import (
    _ "fmt"
    "strconv"
    "strings"
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

     // XXX: Could be better
     if col > max_col { col = 0 }
     if row > max_row { row = 0 }
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

func op_getdouble_args (arg string, lpar_idx int) (string, string) {
    /**
     * When there is a function that takes two arguments as AND
     * this function will be called and the return will be those arguments.
     * **/
    var twocoords string = arg[lpar_idx + 1:len(arg) - 1]
    arguments := strings.Split(twocoords, ";")
    return arguments[0], arguments[1]
}

func op_setvalues_2args (cCell *CELL, arg1 string, arg2 string) (int, int) {
    /**
     * This function will be called when some function takes
     * two arguments and those arguments must be numbers such as
     * AND function.
     * **/
     var iscell CELL
     var rowcll, colcll int
     var r_arg1, r_arg2 int

     if arg1[0] == '=' {
         rowcll, colcll = op_getcoords_cell(arg1)
         iscell = Table[rowcll][colcll]
         if iscell.celltype != INTEGER {
             cCell.celltype = ERROR
             return -1, -1
         }
         r_arg1, _ = strconv.Atoi(iscell.content)
     } else {
         r_arg1, _ = strconv.Atoi(arg1)
     }

     if arg2[0] == '=' {
         rowcll, colcll = op_getcoords_cell(arg2)
         iscell = Table[rowcll][colcll]
         if iscell.celltype != INTEGER {
             cCell.celltype = ERROR
             return -1, -1
         }
         r_arg2, _ = strconv.Atoi(iscell.content)
     } else {
         r_arg2, _ = strconv.Atoi(arg2)
     }
     return r_arg1, r_arg2
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
    } else if cpyCell.celltype == AND_OP {
        Op_and(row_cpy, col_cpy)
    } else if cpyCell.celltype == OR_OP {
        Op_or(row_cpy, col_cpy)
    } else if cpyCell.celltype == XOR_OP {
        Op_xor(row_cpy, col_cpy)
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

    var binvalue string
    if argument[0] == '=' {
        var cpyCell *CELL = op_getsingle_cell(argument)
        if cpyCell.celltype != INTEGER {
            thsCell.content = "!NUM!"
            thsCell.celltype = ERROR
            return
        }
        numint, _ := strconv.Atoi(cpyCell.content)
        binvalue = strconv.FormatInt( int64(numint), 2 )
    } else {
        numint, _ := strconv.Atoi(argument)
        binvalue = strconv.FormatInt( int64(numint), 2 )
    }

    if binvalue[0] == '-' {
        binvalue = "-0b" + binvalue[1:]
    } else {
        binvalue = "0b" + binvalue
    }

    thsCell.content = binvalue
    thsCell.celltype = BINARY_NUM
}

func Op_and (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    argstr1, argstr2 := op_getdouble_args(thsCell.content, 4)
    arg1, arg2 := op_setvalues_2args(thsCell, argstr1, argstr2)

    if thsCell.celltype == ERROR {
        thsCell.content = "!REF!"
        return
    }
    thsCell.content = strconv.Itoa(int(arg1 & arg2))
    thsCell.celltype = INTEGER
}

func Op_or (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    argstr1, argstr2 := op_getdouble_args(thsCell.content, 3)
    arg1, arg2 := op_setvalues_2args(thsCell, argstr1, argstr2)

    if thsCell.celltype == ERROR {
        thsCell.content = "!REF!"
        return
    }
    thsCell.content = strconv.Itoa(int(arg1 | arg2))
    thsCell.celltype = INTEGER
}

func Op_xor (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    argstr1, argstr2 := op_getdouble_args(thsCell.content, 4)
    arg1, arg2 := op_setvalues_2args(thsCell, argstr1, argstr2)

    if thsCell.celltype == ERROR {
        thsCell.content = "!REF!"
        return
    }
    thsCell.content = strconv.Itoa(int(arg1 ^ arg2))
    thsCell.celltype = INTEGER
}
