#include <iostream>
#include <string>
#include "massiv.h"
using namespace std;

void MInitArray(Array& arr, int mesto){
    if (mesto < 1){
        mesto = 1;
    }
    arr.data = new string[mesto];
    arr.size = 0;
    arr.mesto = mesto;
}

void MResize(Array& arr){
    string* new_data = new string[arr.mesto*2];
    for(int i=0; i<arr.size;++i){
        new_data[i] = arr.data[i];
    }
    delete[] arr.data;
    arr.data = new_data;
    arr.mesto = arr.mesto * 2;
}
void MAddEnd(Array& arr, string chislo){
    if (arr.size == arr.mesto){
        MResize(arr);
    }
    arr.data[arr.size] = chislo;
    arr.size++;// O(1)
}
// проверка индекса внутри структуры, результат печатается тут же
void MAddIndex(Array& arr, string chislo, int index){
    if (index < 0 || index > arr.size){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    if (arr.size == arr.mesto){
        MResize(arr);
    }
    for(int i = arr.size; i > index; --i){
        arr.data[i] = arr.data[i-1]; // чтобы не затереть элементы
    }
    arr.data[index] = chislo;
    arr.size++; // O(n)
    cout << "-> " << chislo << "\n";
}
string MGet(const Array& arr, int index){
    return arr.data[index]; // O(1), без проверки
}
void MGetIndex(const Array& arr, int index){
    if (index < 0 || index >= arr.size){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    cout << "-> " << arr.data[index] << "\n"; // O(1)
}
void MDeleteIndex(Array& arr, int index){
    if (index < 0 || index >= arr.size){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    cout << "-> " << arr.data[index] << "\n";
    for(int i = index; i < arr.size-1; ++i){
        arr.data[i] = arr.data[i+1];
    }
    arr.size--;// O(n)
}
void MReplace(Array& arr, string chislo, int index){
    if (index < 0 || index >= arr.size){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    arr.data[index] = chislo;
    cout << "-> " << chislo << "\n";
}// O(1)
int MLength(const Array& arr){
    return arr.size;
}// O(1)
void MPrint(const Array& arr){
    for(int i=0; i<arr.size;++i){
        cout << arr.data[i]<<" ";
    }
    cout << "\n";
}
void MPrintAddresses(const Array& arr){
    for(int i=0; i<arr.size;++i){
        cout << "["<< i << "] " << arr.data[i] << " | адрес: " << &arr.data[i] << '\n';
    }
    cout << "size = " << arr.size << " mesto = " << arr.mesto << "\n\n";
}
void MClear(Array& arr){
    delete[] arr.data;
    arr.data = nullptr;
    arr.size = 0;
    arr.mesto = 0;
}
