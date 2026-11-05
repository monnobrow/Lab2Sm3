#include <iostream>
#include <string>
#include "zadachi/zad_stek.h"
#include "zadachi/zad_atd_mnozh.h"
#include "zadachi/zad_mnozh.h"
#include "zadachi/zad_array.h"
#include "zadachi/zad_tree.h"
#include "zadachi/zad_hesh_tab.h"
#include "zadachi/zad_lru.h"
using namespace std;

int RunTask(int number){
    switch (number){
        case 1: RunBaseball(); break;
        case 2: RunSetTask(); break;
        case 3: RunPartition(); break;
        case 4: RunSubsets(); break;
        case 5: RunHuffman(); break;
        case 6: RunHashTask(); break;
        case 7: RunLRU(); break;
        default: cout << "Нет такого задания\n"; return 1;
    }
    return 0;
}

void PrintMenu(){
    cout << "\nLab2 Ann — номер 19\n"
         << "1 — Бейсбол со стеком\n"
         << "2 — АТД Множество\n"
         << "3 — Разбиение множества\n"
         << "4 — Все подмножества массива\n"
         << "5 — Хаффман\n"
         << "6 — Линейное и двойное хеширование, свёртка\n"
         << "7 — LRU-кэш\n0 — Выход\n> ";
}

int main(int argc, char* argv[]){
    if (argc == 2 && string(argv[1]) == "--help"){
        cout << "lab2 — меню\nlab2 1..7 — отдельное задание\n"
             << "lab2 --file set.txt --query 'SETADD 5' — команда множества\n";
        return 0;
    }
    if (argc == 2 && string(argv[1]).size() == 1 && argv[1][0] >= '1' && argv[1][0] <= '7'){
        return RunTask(argv[1][0] - '0');
    }
    if (argc > 1) return RunMnozh(argc, argv);
    while (true){
        PrintMenu();
        int number;
        if (!(cin >> number)) break;
        string rest; getline(cin, rest);
        if (number == 0) break;
        RunTask(number);
        if (!cin){
            if (cin.eof()) break;
            cin.clear(); getline(cin, rest);
        }
    }
    return 0;
}
