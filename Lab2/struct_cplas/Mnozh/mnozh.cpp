#include "mnozh.h"

void SetInit(Set& set){
    HInitTable(set.table, 8);
}

bool SetAt(const Set& set, const string& value){
    return HContains(set.table, value);
}

bool SetAdd(Set& set, const string& value){
    return HInsert(set.table, value, "");
}

bool SetDel(Set& set, const string& value){
    return HDelete(set.table, value);
}

int SetSize(const Set& set){
    return HLength(set.table);
}

void SetElements(const Set& set, Array& elements){
    for (int i = 0; i < set.table.mesto; ++i){
        HNode* node = set.table.data[i];
        while (node != nullptr){
            MAddEnd(elements, node->key);
            node = node->next;
        }
    }
}

void SetPrint(const Set& set, ostream& out){
    out << '{';
    bool first = true;
    for (int i = 0; i < set.table.mesto; ++i){
        HNode* node = set.table.data[i];
        while (node != nullptr){
            if (!first) out << ", ";
            out << node->key;
            first = false;
            node = node->next;
        }
    }
    out << "}\n";
}

void SetClear(Set& set){
    HClear(set.table);
}

void SetDestroy(Set& set){
    HDestroy(set.table);
}
