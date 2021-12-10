#ifndef MEXCEL_UTILES_H
#define MEXCEL_UTILES_H

#include <string>
#include <vector>

std::string substr (std::string cont, int f, int t) {
    std::string ns = "";
    while ( f != t ) {
        ns += cont[f];
        f++;
    }
    return ns;
}

std::vector<std::string> split_str (std::string *line, char delimiter) {
    std::vector<std::string> nv;
    int p = 0;
    *line = substr(*line, 1, line->size());

    for (int i = 0; i < (int) line->size(); ++i) {
        if ( line->at(i) == delimiter ) {
            nv.push_back( substr(*line, p, i) );
            p = i + 1;
        }
    }
    return nv;
}

#endif
