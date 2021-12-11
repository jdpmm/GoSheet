#ifndef MEXCEL_CELLS_H
#define MEXCEL_CELLS_H

#include "utiles.h"
#include "maths.h"
#define A_ASCII_POS 65

typedef enum {
    NUMBER,
    CLONE,
    COPY,
    ARITHMETIC,
    OPERATION
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
void make_maths (cell tcell);

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

    if ( icol >= ncolumns_definitive || irow > nrows_definitive ) {
        printf("Tyring to get the value of another cell\n");
        printf("The cell is out of range!\n");
        exit(1);
    }

    return columns[irow].nodes.at(icol);
}

void get_values_range_cols (char column, int init, int end, std::vector<int> *v) {
    int idxcolumn = column - A_ASCII_POS;
    while ( init <= end ) {
        v->push_back( columns[init].nodes[idxcolumn]->value );
        init++;
    }
}

void get_values_range_rows (int row, char init, char end, std::vector<int> *v) {
    init = init - A_ASCII_POS;
    while ( init <= (end - A_ASCII_POS) ) {
        v->push_back( columns[row].nodes[init]->value );
        init++;
    }
}

void realize_operation (cell tcell) {
    // tcell->vaat == =ope(A1:B4)
    std::string operation = substr(tcell->vaat, 1, 4);
    int index_colon = idx_of(&tcell->vaat, 0, ':');

    if ( index_colon == -1 ) {
        printf("The operation has fail\n");
        printf("':' excepted");
        printf("The operation has been defined at: (%d, %d)\n", tcell->coord[0] + 1, tcell->coord[1] + 1);
        exit(1);
    }
    std::string coor1 = substr(tcell->vaat, 5, index_colon);
    std::string coor2 = substr(tcell->vaat, index_colon + 1, tcell->vaat.size() - 1);

    int row1 = stoi( substr(coor1, 1, coor1.size()) ) - 1;
    int row2 = stoi( substr(coor2, 1, coor2.size()) ) - 1;

    // save the values from one cell to another cell
    std::vector<int> values;

    if ( coor1[0] == coor2[0] && (coor1[0] - A_ASCII_POS) < ncolumns_definitive ) {
        get_values_range_cols(coor1[0], row1, row2, &values);
    }
    else if ( row1 == row2 && row1 < nrows_definitive ) {
        get_values_range_rows(row1, coor1[0], coor2[0], &values);
    }
    else {
        printf("\nError!!\n");
        printf("Index out range on operation!\n");
        printf("The operation has been defined at: (%d, %d)\n", tcell->coord[0] + 1, tcell->coord[1] + 1);
        exit(1);
    }

    if ( operation == "sum" ) {
        tcell->value = sum_v(values);
    }
    else if ( operation == "max" ) {
        tcell->value = max_v(values);
    }
    else if ( operation == "min" ) {
        tcell->value = min_v(values);
    }
    else if ( operation == "moa" ) {

    }
    else if ( operation == "med" ) {
        tcell->value = sum_v(values) / values.size();
    }
    else if ( operation == "mda" ) {

    }
}

void make_maths (cell tcell) {
    std::vector<char> operations;
    std::vector<int> numbers;

    int idx = 1;
    std::string number = "";
    while ( true ) {
        char _ = tcell->vaat[idx];
        if ( _ == '-' && isdigit(tcell->vaat[idx + 1]) ) {
            number += "-";
            number += tcell->vaat[idx + 1];
            operations.push_back( tcell->vaat[idx] );

            idx+=2;
        }

        if ( _ == '+' || _ == '-' || _ == '/' || _ == '*' || _ == '\0' ) {
            if ( !isdigit( number[0] ) && number[0] != '-' ) {
                numbers.push_back( get_cell_by_coord(number)->value );
            }
            else {
                numbers.push_back( stoi(number) );
            }

            operations.push_back( tcell->vaat[idx] );
            number = "";
            idx++;
        }

        if ( tcell->vaat[idx] == '\0' ) {
            break;
        }

        number+=tcell->vaat[idx];
        idx++;
    }

    operations.erase( operations.end() - 1 );
    tcell->value = (int) arithmetic (&operations, &numbers);
    tcell->value = NUMBER;
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
        // int search_arithmetic (std::string *line, int sidx) {
        if ( search_arithmetic(&value, 1) == -1 ) {
            newcell->type = COPY;
        }
        else {
            if ( value.size() >= 3  ) {
                newcell->type = ARITHMETIC;
            }
            else {
                printf("Trying to make maths!\n");
                printf("There aren't the enough number or operators\n");
                printf("(%d, %d)\n", i_row + 1, i_col + 1);
                exit(1);
            }
        }

        newcell->vaat = value;
    }
    else if ( value[0] == '=' ) {
        newcell->type = OPERATION;
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

            if ( columns[i].nodes[j]->type == ARITHMETIC ) {
                make_maths(columns[i].nodes[j]);
            }

            if ( columns[i].nodes[j]->type == OPERATION ) {
                realize_operation(columns[i].nodes[j]);
            }

            printf("%d ", columns[i].nodes[j]->value);
        }
        printf("\n");
    }

}


#endif

