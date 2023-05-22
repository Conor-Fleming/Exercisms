#include "reverse_string.h"

namespace reverse_string
{
    string reverse_string(string input)
    {
        if (input.length() == 0)
        {
            return input;
        }

        string result;
        for (int i = input.length() - 1; i >= 0; i--)
        {
            result += input[i];
        }

        cout << result << endl;
        return result;
    }
}
