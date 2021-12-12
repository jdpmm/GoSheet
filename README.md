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

### Loops

A loop error will be generated when one cell point to another cell and this cell point to the first cell, looks like:

```
| 1 | 2 |
| > | < |
| 2 | 1 |
```

The output will be:
```
[anothers cells values]
Loop detected
coord: (1, 2)
```

**The parser always read from the top to the bottom and from left to the right**

```
| 1 | > | v |    | 1 | 9 | 9 |
| > | 9 | v | == | 9 | 9 | 9 |
| ^ | 9 | v | == | 9 | 9 | 9 |
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
| 5  | :A1 |
```

The output will be:

```
| 1  | 2 |
| 3  | 4 |
| 5  | 1 |
```

### Loops

A loop error will be generated when one cell try to copy the value of another and this cell try to copy the value of the first cell!!
<br/>
Looks like:
```
| :B3  | 2   |
| 3    | 4   |
| 5    | :A1 |
```

The output will be:

```
[anothers cells values]
Loop detected
coord: (3, 2)
```

## Arithmetic (:3+1)

To make one arithmetic operation you need at leaste one operator (+, *, / or -) and two numbers, looks like:
```
| :3+3  | :2-1  |
| 3     | 4     |
| :4/2  | :A1*2 |
```

The output will be:
```
| 6  | 1  |
| 3  | 4  |
| 2  | 12 |
```

## TODO
- [x] Arithmetic and negative numbers
- [x] =sum(from : to)
- [x] =max(from : to)
- [x] =min(from : to)
- [x] =med(from : to)
- [x] =mda(from : to)
- [ ] =moa(col1 : to)
- [ ] print with format

---

## ???

- [moa](https://es.wikipedia.org/wiki/Moda_(estad%C3%ADstica))
- [med](https://es.wikipedia.org/wiki/Media_(matem%C3%A1ticas))
- [mda](https://es.wikipedia.org/wiki/Mediana_(estad%C3%ADstica)#:~:text=En%20el%20%C3%A1mbito%20de%20la,un%20conjunto%20de%20datos%20ordenados.)

[idea from](https://github.com/tsoding/minicel)
