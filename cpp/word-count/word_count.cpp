#include "word_count.h"
#include<iterator>
#include<unordered_map>
#include<cctype>
#include<bits/stdc++.h>

using namespace std;


namespace word_count {

map<string, int> words(string input){
    unordered_map<string, int> wordCounts;
    input = tolower(input)
    stringstream s(input);
    string currentWord;
    while(s >> currentWord){
        unordered_map<string, int>::iterator it = wordCounts.find(currentWord);
        if(it != wordCounts.end())
            it->second++;
        else
            wordCounts.insert(std::make_pair(currentWord, 1));
    }
    return wordCounts
}

}  // namespace word_count
