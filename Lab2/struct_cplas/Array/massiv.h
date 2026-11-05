#pragma once
#include <iostream>
#include <string>
using namespace std;

struct Array{
    string* data;
    int size;
    int mesto;
};

void MInitArray(Array& arr, int mesto);
void MResize(Array& arr);
void MAddEnd(Array& arr, string chislo);
void MAddIndex(Array& arr, string chislo, int index);
string MGet(const Array& arr, int index);
void MGetIndex(const Array& arr, int index);
void MDeleteIndex(Array& arr, int index);
void MReplace(Array& arr, string chislo, int index);
int MLength(const Array& arr);
void MPrint(const Array& arr);
void MPrintAddresses(const Array& arr);
void MClear(Array& arr);
