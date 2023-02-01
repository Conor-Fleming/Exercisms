#include "reverse_string.h"

namespace reverse_string
{
    string reverse_string(string input)
    {
        string result = "";
        for (int i = input.length(); i > 0; i--)
        {
            result += input[i];
        }

        return result;
    }
} // namespace reverse_string
