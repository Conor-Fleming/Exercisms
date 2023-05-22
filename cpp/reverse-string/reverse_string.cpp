#include "reverse_string.h"

namespace reverse_string
{
    string reverse_string(string input)
    {
        if (input.length() == 0)
        {
            return input;
        }
        string result = "\0";
        for (int i = input.length(); i >= 0; i--)
        {
            result += tolower(input[i]);
        }

        return result;
    }
}
