#include "zad_mnozh.h"
#include "../utils/input.h"
#include <climits>
#include <sstream>
using namespace std;

bool ReadPartitionSet(string text, Set& set){
    for (char& c : text){
        if (c == '{' || c == '}' || c == ',') c = ' ';
    }
    istringstream input(text);
    string item;
    while (input >> item){
        int value;
        if (!ReadPositiveInt(item, value)) return false;
        SetAdd(set, to_string(value));
    }
    return SetSize(set) > 0;
}

void SortPartition(Array& elements, int left, int right){
    if (left >= right) return;
    int middle = stoi(MGet(elements, (left + right) / 2));
    int i = left, j = right;
    while (i <= j){
        while (stoi(MGet(elements, i)) > middle) ++i;
        while (stoi(MGet(elements, j)) < middle) --j;
        if (i <= j){
            string value = elements.data[i];
            elements.data[i] = elements.data[j];
            elements.data[j] = value;
            ++i; --j;
        }
    }
    SortPartition(elements, left, j);
    SortPartition(elements, i, right);
}

bool FillEqualGroups(const Array& elements, int index, int target,
                     int group_count, int* sums, int* groups){
    if (index == MLength(elements)) return true;
    int value = stoi(MGet(elements, index));
    int previous_sum = -1;
    for (int group = 0; group < group_count; ++group){
        if (sums[group] == previous_sum) continue;
        previous_sum = sums[group];
        if (value > target - sums[group]) continue;
        sums[group] += value;
        groups[index] = group;
        if (FillEqualGroups(elements, index + 1, target, group_count, sums, groups)){
            return true;
        }
        sums[group] -= value;
        groups[index] = -1;
        if (sums[group] == 0) break;
    }
    return false;
}

bool EqualPartition(Array& elements, int target, int*& groups, int& group_count){
    groups = nullptr;
    group_count = 0;
    if (target < 1 || MLength(elements) == 0) return false;
    int total = 0;
    for (int i = 0; i < MLength(elements); ++i){
        int value = stoi(MGet(elements, i));
        if (value > target || total > INT_MAX - value) return false;
        total += value;
    }
    if (total % target != 0) return false;
    group_count = total / target;
    if (group_count > MLength(elements)) return false;
    SortPartition(elements, 0, MLength(elements) - 1);
    groups = new int[MLength(elements)];
    for (int i = 0; i < MLength(elements); ++i) groups[i] = -1;
    int* sums = new int[group_count]{};
    bool found = FillEqualGroups(elements, 0, target, group_count, sums, groups);
    delete[] sums;
    if (!found){ delete[] groups; groups = nullptr; }
    return found;
}

void RunPartition(){
    cout << "Введите множество, например {4,10,5,1,3,7}:\n";
    string line;
    getline(cin, line);
    Set set;
    SetInit(set);
    if (!ReadPartitionSet(line, set)){
        cout << "Ошибка: нужно непустое множество натуральных чисел\n";
        SetDestroy(set);
        return;
    }
    int target;
    cout << "Сумма каждого подмножества: ";
    if (!(cin >> target) || target < 1){
        cout << "Ошибка: нужна положительная сумма\n";
        SetDestroy(set);
        return;
    }
    Array elements;
    MInitArray(elements, SetSize(set));
    SetElements(set, elements);
    int* groups = nullptr;
    int group_count;
    if (EqualPartition(elements, target, groups, group_count)){
        for (int group = 0; group < group_count; ++group){
            cout << "Подмножество " << group + 1 << ": {";
            bool first = true;
            for (int i = 0; i < MLength(elements); ++i){
                if (groups[i] == group){
                    if (!first) cout << ", ";
                    cout << MGet(elements, i);
                    first = false;
                }
            }
            cout << "}, сумма: " << target << '\n';
        }
    } else {
        cout << "Ошибка: множество нельзя разбить на подмножества с заданной суммой\n";
    }
    delete[] groups;
    MClear(elements);
    SetDestroy(set);
}
