#pragma once
#include "../hash/hesh_table.h"
#include "../spisok/spisok2.h"

struct LRUEntry{
    string value;
    LNode* node = nullptr;
};

struct LRUCache{
    int capacity;
    int size;
    HashTable index;
    LList keys;
    LRUEntry* entries;

    LRUCache(int cap);
    ~LRUCache();
    string Get(const string& key);
    void Set(const string& key, const string& value);
};
