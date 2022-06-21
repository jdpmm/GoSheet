package main
import (
    "strings"
    "fmt"
    "regexp"
)

type CELL_TYPE int
const (
    NUMBER = iota
    STRING
    BOOL
    UNKNOWN
    COPY_OP
)

type CELL struct {
    content  string
    row      int
    col      int
    celltype CELL_TYPE
}

/**
 * max_row and max_col works to know how long the table is, and know which cells
 * are used.
 * max_dig works to print the new table with format.
 * **/
var Table[500][26] CELL
var max_row int = 0
var max_col int = 0
var max_dig int = 5

func tbl_aux_stripwhitespaces (str string) string {
    var newstr string
    var instr bool = false
    for idx := 0; idx < len(str); idx++ {
        if str[idx] == '"' {
            instr = !instr
        }
        if str[idx] != ' ' && !instr {
            newstr += string(str[idx])
        }
        if instr {
            newstr += string(str[idx])
        }
    }

    if len(newstr) > max_dig {
        max_dig = len(newstr)
    }
    return newstr
}

func tbl_setcell (content string, row int, col int) {
    isnumber, _ := regexp.Compile("^(\\d+|\\d+\\.\\d+)$")
    isbool,   _ := regexp.Compile("^(TRUE|FALSE)$")
    isstring, _ := regexp.Compile("^\".*\"$")
    iscopyop, _ := regexp.Compile("^=[A-Z]{1}[0-9]{1,3}$");

    var newC CELL
    newC.row = row;
    newC.col = col;
    newC.content = content;
    if isnumber.MatchString(content) {
        newC.celltype = NUMBER
    } else if isbool.MatchString(content) {
        newC.celltype = BOOL
    } else if isstring.MatchString(content) {
        newC.celltype = STRING
    } else if iscopyop.MatchString(content) {
        newC.celltype = COPY_OP
    } else {
        newC.celltype = UNKNOWN
        newC.content = "!ERR!"
    }
    Table[row][col] = newC
}

func tbl_print (content string) {
    fmt.Printf(" %s ", content)
    for spc := 0; spc < (max_dig - len(content)); spc++ {
        fmt.Printf(" ")
    }
}

func Tbl_getrow (rowstr string, rowint int) {
    if len(rowstr) == 0 {
        fmt.Printf("(%d line): Empty.\n", rowint)
        return
    }

    cellsVec := strings.Split(rowstr, "|")
    var colint int
    for colint = 0; colint < len(cellsVec) - 1; colint++ {
        cellsVec[colint] = tbl_aux_stripwhitespaces(cellsVec[colint])
        tbl_setcell (cellsVec[colint], rowint, colint);
    }

    if colint > max_col {
        max_col = colint
    }
    max_row++
}

func Tbl_maketable () {
    var cCell *CELL
    for c_row := 0; c_row < max_row; c_row++ {
        for c_col := 0; c_col < max_col; c_col++ {
            cCell = &Table[c_row][c_col]
            Op_setnoloops(-1, -1)

            if cCell.celltype == COPY_OP {
                Op_copy(c_row, c_col)
            }

            tbl_print(cCell.content)
        }
        fmt.Printf("\n");
    }
}
