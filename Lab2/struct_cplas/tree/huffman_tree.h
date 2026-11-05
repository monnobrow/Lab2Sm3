#pragma once
#include <iostream>
#include <string>
#include "../hash/hesh_table.h"
using namespace std;

// дерево Хаффмана: символы только в листьях, у внутренних узлов только вес (сумма весов детей).
// путь от корня до листа = код символа: влево 0, вправо 1
struct HufNode{
    unsigned char symbol;
    int weight;
    HufNode* left;
    HufNode* right;
};

bool HufIsLeaf(HufNode* node);
HufNode* HufBuildTree(HufNode** nodes, int count);
void HufBuildCodes(HufNode* node, string code, HashTable& codes);
string HufEncode(const string& text, const HashTable& codes);
bool HufDecode(HufNode* root, const string& bits, string& text);
void HufDeleteTree(HufNode* node);
