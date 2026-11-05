#pragma once
#include <string>
using namespace std;

enum OpenState{ OPEN_EMPTY, OPEN_OCCUPIED, OPEN_DELETED };

struct OpenCell{
    string key;
    string value;
    OpenState state = OPEN_EMPTY;
};

struct OpenTable{
    OpenCell* data = nullptr;
    int size = 0;
    int mesto = 0;
    int deleted = 0;
    bool double_hash = false;
    bool folding = false;
};

int FoldHash(const string& key, int mesto);
bool FoldingSum(const string& key, int& sum, string& formula);
void OInit(OpenTable& table, int mesto, bool double_hash, bool folding);
bool OInsert(OpenTable& table, const string& key, const string& value);
bool OGet(const OpenTable& table, const string& key, string& value);
bool ODelete(OpenTable& table, const string& key);
void OPrint(const OpenTable& table);
void ODestroy(OpenTable& table);
