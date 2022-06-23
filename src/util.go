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
