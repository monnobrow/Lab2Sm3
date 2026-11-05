#include "zad_atd_mnozh.h"
#include "../struct_cplas/Mnozh/mnozh.h"
#include <fstream>
#include <sstream>
using namespace std;

// В файле каждый элемент занимает отдельную строку.
bool LoadSet(Set& set, const string& filename){
    ifstream file(filename);
    if (!file.is_open()){
        return true; // Если файл не открылся, начинаем с пустого множества.
    }
    string value;
    while (getline(file, value)){
        if (!value.empty() && value.back() == '\r'){
            value.pop_back();
        }
        if (!value.empty()){
            SetAdd(set, value);
        }
    }
    if (file.bad()){
        cerr << "Ошибка: не удалось прочитать файл " << filename << '\n';
        return false;
    }
    return true;
}

bool SaveSet(const Set& set, const string& filename){
    ofstream file(filename);
    if (!file.is_open()){
        cerr << "Ошибка: не удалось записать файл " << filename << '\n';
        return false;
    }
    Array elements;
    MInitArray(elements, SetSize(set));
    SetElements(set, elements);
    for (int i = 0; i < MLength(elements); ++i){
        file << MGet(elements, i) << '\n';
    }
    MClear(elements);
    file.close();
    if (!file){
        cerr << "Ошибка: запись файла не завершена " << filename << '\n';
        return false;
    }
    return true;
}

int RunSetQuery(const string& filename, const string& query){
    istringstream input(query);
    string command, value, extra;
    input >> command;
    if (command != "SETADD" && command != "SETDEL" && command != "SET_AT" && command != "SETPRINT"){
        cerr << "Ошибка: неизвестная команда " << command << '\n';
        return 1;
    }
    if (command != "SETPRINT" && !(input >> value)){
        cerr << "Ошибка: нужно значение элемента\n";
        return 1;
    }
    if (input >> extra){
        cerr << "Ошибка: слишком много аргументов команды\n";
        return 1;
    }
    Set set;
    SetInit(set);
    if (!LoadSet(set, filename)){
        SetDestroy(set);
        return 1;
    }
    int result = 0;
    if (command == "SETADD"){
        if (!SetAdd(set, value)){
            cout << value << " уже есть в множестве\n";
        } else if (SaveSet(set, filename)){
            cout << "-> " << value << '\n';
        } else {
            result = 1;
        }
    } else if (command == "SETDEL"){
        if (!SetDel(set, value)){
            cout << value << " нет в множестве\n";
        } else if (SaveSet(set, filename)){
            cout << "-> " << value << " удалён\n";
        } else {
            result = 1;
        }
    } else if (command == "SET_AT"){
        cout << "-> " << (SetAt(set, value) ? "TRUE" : "FALSE") << '\n';
    } else {
        SetPrint(set);
    }
    SetDestroy(set);
    return result;
}


int RunMnozh(int argc, char* argv[]){
    string filename, query;
    for (int i = 1; i < argc; ++i){
        string arg = argv[i];
        if (arg == "--file" && i + 1 < argc){
            filename = argv[++i];
        } else if (arg == "--query" && i + 1 < argc){
            query = argv[++i];
        } else {
            cerr << "Ошибка: неизвестный или неполный параметр " << arg << '\n';
            return 1;
        }
    }
    if (filename.empty() || query.empty()){
        cerr << "Запуск: lab2 --file set.txt --query 'SETADD x | SETDEL x | SET_AT x'\n";
        return 1;
    }
    return RunSetQuery(filename, query);
}

void RunSetTask(){
    string filename, query;
    cout << "Файл множества: ";
    getline(cin >> ws, filename);
    cout << "Команда (SETADD x / SETDEL x / SET_AT x / SETPRINT): ";
    getline(cin, query);
    RunSetQuery(filename, query);
}
