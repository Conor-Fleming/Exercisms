#if !defined(WORD_COUNT_H)
#define WORD_COUNT_H
#include<string>
#include<unordered_map>

using namespace std;

namespace word_count {
unordered_map<string, int> words(string);
}  // namespace word_count

#endif // WORD_COUNT_H