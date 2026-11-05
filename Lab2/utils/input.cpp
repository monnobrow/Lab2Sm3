#include "input.h"
#include <climits>

bool ReadPositiveInt(const std::string& text, int& value){
    if (text.empty()) return false;
    value = 0;
    for (char c : text){
        if (c < '0' || c > '9') return false;
        int digit = c - '0';
        if (value > (INT_MAX - digit) / 10) return false;
        value = value * 10 + digit;
    }
    return value > 0;
}
