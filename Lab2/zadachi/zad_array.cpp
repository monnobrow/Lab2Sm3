#include "zad_array.h"
using namespace std;

void SortSubsetElements(Array& elements, int left, int right){
    if (left >= right) return;
    string middle = MGet(elements, (left + right) / 2);
    int i = left, j = right;
    while (i <= j){
        while (MGet(elements, i) < middle) ++i;
        while (MGet(elements, j) > middle) --j;
        if (i <= j){
            string value = elements.data[i];
            elements.data[i] = elements.data[j];
            elements.data[j] = value;
            ++i; --j;
        }
    }
    SortSubsetElements(elements, left, j);
    SortSubsetElements(elements, i, right);
}

void PrintSubset(const Array& current, bool& first){
    if (!first) cout << ", ";
    cout << '{';
    for (int i = 0; i < MLength(current); ++i){
        if (i > 0) cout << ", ";
        cout << MGet(current, i);
    }
    cout << '}';
    first = false;
}

void BuildSubsets(const Array& elements, Array& current, int start, int left, bool& first){
    if (left == 0){ PrintSubset(current, first); return; }
    for (int i = start; i <= MLength(elements) - left; ++i){
        if (i > start && MGet(elements, i) == MGet(elements, i - 1)) continue;
        MAddEnd(current, MGet(elements, i));
        BuildSubsets(elements, current, i + 1, left - 1, first);
        --current.size;
    }
}

void PrintAllSubsets(Array& elements){
    SortSubsetElements(elements, 0, MLength(elements) - 1);
    Array current;
    MInitArray(current, MLength(elements));
    bool first = true;
    cout << '[';
    for (int size = 0; size <= MLength(elements); ++size){
        BuildSubsets(elements, current, 0, size, first);
    }
    cout << "]\n";
    MClear(current);
}

void RunSubsets(){
    int n;
    cout << "Количество элементов массива: ";
    if (!(cin >> n) || n < 0){ cout << "Ошибка: неверный размер\n"; return; }
    Array elements;
    MInitArray(elements, n);
    cout << "Введите элементы через пробел:\n";
    for (int i = 0; i < n; ++i){
        string value;
        if (!(cin >> value)){
            cout << "Ошибка: элементов меньше, чем указано\n";
            MClear(elements);
            return;
        }
        MAddEnd(elements, value);
    }
    cout << "Подмножества: ";
    PrintAllSubsets(elements);
    MClear(elements);
}
