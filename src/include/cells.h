#ifndef MEXCEL_CELLS_H
#define MEXCEL_CELLS_H

#include "utiles.h"
#define A_ASCII_POS 65

typedef enum {
    NUMBER,
    CLONE,
    COPY
} c_Type;

typedef struct CELL {
    int value { 0 };
    c_Type type { NUMBER };
    int coord[2];

    // value as another type, if isn't a number yet
    std::string vaot { "-" };
} *cell;

typedef struct COL {
    std::vector<cell> child_s;
} col;

col columns[26];
int ncols;
int nrows;

void set_cell (int i_row, int i_col, std::string value) {
    cell newcell = new (struct CELL);
    newcell->coord[0] = i_row;
    newcell->coord[1] = i_col;

    if ( value == "v" || value == "^" || value == ">" || value == "<" ) {
        newcell->type = CLONE;
        newcell->vaot = value;
    }
    else if ( value[0] == ':' ) {
        newcell->type = COPY;
        newcell->vaot = value;
    }
    else {
        newcell->type = NUMBER;
        newcell->value = stoi(value);
    }

    // i_row as index because i_row (nrows_defined as paramter) inc his value for each row on the table,
    // so if we will use i_col (i as paramter) the table will be vertial (0, 1) -> (1, 0)
    // because i_col inc his value for each value on the row
    columns[i_row].child_s.push_back(newcell);
}

void clone_value (cell tcell, int idxrow, int idxcol) {
    // only could be >, <, ^, v
    std::string to = tcell->vaot;
    cell toclone;

    try {
        if ( to == "^" && idxrow > 0 ) idxrow--;
        else if ( to == "v" && idxrow < (nrows - 1) ) idxrow++;
        else if ( to == "<" && idxcol > 0 ) idxcol--;
        else if ( to == ">" && idxcol < (ncols - 1) ) idxcol++;

        else {
            throw "Out of range";
        }
    }
    catch (const char* error) {
        printf("\nTrying to clone one value\n");
        printf("%s\n", error);
        printf("-> (%d, %d)\n", idxrow, idxcol);
        exit(1);
    }

    toclone = columns[idxrow].child_s.at(idxcol);
    tcell->value = toclone->value;
    tcell->type = NUMBER;
}

void copy_value (cell tcell, int idxrow, int idxcol) {
    const int i_coltocp = tcell->vaot[1] - A_ASCII_POS;
    const int i_rowtocp = stoi( substr(tcell->vaot, 2, tcell->vaot.size()) ) - 1;

    if ( i_coltocp > ncols || i_rowtocp > nrows ) {
        printf("Tyring to copy the value of another cell\n");
        printf("This cell is out of range!\n");
        printf("copy operation was defined at: (%d, %d)\n", idxrow, idxcol);
        exit(1);
    }


    cell celltopoint = columns[i_rowtocp].child_s.at(i_coltocp);
    if ( celltopoint->type == CLONE ) {
        clone_value( celltopoint, i_rowtocp, i_coltocp );
    }
    tcell->value = celltopoint->value;
}

void start (const int _nrows, const int _ncols) {
    ncols = _ncols;
    nrows = _nrows;

    for (int i = 0; i < nrows; ++i) {
        for (int j = 0; j < ncols; ++j) {

            if ( columns[i].child_s[j]->type == CLONE ) {
                clone_value( columns[i].child_s.at(j), i, j );
            }

            if ( columns[i].child_s[j]->type == COPY ) {
                copy_value( columns[i].child_s.at(j), i, j );
            }

            printf("%d ", columns[i].child_s[j]->value);
        }
        printf("\n");
    }

}


#endif
