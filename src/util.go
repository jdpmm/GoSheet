package main
import (
    "strconv"
    "strings"
)

func Utl_int32 (num string) string {
    numint, _ := strconv.Atoi(num)
    if numint >= -2147483648 && numint <= 2147483647 { 
        return num
    }
    return "!INT!"
}

func Utl_bin32 (num string) string {
    if len(num) == 33 {
        if strings.Count(num, "1") == 31 {
            return "!BIN!"
        }
    }
    return num
}

func Utl_minmax (field []CELL, type_ CELL_TYPE) (int, float32, string) {
    var _int int = 0
    var _float float32 = 0.0

    var n_ints, n_floats int = 0, 0
    var cu_cell CELL

    for idx := 0; idx < len(field); idx++ {
        cu_cell = field[idx]
        if cu_cell.celltype == INTEGER {
            if cu_cell.asint < _int && type_ == MIN_OP { _int = cu_cell.asint }
            if cu_cell.asint > _int && type_ == MAX_OP { _int = cu_cell.asint }
            n_ints++
        } else {
            if cu_cell.asfloat < _float && type_ == MIN_OP { _float = cu_cell.asfloat }
            if cu_cell.asfloat > _float && type_ == MIN_OP { _float = cu_cell.asfloat }
            n_floats++
        }
    }

    if n_ints > n_floats {
        return _int, _float, "int"
    }
    return _int, _float, "float"
}
