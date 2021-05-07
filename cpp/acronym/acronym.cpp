#include "acronym.h"
#include<string>
#include<ctype.h>
#include<string>
#include<iostream>

using namespace std;

namespace acronym {

string acronym(string input){
    //input = input.replace(input.begin(), input.end(), '-', ' ');
    string output;
    output = input[0];
    for(int i = 0; i < int(input.length()); i++){
        if (input[i] == ' ' || input[i] == '-'){
            output += toupper(input[i+1]);
        }
    }
    return output;
}

}  // namespace acronym

