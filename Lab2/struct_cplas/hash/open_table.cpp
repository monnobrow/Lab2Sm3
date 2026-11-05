#include "open_table.h"
#include "hesh_table.h"
#include <climits>
#include <iostream>

bool IsPrime(int value){
    if (value < 2) return false;
    for (int i = 2; i <= value / i; ++i){
        if (value % i == 0) return false;
    }
    return true;
}

int NextPrime(int value){
    if (value < 3) value = 3;
    while (!IsPrime(value)) ++value;
    return value;
}

int FoldHash(const string& key, int mesto){
    unsigned int result = 0;
    for (int i = 0; i < (int)key.size(); i += 3){
        int part = 0;
        for (int j = i; j < i + 3 && j < (int)key.size(); ++j){
            if (key[j] < '0' || key[j] > '9') return HHash(key, mesto);
            part = part * 10 + key[j] - '0';
        }
        result = (result + part) % mesto;
    }
    return result;
}

bool FoldingSum(const string& key, int& sum, string& formula){
    sum = 0;
    formula.clear();
    if (key.empty()) return false;
    for (int i = 0; i < (int)key.size(); i += 3){
        int part = 0;
        string digits;
        for (int j = i; j < i + 3 && j < (int)key.size(); ++j){
            if (key[j] < '0' || key[j] > '9') return false;
            digits += key[j];
            part = part * 10 + key[j] - '0';
        }
        if (sum > INT_MAX - part) return false;
        sum += part;
        if (!formula.empty()) formula += '+';
        formula += digits;
    }
    return true;
}

int OpenStart(const OpenTable& table, const string& key){
    return table.folding ? FoldHash(key, table.mesto) : HHash(key, table.mesto);
}

int OpenStep(const OpenTable& table, const string& key){
    if (!table.double_hash) return 1;
    unsigned int hash = 2166136261u;
    for (unsigned char c : key){ hash ^= c; hash *= 16777619u; }
    return 1 + hash % (table.mesto - 1);
}

void OInit(OpenTable& table, int mesto, bool double_hash, bool folding){
    ODestroy(table);
    table.mesto = NextPrime(mesto);
    table.data = new OpenCell[table.mesto];
    table.double_hash = double_hash;
    table.folding = folding;
}

int OpenFind(const OpenTable& table, const string& key){
    if (table.data == nullptr) return -1;
    int index = OpenStart(table, key), step = OpenStep(table, key);
    for (int i = 0; i < table.mesto; ++i){
        if (table.data[index].state == OPEN_EMPTY) return -1;
        if (table.data[index].state == OPEN_OCCUPIED && table.data[index].key == key){
            return index;
        }
        index = (index + step) % table.mesto;
    }
    return -1;
}

void OpenRehash(OpenTable& table, int mesto){
    OpenTable next;
    OInit(next, mesto, table.double_hash, table.folding);
    for (int i = 0; i < table.mesto; ++i){
        if (table.data[i].state == OPEN_OCCUPIED){
            int index = OpenStart(next, table.data[i].key);
            int step = OpenStep(next, table.data[i].key);
            while (next.data[index].state == OPEN_OCCUPIED){
                index = (index + step) % next.mesto;
            }
            next.data[index].key = table.data[i].key;
            next.data[index].value = table.data[i].value;
            next.data[index].state = OPEN_OCCUPIED;
            ++next.size;
        }
    }
    delete[] table.data;
    table = next;
}

bool OInsert(OpenTable& table, const string& key, const string& value){
    if (table.data == nullptr) OInit(table, 11, false, false);
    int existing = OpenFind(table, key);
    if (existing >= 0){ table.data[existing].value = value; return false; }
    if (table.size + table.deleted + 1 >= table.mesto - table.mesto / 4){
        int mesto = table.size + 1 >= table.mesto / 2 ? table.mesto * 2 : table.mesto;
        OpenRehash(table, mesto);
    }
    int index = OpenStart(table, key), step = OpenStep(table, key);
    while (table.data[index].state == OPEN_OCCUPIED){ index = (index + step) % table.mesto; }
    if (table.data[index].state == OPEN_DELETED) --table.deleted;
    table.data[index].key = key;
    table.data[index].value = value;
    table.data[index].state = OPEN_OCCUPIED;
    ++table.size;
    return true;
}

bool OGet(const OpenTable& table, const string& key, string& value){
    int index = OpenFind(table, key);
    if (index < 0) return false;
    value = table.data[index].value;
    return true;
}

bool ODelete(OpenTable& table, const string& key){
    int index = OpenFind(table, key);
    if (index < 0) return false;
    table.data[index].state = OPEN_DELETED;
    table.data[index].key.clear();
    table.data[index].value.clear();
    --table.size;
    ++table.deleted;
    return true;
}

void OPrint(const OpenTable& table){
    for (int i = 0; i < table.mesto; ++i){
        cout << '[' << i << "] ";
        if (table.data[i].state == OPEN_OCCUPIED){
            cout << table.data[i].key << " -> " << table.data[i].value;
        } else if (table.data[i].state == OPEN_DELETED){ cout << "удалено"; }
        else cout << '-';
        cout << '\n';
    }
}

void ODestroy(OpenTable& table){
    delete[] table.data;
    table = OpenTable{};
}
