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
| 5 | 2 | 3 | 5 |
```

### No loops on clone operation

Can not there be loops on this operation since each cell starts with a value of 0, so if you made something like:
```
| 1 | 2 |
| > | < |
| 2 | 1 |
```

The output will be:
```
| 1 | 2 |
| 0 | 0 |
| 2 | 1 |
```

because ```(1, 0)``` will copy the value ```(1, 1)``` but the value of ```(1, 1)``` is 0 since this cell is not a number, so
```(1, 0) == (1, 1)``` 

**The parser read from the top to the bottom and from left to the right**

```
| 1 | > | v |    | 1 | 0 | 0 |
| > | 9 | v | == | 9 | 9 | 0 |
| ^ | 9 | v | == | 9 | 9 | 0 |
| ^ | < | < |    | 9 | 9 | 9 |
```

## Copy operation (:A1)

This operation will copy the value of the cell indicated, to indicate one cell you must start with a ":", then the letter of the column and the number of the row, for example:

```
| 1 | 2 |
| 3 | 4 |
| 5 | 6 |

A1 = 1 :: B1 = 2
A2 = 3 :: B2 = 4
A3 = 5 :: B3 = 6
```

So to copy one value:

```
| 1  | 2   |
| 3  | 4   |
| 5  | :B3 |
```

The output will be:

```
| 1  | 2 |
| 3  | 4 |
| 5  | 6 |
```

there are more operations! just give a few seconds...!

[idea from](https://github.com/tsoding/minicel)
