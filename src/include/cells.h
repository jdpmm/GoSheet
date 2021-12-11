#ifndef MEXCEL_CELLS_H
#define MEXCEL_CELLS_H

#include "utiles.h"
#define A_ASCII_POS 65

typedef enum {
    NUMBER,
    CLONE,
    COPY
} Type;

typedef struct CELL {
    int value { 0 };
    Type type { NUMBER };
    int coord[2];

    CELL* noloops { nullptr };
    std::string vaat { "-" };
} *cell;

typedef struct COL {
    std::vector<cell> nodes;
} col;

col columns[26];
int ncolumns_definitive;
int nrows_definitive;

void clone_value (cell tcell, int irow, int icol);
void copy_value (cell tcell);

void check_no_loop (cell currentc) {
    cell auxnoloop = currentc->noloops->noloops;
    if ( auxnoloop != nullptr ) {
        bool causeloop = auxnoloop->coord[0] == currentc->coord[0] && auxnoloop->coord[1] == currentc->coord[1];

        if ( causeloop ) {
            printf("Loop detected\n");
            printf("coord: (%d, %d)\n", currentc->coord[0] + 1, currentc->coord[1] + 1 );
            exit(1);
        }
    }
}

cell get_cell_by_coord (const std::string _ord) {
    const int pre_icol = _ord[0] == ':' ? 1 : 0;
    const int pre_irow = _ord[0] == ':' ? 2 : 1;

    const int icol = _ord[pre_icol] - A_ASCII_POS;
    const int irow = stoi( substr(_ord, pre_irow, _ord.size()) ) - 1;

    if ( icol > ncolumns_definitive || irow > nrows_definitive ) {
        printf("Tyring to get the value of another cell\n");
        printf("The cell is out of range!\n");
        exit(1);
    }

    return columns[irow].nodes.at(icol);
}

void copy_value (cell tcell) {
    // tcell->vaat == [CHAR COLUMN][N ROW]: A1
    // tcell->vaat = position of the cell to copy
    cell ctocopy = get_cell_by_coord(tcell->vaat);

    if ( ctocopy->type != NUMBER ) {
        tcell->noloops = ctocopy;
        check_no_loop(tcell);

        if ( ctocopy->type == COPY ) {
            copy_value(ctocopy);
        }
        else if ( ctocopy->type == CLONE ) {
            clone_value(ctocopy, ctocopy->coord[0], ctocopy->coord[1]);
        }
    }

    tcell->value = ctocopy->value;
    tcell->type = NUMBER;
    tcell->noloops = nullptr;
}

void clone_value (cell tcell, int irow, int icol) {
    char oper = tcell->vaat[0];
    cell toclone;

    try {
        if ( oper == '^' && irow > 0 ) irow--;
        else if ( oper == 'v' && irow < (nrows_definitive - 1) ) irow++;
        else if ( oper == '>' && icol < (ncolumns_definitive - 1) ) icol++;
        else if ( oper == '<' && irow > 0 ) icol--;

        else {
            throw "Out of range";
        }
    }
    catch (const char* error) {
        printf("\nTrying to clone one value\n");
        printf("%s\n", error);
        printf("-> (%d, %d)\n", irow, icol);
        exit(1);
    }

    toclone = columns[irow].nodes.at(icol);
    if ( toclone->type != NUMBER ) {
        tcell->noloops = toclone;
        check_no_loop(tcell);

        if ( toclone->type == CLONE ) {
            clone_value(toclone, irow, icol);
        }
        else if ( toclone->type == COPY ) {
            copy_value(toclone);
        }
    }

    tcell->value = toclone->value;
    tcell->type = NUMBER;
    tcell->noloops = nullptr;
}

void set_cell (int i_row, int i_col, std::string value) {
    cell newcell = new (struct CELL);
    newcell->coord[0] = i_row;
    newcell->coord[1] = i_col;

    if ( value == "v" || value == "^" || value == ">" || value == "<" ) {
        newcell->type = CLONE;
        newcell->vaat = value;
    }
    else if ( value[0] == ':' ) {
        newcell->type = COPY;
        newcell->vaat = value;
    }
    else {
        newcell->type = NUMBER;
        newcell->value = stoi(value);
    }

    // i_row as index because i_row (nrows_defined as paramter) inc his value for each row on the table,
    // so if we will use i_col (i as paramter) the table will be vertial (0, 1) -> (1, 0)
    // because i_col inc his value for each value on the row
    columns[i_row].nodes.push_back(newcell);
}

void start (const int _nrows, const int _ncols) {
    ncolumns_definitive = _ncols;
    nrows_definitive = _nrows;

    for (int i = 0; i < nrows_definitive; ++i) {
        for (int j = 0; j < ncolumns_definitive; ++j) {

            if ( columns[i].nodes[j]->type == CLONE ) {
                clone_value( columns[i].nodes.at(j), i, j );
            }

            if ( columns[i].nodes[j]->type == COPY ) {
                copy_value( columns[i].nodes.at(j) );
            }

            /*
            if ( columns[i].child_s[j]->type == ARITHMETIC ) {
                printf("(%d, %d)\n", i, j);
            }*/

            printf("%d ", columns[i].nodes[j]->value);
        }
        printf("\n");
    }

}


#endif



/*
#include "utiles.h"
#define A_ASCII_POS 65

typedef enum {
    NUMBER,
    CLONE,
    COPY,
    ARITHMETIC
} c_Type;

typedef struct CELL {
    int value { 0 };
    c_Type type { NUMBER };
    int coord[2];
    int coordnl[2];
    CELL* noloops { nullptr };

    // value as another type, if isn't a number yet
    std::string vaot { "-" };
} *cell;

typedef struct COL {
    std::vector<cell> child_s;
} col;

col columns[26];
int ncols;
int nrows;
cell detect_loop = nullptr;

void clone_value (cell tcell, int idxrow, int idxcol);
cell get_cell_by_coord (const std::string coord);
void copy_value (cell tcell);

void noloops_check (cell current, cell couldbeloop) {

}

void set_cell (int i_row, int i_col, std::string value) {
    cell newcell = new (struct CELL);
    newcell->coord[0] = i_row;
    newcell->coord[1] = i_col;

    if ( value == "v" || value == "^" || value == ">" || value == "<" ) {
        newcell->type = CLONE;
        newcell->vaot = value;
    }
    else if ( value[0] == ':' ) {
        if ( search_arithmetic(&value, 0) != -1 ) {
            newcell->type = ARITHMETIC;
            newcell->vaot = value;
        }
        else {
            newcell->type = COPY;
            newcell->vaot = value;
        }
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

    toclone = columns[idxrow].child_s.at(idxcol);

    if ( toclone->type == COPY ) {
        copy_value(toclone);
    }

    tcell->value = toclone->value;
    tcell->type = NUMBER;
}

void copy_value (cell tcell) {
    cell cellto_cp = get_cell_by_coord(tcell->vaot);

    if ( tcell->noloops == nullptr ) {
        tcell->noloops = cellto_cp;
    }

    if ( cellto_cp->type == CLONE ) {
        clone_value ( cellto_cp, cellto_cp->coord[0], cellto_cp->coord[1] );
    }
    if ( cellto_cp->type == COPY ) {
        copy_value(cellto_cp);
    }

    tcell->value = cellto_cp->value;
    tcell->type = NUMBER;
}




#endif
*/
