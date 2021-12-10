#include "include/lexer.h"
#include <iostream>
#include <fstream>
using namespace std;

void read (const string &filename) {
    ifstream file (filename);
    string content = "";

    while ( !file.eof() ) {
        getline(file, content);
        if ( !content.empty() ) {
            lexer(content);
        }
    }

    pre_start();
}

int main (int argc, char *argv[]) {
    if ( argc == 2 ) {
        read ( string(argv[1]) );
    }
    return 0;
}
