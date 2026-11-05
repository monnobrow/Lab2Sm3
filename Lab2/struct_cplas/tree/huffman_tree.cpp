#include <iostream>
#include <string>
#include "huffman_tree.h"
using namespace std;

bool HufIsLeaf(HufNode* node){
    return node->left == nullptr && node->right == nullptr;
}

// nodes - список свободных узлов (сначала по листу на символ).
// пока узлов больше одного: берём два с наименьшим весом и заменяем их родителем.
// оставшийся узел - корень. O(k^2), k - число различных символов (не больше 256)
HufNode* HufBuildTree(HufNode** nodes, int count){
    if (count == 0) return nullptr;
    while (count > 1){
        int min1 = 0; // индекс самого лёгкого
        for (int i = 1; i < count; ++i){
            if (nodes[i]->weight < nodes[min1]->weight){
                min1 = i;
            }
        }
        int min2 = -1; // индекс второго по лёгкости
        for (int i = 0; i < count; ++i){
            if (i != min1 && (min2 == -1 || nodes[i]->weight < nodes[min2]->weight)){
                min2 = i;
            }
        }
        // левая дуга (0) - к самому лёгкому, правая (1) - ко второму
        HufNode* parent = new HufNode{0, nodes[min1]->weight + nodes[min2]->weight, nodes[min1], nodes[min2]};
        nodes[min1] = parent;           // родитель встаёт на место первого
        nodes[min2] = nodes[count - 1]; // на место второго - последний узел списка
        count--;
    }
    return nodes[0];
}

// обход в глубину: влево дописываем 0, вправо 1; дошли до листа - код готов.
// коды кладём в свою хеш-таблицу: символ -> код
void HufBuildCodes(HufNode* node, string code, HashTable& codes){
    if (node == nullptr) return;
    if (HufIsLeaf(node)){
        if (code.empty()){ // в строке один различный символ: дерево из одного листа
            code = "0";
        }
        HInsert(codes, string(1, node->symbol), code);
        return;
    }
    HufBuildCodes(node->left, code + "0", codes);
    HufBuildCodes(node->right, code + "1", codes);
}

string HufEncode(const string& text, const HashTable& codes){
    string bits;
    for (char c : text){
        string code;
        HGet(codes, string(1, c), code);
        bits += code;
    }
    return bits;
}

// идём от корня по битам; в листе записываем символ и возвращаемся в корень.
// false - в коде есть не 0/1 или он оборван посередине символа
bool HufDecode(HufNode* root, const string& bits, string& text){
    text = "";
    if (root == nullptr) return bits.empty();
    if (HufIsLeaf(root)){ // один символ: каждый бит 0 - это он
        for (char b : bits){
            if (b != '0'){
                return false;
            }
            text += root->symbol;
        }
        return true;
    }
    HufNode* cur_node = root;
    for (char b : bits){
        if (b == '0'){
            cur_node = cur_node->left;
        } else if (b == '1'){
            cur_node = cur_node->right;
        } else {
            return false;
        }
        if (HufIsLeaf(cur_node)){
            text += cur_node->symbol;
            cur_node = root;
        }
    }
    return cur_node == root;
}

void HufDeleteTree(HufNode* node){
    if (node == nullptr) return;
    HufDeleteTree(node->left);
    HufDeleteTree(node->right);
    delete node;
}
