#pragma once
#include "../hash/hesh_table.h"
#include "../Array/massiv.h"

struct Set{
    HashTable table;
};

void SetInit(Set& set);
bool SetAt(const Set& set, const string& value);
bool SetAdd(Set& set, const string& value);
bool SetDel(Set& set, const string& value);
int SetSize(const Set& set);
void SetElements(const Set& set, Array& elements);
void SetPrint(const Set& set, ostream& out = cout);
void SetClear(Set& set);
void SetDestroy(Set& set);
