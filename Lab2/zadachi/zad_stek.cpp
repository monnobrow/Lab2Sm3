#include "zad_stek.h"
#include "../struct_cplas/Stack/stack.h"
#include "../utils/input.h"
#include <climits>
#include <sstream>
using namespace std;

bool BaseballScore(string operations, int& result){
    for (char& c : operations){
        if (c == '[' || c == ']' || c == ',') c = ' ';
    }
    Stack stack;
    SInitStack(&stack);
    int count = 0;
    bool correct = true;
    string operation;
    istringstream input(operations);
    while (input >> operation){
        int score = 0;
        if (operation == "C"){
            if (count == 0){ correct = false; break; }
            SPop(stack);
            --count;
            continue;
        } else if (operation == "D"){
            if (count == 0){ correct = false; break; }
            int last = stoi(STop(stack));
            if (last > INT_MAX / 2){ correct = false; break; }
            score = last * 2;
        } else if (operation == "+"){
            if (count < 2){ correct = false; break; }
            string second = SPop(stack);
            int first = stoi(STop(stack));
            int last = stoi(second);
            SPush(stack, second);
            if (first > INT_MAX - last){ correct = false; break; }
            score = first + last;
        } else if (!ReadPositiveInt(operation, score)){
            correct = false;
            break;
        }
        SPush(stack, to_string(score));
        ++count;
    }
    result = 0;
    if (correct){
        while (!SIsEmptyStack(&stack)){
            int score = stoi(SPop(stack));
            if (result > INT_MAX - score){ correct = false; break; }
            result += score;
        }
    }
    SClear(stack);
    return correct;
}

void RunBaseball(){
    cout << "Введите операции, например [5, 2, C, D, +]:\n";
    string operations;
    getline(cin, operations);
    int result;
    if (BaseballScore(operations, result)){
        cout << "Сумма очков: " << result << '\n';
    } else {
        cout << "Ошибка: неверная операция, недостаточно очков или переполнение int\n";
    }
}
