#include "bob.h"
#include<iostream>
#include<string>
#include<cctype>

using namespace std;
using std::string;

namespace bob {
string hey(string input) { 
    if(isupper(input[2])){
        if(input[input.length()] == '?')
            return "Calm down, I know what I'm doing!";
        else
           return "Whoa, chill out!";
    }else if(!isupper(input[2] || (input[1] != ' ' && input[input.length()] != ' '))){
        if(input[input.length()] == '?')
             return "Sure.";
        else
            return "Whatever.";
    }else{
        return "Fine. Be that way!";
    }
}
}  // namespace bob
