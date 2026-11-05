#include "zad_tree.h"
#include "../struct_cplas/tree/huffman_tree.h"
#include <iostream>
using namespace std;

void PrintHuffmanSymbol(unsigned char symbol){
    if (symbol == ' '){
        cout << "пробел";
    } else if (symbol > 32 && symbol < 127){
        cout << symbol;
    } else {
        cout << "байт " << (int)symbol;
    }
}

void RunHuffman(){
    string text;
    cout << "Введите строку:\n";
    getline(cin, text);
    if (text.empty()){
        cout << "Строка пустая\n";
        return;
    }
    int frequency[256]{};
    for (unsigned char symbol : text){
        ++frequency[symbol];
    }
    HufNode* nodes[256];
    int count = 0;
    for (int i = 0; i < 256; ++i){
        if (frequency[i] > 0){
            nodes[count++] = new HufNode{(unsigned char)i, frequency[i], nullptr, nullptr};
        }
    }
    HufNode* root = HufBuildTree(nodes, count);
    HashTable codes;
    HInitTable(codes, 8);
    HufBuildCodes(root, "", codes);
    cout << "Символ — частота — код:\n";
    for (int i = 0; i < 256; ++i){
        if (frequency[i] > 0){
            string code;
            HGet(codes, string(1, (char)i), code);
            PrintHuffmanSymbol((unsigned char)i);
            cout << " — " << frequency[i] << " — " << code << '\n';
        }
    }
    string bits = HufEncode(text, codes);
    string decoded;
    HufDecode(root, bits, decoded);
    cout << "Код: " << bits << '\n';
    cout << "Расшифровка: " << decoded << '\n';
    cout << "Введите код для расшифровки по этому дереву (пустая строка — пропустить):\n";
    if (getline(cin, bits) && !bits.empty()){
        if (HufDecode(root, bits, decoded)){
            cout << "Результат: " << decoded << '\n';
        } else {
            cout << "Ошибка: код содержит не 0/1 или обрывается внутри символа\n";
        }
    }
    HDestroy(codes);
    HufDeleteTree(root);
}
