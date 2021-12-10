# MEXCEL - Mini excel

Is a program that can read one file with a "table" inside, each cell has one number or one formula or one operation to make
the main idea is try to make something like Excel but more simple :D

## Clone operation (v, >, ^, <)

This operation will copy the value of the cell that point to the character, if we have something like:
```
| 1 | 2 | 3 | 5 |
| 2 | 0 | < | 2 |
| 3 | 2 | > | 5 |
| 4 | v | 3 | 5 |
| 5 | 2 | ^ | 5 |
```
The output will be:
```
| 1 | 2 | 3 | 5 |
| 2 | 0 | 0 | 2 |
| 3 | 2 | 5 | 5 |
| 4 | 2 | 3 | 5 |
| 5 | 2 | ^ | 5 |
```

there are more operations! just give a few seconds...!
[idea from](https://github.com/tsoding/minicel)
