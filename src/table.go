package main
import (
    "strings"
    "fmt"
    "regexp"
    "strconv"
)

type CELL_TYPE int
const (
    INTEGER = iota
    FLOAT
    STRING
    BOOL
    BINARY_NUM

    COPY_OP
    ABS_OP
    AND_OP
    OR_OP
    XOR_OP
    MIN_OP
    MAX_OP
    ARITH_OP

    ERROR
)

type CELL struct {
    content  string
    row      int
    col      int
    celltype CELL_TYPE

    asint    int
    asfloat  float32
}

/**
 * regex used not only in this golang file
 * **/
var isinteger, _ = regexp.Compile("^(-|)\\d+$")
var isfloat,   _ = regexp.Compile("^(-|)\\d+.\\d+$")
var iscopyop,  _ = regexp.Compile("^=[A-Z]{1}[0-9]{1,3}$")

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
    return newstr
}

func tbl_print (content string) {
    fmt.Printf("%s ", content)
    for spc := 0; spc < (max_dig - len(content)); spc++ {
        fmt.Printf(" ")
    }
}

func tbl_setcell (content string, row int, col int) {
    isbool,    _ := regexp.Compile("^(TRUE|FALSE)$")
    isstring,  _ := regexp.Compile("^\".*\"$")

    isabsop,   _ := regexp.Compile("^=ABS\\(=[A-Z]{1}[0-9]{1,3}\\)$")
    isminop,   _ := regexp.Compile("^=MIN\\((=[A-Z]{1}[0-9]{1,3};){2}\\)$")
    ismaxop,   _ := regexp.Compile("^=MAX\\((=[A-Z]{1}[0-9]{1,3};){2}\\)$")
    isarithop, _ := regexp.Compile("^=ARITH\\((.+,){2,}\\)$")
    
    isandop,   _ := regexp.Compile("^=AND\\((=[A-Z]{1}[0-9]{1,3};|(-|)[0-9]+;){2}\\)$");
    isorop,    _ := regexp.Compile("^=OR\\((=[A-Z]{1}[0-9]{1,3};|(-|)[0-9]+;){2}\\)$");
    isxorop,   _ := regexp.Compile("^=XOR\\((=[A-Z]{1}[0-9]{1,3};|(-|)[0-9]+;){2}\\)$");

    var newC CELL
    newC.row = row;
    newC.col = col;
    newC.content = content;

    // XXX: Could be better
    if isinteger.MatchString(content) {
        newC.content = Utl_int32(content)
        if newC.content == "!INT!" {
            newC.celltype = ERROR
        } else {
            newC.celltype = INTEGER
            newC.asint, _ = strconv.Atoi(content)
        }
    } else if isbool.MatchString(content) {
        newC.celltype = BOOL
    } else if isstring.MatchString(content) {
        newC.celltype = STRING
        newC.content = content[1:len(content) - 1] // removes the quotes
    } else if isfloat.MatchString(content) {
        newC.celltype = FLOAT
        asfloat, _ := strconv.ParseFloat(content, 32)
        newC.asfloat = float32(asfloat)
    } else if iscopyop.MatchString(content) {
        newC.celltype = COPY_OP
    } else if isabsop.MatchString(content) {
        newC.celltype = ABS_OP
    } else if isandop.MatchString(content) {
        newC.celltype = AND_OP
    } else if isorop.MatchString(content) {
        newC.celltype = OR_OP
    } else if isxorop.MatchString(content) {
        newC.celltype = XOR_OP
    } else if isminop.MatchString(content) {
        newC.celltype = MIN_OP
    } else if ismaxop.MatchString(content) {
        newC.celltype = MAX_OP
    } else if isarithop.MatchString(content) {
        newC.celltype = ARITH_OP
    } else {
        newC.celltype = ERROR
        newC.content = "!UNK!"
    }
    Table[row][col] = newC
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
            if cCell.celltype == ABS_OP {
                Op_abs(c_row, c_col)
            }
            if cCell.celltype == AND_OP || cCell.celltype == OR_OP || cCell.celltype == XOR_OP {
                Op_bitwise(c_row, c_col, cCell.celltype)
            }
            if cCell.celltype == MIN_OP || cCell.celltype == MAX_OP {
                Op_minmax(c_row, c_col, cCell.celltype)
            }
            if cCell.celltype == ARITH_OP {
                Op_arith(c_row, c_col)
            }

            if len(cCell.content) > max_dig {
                max_dig = len(cCell.content)
            }
        }
    }
}

func Tbl_printable () {
    for c_row := 0; c_row < max_row; c_row++ {
        for c_col := 0; c_col < max_col; c_col++ {
            tbl_print(Table[c_row][c_col].content)
        }
        fmt.Printf("\n")
    }
}
