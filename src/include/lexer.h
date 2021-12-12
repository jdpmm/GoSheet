#ifndef MEXCEL_LEXER_H
#define MEXCEL_LEXER_H

#include "cells.h"
#include "utiles.h"

int ncolumns_defined = 0;
int nrows_defined = 0;

void lexer (const std::string &line) {
    std::string cleanline = "";
    int idx = 0;

    while ( line[idx] != '\0' ) {
        while ( line[idx] == ' ' ) {
            idx++;
        }
        cleanline += line[idx];
        idx++;
    }

    std::vector<std::string> content = split_str(&cleanline, '|');
    if ( nrows_defined == 0 ) {
        ncolumns_defined = content.size();
    }

    if ( (int) content.size() != ncolumns_defined ) {
        printf("There are not the same number of values\n");
        printf("%d was defined on the first row\n", ncolumns_defined);
        printf("%d was defined on this row\n", (int) content.size());
        printf("ROW: %d\n", nrows_defined + 1);
        exit(1);
    }

    for (int i = 0; i < (int) content.size(); ++i) {
        set_cell(nrows_defined, i, content.at(i));
    }

    nrows_defined++;
}

void pre_start () {
    start(nrows_defined, ncolumns_defined);
}



#endif
