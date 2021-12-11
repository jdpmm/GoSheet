#ifndef MEXCEL_MATHS_H
#define MEXCEL_MATHS_H

#include <vector>
#include <string>

double arithmetic (std::vector<char> *op, std::vector<int> *numbers) {
    if ( op->size() == 1 ) {
        if ( op->at(0) == '+' ) return numbers->at(0) + numbers->at(1);
        if ( op->at(0) == '-' ) return numbers->at(0) - numbers->at(1);
        if ( op->at(0) == '*' ) return numbers->at(0) * numbers->at(1);
        if ( op->at(0) == '/' ) return numbers->at(0) / numbers->at(1);
    }


    double res = numbers->at(0);
    for (int i = 0; i < (int) op->size(); ++i) {
        if ( op->at(i) == '+' ) {
            if ( op->at( i + 1 ) == '*' ) {
                res += numbers->at(i + 1) * numbers->at(i + 2);
                i+=1;
            }
            else if ( op->at(i + 1) == '/' ) {
                res += numbers->at(i + 1) / numbers->at(i + 2);
                i+=1;
            }
            else {
                res += numbers->at(i + 1);
            }
        }

        else if ( op->at(i) == '-' ) {
            if ( op->at( i + 1 ) == '*' ) {
                res -= numbers->at(i + 1) * numbers->at(i + 2);
                i+=1;
            }
            else if ( op->at(i + 1) == '/' ) {
                res -= numbers->at(i + 1) / numbers->at(i + 2);
                i+=1;
            }
            else {
                res -= numbers->at(i + 1);
            }
        }

        else if ( op->at(i) == '*' ) {
            res *= numbers->at(i + 1);
        }
        else {
            res /= numbers->at(i + 1);
        }
    }

    return res;
}


#endif
