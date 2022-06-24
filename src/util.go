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

func Utl_min (field []CELL) (int, float32, string) {
    var minint int = 0
    var minfloat float32 = 0.0

    var n_ints, n_floats int = 0, 0
    var cu_cell CELL

    for idx := 0; idx < len(field); idx++ {
        cu_cell = field[idx]
        if cu_cell.celltype == INTEGER {
            if cu_cell.asint < minint {
                minint = cu_cell.asint
            }
            n_ints++
        } else {
            if cu_cell.asfloat < minfloat {
                minfloat = cu_cell.asfloat
            }
            n_floats++
        }
    }

    if n_ints > n_floats {
        return minint, minfloat, "int"
    }
    return minint, minfloat, "float"
}
