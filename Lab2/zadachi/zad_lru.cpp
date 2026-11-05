#include "zad_lru.h"
#include "../struct_cplas/cache/lru_cache.h"
using namespace std;

void RunLRU(){
    int capacity, q;
    cout << "Ёмкость кэша и количество запросов: ";
    if (!(cin >> capacity >> q) || capacity < 0 || q < 0){
        cout << "Ошибка: нужны неотрицательные числа\n";
        return;
    }
    LRUCache cache(capacity);
    cout << "Введите запросы SET ключ значение или GET ключ:\n";
    string result;
    for (int i = 0; i < q; ++i){
        string command, key, value;
        if (!(cin >> command >> key)){ cout << "Ошибка: не хватает запросов\n"; return; }
        if (command == "SET"){
            if (!(cin >> value)){ cout << "Ошибка: нужно значение\n"; return; }
            cache.Set(key, value);
        } else if (command == "GET"){
            if (!result.empty()) result += ' ';
            result += cache.Get(key);
        } else { cout << "Ошибка: неизвестная команда\n"; return; }
    }
    cout << "Результат: " << result << '\n';
}
