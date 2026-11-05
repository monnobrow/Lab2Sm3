#include "hesh_table.h"

void HInitTable(HashTable& table, int mesto){
    HDestroy(table);
    if (mesto < 1){
        mesto = 1;
    }
    table.data = new HNode*[mesto]{};
    table.mesto = mesto;
}

int HHash(const string& key, int mesto){
    unsigned int hash = 0;
    for (unsigned char c : key){
        hash = hash * 31 + c;
    }
    return hash % mesto;
}

void HResize(HashTable& table){
    int new_mesto = table.mesto * 2;
    HNode** new_data = new HNode*[new_mesto]{};
    for (int i = 0; i < table.mesto; ++i){
        HNode* node = table.data[i];
        while (node != nullptr){
            HNode* next = node->next;
            int index = HHash(node->key, new_mesto);
            node->next = new_data[index];
            new_data[index] = node;
            node = next;
        }
    }
    delete[] table.data;
    table.data = new_data;
    table.mesto = new_mesto;
}

// true — добавили новый ключ; false — обновили существующий.
bool HInsert(HashTable& table, const string& key, const string& value){
    if (table.data == nullptr){
        HInitTable(table, 8);
    }
    int index = HHash(key, table.mesto);
    HNode* node = table.data[index];
    while (node != nullptr){
        if (node->key == key){
            node->value = value;
            return false;
        }
        node = node->next;
    }
    table.data[index] = new HNode{key, value, table.data[index]};
    ++table.size;
    // Примерно 75%: считаем без умножения размеров и больших типов.
    if (table.size >= table.mesto - table.mesto / 4){
        HResize(table);
    }
    return true;
}

bool HGet(const HashTable& table, const string& key, string& value){
    if (table.data == nullptr){
        return false;
    }
    HNode* node = table.data[HHash(key, table.mesto)];
    while (node != nullptr){
        if (node->key == key){
            value = node->value;
            return true;
        }
        node = node->next;
    }
    return false;
}

bool HContains(const HashTable& table, const string& key){
    string value;
    return HGet(table, key, value);
}

bool HDelete(HashTable& table, const string& key){
    if (table.data == nullptr){
        return false;
    }
    int index = HHash(key, table.mesto);
    HNode* node = table.data[index];
    HNode* previous = nullptr;
    while (node != nullptr && node->key != key){
        previous = node;
        node = node->next;
    }
    if (node == nullptr){
        return false;
    }
    if (previous == nullptr){
        table.data[index] = node->next;
    } else {
        previous->next = node->next;
    }
    delete node;
    --table.size;
    return true;
}

int HLength(const HashTable& table){
    return table.size;
}

void HPrint(const HashTable& table){
    for (int i = 0; i < table.mesto; ++i){
        HNode* node = table.data[i];
        while (node != nullptr){
            cout << node->key << " -> " << node->value << '\n';
            node = node->next;
        }
    }
}

void HPrintAddresses(const HashTable& table){
    for (int i = 0; i < table.mesto; ++i){
        cout << '[' << i << ']';
        HNode* node = table.data[i];
        while (node != nullptr){
            cout << " -> " << node->key << ':' << node->value << " (" << node << ')';
            node = node->next;
        }
        cout << '\n';
    }
    cout << "size = " << table.size << " mesto = " << table.mesto << '\n';
}

void HClear(HashTable& table){
    for (int i = 0; i < table.mesto; ++i){
        HNode* node = table.data[i];
        while (node != nullptr){
            HNode* next = node->next;
            delete node;
            node = next;
        }
        table.data[i] = nullptr;
    }
    table.size = 0;
}

void HDestroy(HashTable& table){
    HClear(table);
    delete[] table.data;
    table = HashTable{};
}
