#include "word_count.h"
#include<iterator>
#include<map>
#include<cctype>
#include<sstream>
#include<string>
#include<algorithm>

using namespace std;


namespace word_count {

map<string, int> words(string input){
    map<string, int> wordCounts;
    transform(input.begin(), input.end(), input.begin(), ::tolower);
    stringstream s(input);
    string currentWord;
    while(s >> currentWord){
        map<string, int>::iterator it = wordCounts.find(currentWord);
        if(it != wordCounts.end())
            it->second++;
        else
            wordCounts.insert(std::make_pair(currentWord, 1));
    }
    return wordCounts;
}

}  // namespace word_count
