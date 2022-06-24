package main
import (
    "fmt"
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

func op_getargument (cellcont string) string {
    return cellcont[5:len(cellcont) - 1]
}

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

     // XXX: could be improve it
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
    return &Table[row_arg][col_arg]
}

func op_getdouble_args (arg string, lpar_idx int) (string, string) {
    /**
     * When there is a operation that takes two arguments as AND
     * this function will be called and the return value
     * will be those arguments (not the values, just the representations).
     * =AND(4;=A0;)
     * returns:
     *  ~ 4
     *  ~ =A0
     *
     * =AND(=A4;4;)
     *      \
     *      lpar_idx points here, but is added 1 to get all argument
     *      without the parenthesis.
     * **/
    var twocoords string = arg[lpar_idx + 1:len(arg) - 1]
    arguments := strings.Split(twocoords, ";")
    return arguments[0], arguments[1]
}

func op_setvalues_2args (cCell *CELL, arg1 string, arg2 string) (int, int) {
    /**
     * This function will be called when some operation takes
     * two arguments and those arguments must be numbers such as
     * AND operation.
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
         r_arg1 = iscell.asint
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
         r_arg2 = iscell.asint
     } else {
         r_arg2, _ = strconv.Atoi(arg2)
     }
     return r_arg1, r_arg2
}

func op_getcells_field (args string, lpar_idx int) []CELL {
    /**
     * This function will be called when the operation takes two
     * arguments and those arguments must be another cells coordinates
     * such as MIN.
     * **/
    cellstr1, cellstr2 := op_getdouble_args(args, lpar_idx)

    /**
     * There are two posibilites (May be two):
     *  ~ Wants one row. (A0, D0) (The row is the same).
     *  ~ Wants one col. (A0, A3) (The col is the same).
     * All elements in the range must be numbers! if some element
     * is not, will not be added.
     * **/
     var field []CELL
     row_C1, col_C1 := op_getcoords_cell(cellstr1)
     row_C2, col_C2 := op_getcoords_cell(cellstr2)

     if cellstr1[1] != cellstr2[1] {
         var ccell CELL
         for col := col_C1; col <= col_C2; col++ {
             ccell = Table[row_C1][col]
             if ccell.celltype == INTEGER || ccell.celltype == FLOAT {
                field = append(field, ccell)
             }
         }
     } else {
         var ccell CELL
         for row := row_C1; row <= row_C2; row++ {
             ccell = Table[row][col_C1]
             if ccell.celltype == INTEGER || ccell.celltype == FLOAT {
                field = append(field, ccell)
             }
         }
     }
     return field
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
    var cpType CELL_TYPE = cpyCell.celltype

    if cpType == COPY_OP {
        Op_copy(row_cpy, col_cpy)
    } else if cpType == ABS_OP {
        Op_abs(row_cpy, col_cpy)
    } else if cpType == AND_OP || cpType == OR_OP || cpType == XOR_OP {
        Op_bitwise(row_cpy, col_cpy, cpType)
    } 


    if cpyCell.celltype == INTEGER { thsCell.asint = cpyCell.asint }
    if cpyCell.celltype == FLOAT { thsCell.asfloat = cpyCell.asfloat }
    thsCell.content = cpyCell.content
    thsCell.celltype = cpyCell.celltype
}

func Op_abs (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    var argument string = op_getargument(thsCell.content)
    var cpyCell *CELL = op_getsingle_cell(argument)

    /**
     * Abs operation is only to cells that already has some value.
     * It means abs function won't call another operation to fill the value
     * in the current cell.
     * **/
    if cpyCell.celltype == INTEGER || cpyCell.celltype == FLOAT {
        if cpyCell.content[0] == '-' {
            thsCell.content = cpyCell.content[1:]
        } else {
            thsCell.content = cpyCell.content
        }

        if cpyCell.celltype == INTEGER { thsCell.asint = cpyCell.asint }
        if cpyCell.celltype == FLOAT { thsCell.asfloat = cpyCell.asfloat }
        thsCell.celltype = cpyCell.celltype
    } else {
        thsCell.content = "!REF!"
        thsCell.celltype = ERROR
    }
}

func Op_bitwise (row int, col int, op CELL_TYPE) {
    var thsCell *CELL = &Table[row][col]
    var argstr1, argstr2 string
    if op == AND_OP || op == XOR_OP {
        argstr1, argstr2 = op_getdouble_args(thsCell.content, 4)
    } else {
        argstr1, argstr2 = op_getdouble_args(thsCell.content, 3)
    }

    arg1, arg2 := op_setvalues_2args(thsCell, argstr1, argstr2)
    if thsCell.celltype == ERROR {
        thsCell.content = "!REF!"
        return
    }

    var bitwise_op int
    if op == AND_OP { bitwise_op = (arg1 & arg2) }
    if op == OR_OP  { bitwise_op = (arg1 | arg2) }
    if op == XOR_OP { bitwise_op = (arg1 ^ arg2) }

    thsCell.content = strconv.Itoa(bitwise_op)
    thsCell.celltype = INTEGER
    thsCell.asint = bitwise_op
}

func Op_min (row int, col int) {
    var thsCell *CELL = &Table[row][col]
    var cells []CELL = op_getcells_field(thsCell.content, 4)

    asint, asfloat, should_be := Utl_min(cells)
    if should_be == "int" {
        thsCell.content = strconv.Itoa(asint)
        thsCell.asint = asint
        thsCell.celltype = INTEGER
        return
    }

    thsCell.content = fmt.Sprintf("%f", asfloat)
    thsCell.asfloat = asfloat
    thsCell.celltype = FLOAT
}
