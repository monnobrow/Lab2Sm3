#pragma once
#include <iostream>
#include <string>
using namespace std;

struct LNode{
    string chislo;
    LNode* next, *prev;
};
struct LList{
    LNode* head;
    LNode* tail;
};

void LInitList(LList& list);

// поиск target делается внутри операции: один проход, потом перевязка за O(1)
// результат ("-> x") или ошибку операция печатает сама
LNode* LSearch(LNode* flist, string chislo);
void LLinkBefore(LList& list, LNode* flist, string chislo);
void LLinkAfter(LList& list, LNode* flist, string chislo);
void LDeleteNode(LList& list, LNode* flist);
void LAddFirst(LList& list, string chislo);
void LAddEnd(LList& list, string chislo);
void LAddAfter(LList& list, string target, string chislo);
void LAddBefore(LList& list, string target, string chislo);
void LAddIndex(LList& list, string chislo, int index);
void LDeleteFirst(LList& list);
void LDeleteEnd(LList& list);
void LDeleteAfter(LList& list, string target);
void LDeleteBefore(LList& list, string target);
void LDeleteIndex(LList& list, int index);
void LDeleteValue(LList& list, string chislo);
void LGet(const LList& list, int index);
void LFind(const LList& list, string chislo);
void LPrint(const LList& list);
void LPrintBack(const LList& list);
void LPrintAddresses(const LList& list);
void LClear(LList& list);
