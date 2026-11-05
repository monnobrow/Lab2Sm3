#pragma once
#include <iostream>
#include <string>
using namespace std;

struct SNode{
    string chislo;
    SNode* next;
};
struct Stack{
    SNode* head;
};

void SInitStack(Stack* stack);
bool SIsEmptyStack(Stack* stack);
void SPush(Stack& stack, string chislo);
string SPop(Stack& stack); // снимает вершину и сразу отдаёт её значение
string STop(const Stack& stack);
void SPrint(SNode* flist);
void SPrintAddresses(const Stack& stack);
void SClear(Stack& stack);
