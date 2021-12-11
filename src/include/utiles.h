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

int search_arithmetic (std::string *line, int sidx) {
    for (sidx; sidx < line->size(); ++sidx) {
        if ( line->at(sidx) == '+' || line->at(sidx) == '-' || line->at(sidx) == '*' || line->at(sidx) == '/' ) {
            return sidx;
        }
    }

    return -1;
}

int idx_of (const std::string *line, int from, char search) {
    int idx = from;
    while ( line->at(idx) != '\0' ) {
        if ( line->at(idx) == search ) {
            return idx;
        }
        idx++;
    }
    return -1;
}

int sum_v (std::vector<int> v) {
    int s = 0;
    for (int i = 0; i < v.size(); ++i) {
        s += v.at(i);
    }
    return s;
}


#endif
