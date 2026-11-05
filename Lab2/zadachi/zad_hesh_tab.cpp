#include "zad_hesh_tab.h"
#include "../struct_cplas/hash/open_table.h"
#include "../struct_cplas/Array/massiv.h"
#include <ctime>
#include <iomanip>
#include <sstream>
using namespace std;

double OpenOperationTime(clock_t start, int n){
    return (double)(clock() - start) * 1000000 / CLOCKS_PER_SEC / n;
}

void MeasureOpenHash(int n, bool double_hash, double* times){
    Array keys, missing;
    MInitArray(keys, n); MInitArray(missing, n);
    for (int i = 0; i < n; ++i){
        MAddEnd(keys, to_string(i * 37 + 11));
        MAddEnd(missing, to_string((i + n) * 37 + 11));
    }
    OpenTable table;
    OInit(table, 11, double_hash, false);
    clock_t start = clock();
    for (int i = 0; i < n; ++i) OInsert(table, MGet(keys, i), "value");
    times[0] = OpenOperationTime(start, n);
    int found = 0;
    string value;
    start = clock();
    for (int i = 0; i < n; ++i) found += OGet(table, MGet(keys, i), value);
    times[1] = OpenOperationTime(start, n);
    start = clock();
    for (int i = 0; i < n; ++i) found += OGet(table, MGet(missing, i), value);
    times[2] = OpenOperationTime(start, n);
    start = clock();
    for (int i = 0; i < n; ++i) ODelete(table, MGet(keys, i));
    times[3] = OpenOperationTime(start, n);
    if (found != n || table.size != 0) cout << "Ошибка: таблица работает неправильно\n";
    ODestroy(table); MClear(keys); MClear(missing);
}

void BenchmarkOpenHash(int n){
    double linear[4], double_hash[4];
    MeasureOpenHash(n, false, linear);
    MeasureOpenHash(n, true, double_hash);
    string names[4] = {"Добавление", "Поиск существующего", "Поиск отсутствующего", "Удаление"};
    cout << "\nN = " << n << ", микросекунды на операцию\n";
    cout << "Операция | Линейное | Двойное\n" << fixed << setprecision(3);
    for (int i = 0; i < 4; ++i){
        cout << names[i] << " | " << linear[i] << " | " << double_hash[i] << '\n';
    }
}

void DemoFolding(){
    string key, formula;
    cout << "Введите числовой ключ: ";
    getline(cin, key);
    int sum;
    if (!FoldingSum(key, sum, formula)){
        cout << "Ошибка: нужен числовой ключ с суммой частей в пределах int\n";
        return;
    }
    cout << formula << " = " << sum << '\n';
    OpenTable linear, double_hash;
    OInit(linear, 11, false, true);
    OInit(double_hash, 11, true, true);
    OInsert(linear, key, to_string(sum));
    OInsert(double_hash, key, to_string(sum));
    string value;
    OGet(linear, key, value);
    cout << "Линейное: найдено " << value << ", начальная ячейка " << FoldHash(key, linear.mesto) << '\n';
    OGet(double_hash, key, value);
    cout << "Двойное: найдено " << value << ", начальная ячейка " << FoldHash(key, double_hash.mesto) << '\n';
    ODestroy(linear); ODestroy(double_hash);
}

void DemoOpenTables(){
    string line;
    cout << "Введите числовые ключи через пробел:\n";
    getline(cin, line);
    istringstream input(line);
    OpenTable linear, double_hash;
    OInit(linear, 11, false, true); OInit(double_hash, 11, true, true);
    string key, formula;
    int sum;
    while (input >> key){
        if (!FoldingSum(key, sum, formula)){ cout << "Пропущен неверный ключ: " << key << '\n'; continue; }
        OInsert(linear, key, to_string(sum));
        OInsert(double_hash, key, to_string(sum));
    }
    cout << "Линейное пробирование:\n"; OPrint(linear);
    cout << "Двойное хеширование:\n"; OPrint(double_hash);
    ODestroy(linear); ODestroy(double_hash);
}

void RunHashTask(){
    int choice;
    cout << "1 — метод свёртки, 2 — сравнение таблиц, 3 — показать коллизии: ";
    if (!(cin >> choice)) return;
    string rest; getline(cin, rest);
    if (choice == 1) DemoFolding();
    else if (choice == 2){ BenchmarkOpenHash(1000); BenchmarkOpenHash(5000); BenchmarkOpenHash(10000); }
    else if (choice == 3) DemoOpenTables();
    else cout << "Нет такого режима\n";
}
