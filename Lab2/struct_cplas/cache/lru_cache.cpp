#include "lru_cache.h"

LRUCache::LRUCache(int cap){
    capacity = cap < 0 ? 0 : cap;
    size = 0;
    HInitTable(index, 8);
    LInitList(keys);
    entries = capacity > 0 ? new LRUEntry[capacity] : nullptr;
}

LRUCache::~LRUCache(){
    LClear(keys);
    HDestroy(index);
    delete[] entries;
}

void TouchLRU(LRUCache& cache, int slot){
    string key = cache.entries[slot].node->chislo;
    LDeleteNode(cache.keys, cache.entries[slot].node);
    LAddEnd(cache.keys, key);
    cache.entries[slot].node = cache.keys.tail;
}

string LRUCache::Get(const string& key){
    string saved;
    if (!HGet(index, key, saved)) return "-1";
    int slot = stoi(saved);
    TouchLRU(*this, slot);
    return entries[slot].value;
}

void LRUCache::Set(const string& key, const string& value){
    if (capacity == 0) return;
    string saved;
    if (HGet(index, key, saved)){
        int slot = stoi(saved);
        entries[slot].value = value;
        TouchLRU(*this, slot);
        return;
    }
    int slot;
    if (size == capacity){
        string victim = keys.head->chislo;
        HGet(index, victim, saved);
        slot = stoi(saved);
        LDeleteFirst(keys);
        HDelete(index, victim);
    } else {
        slot = size;
        ++size;
    }
    LAddEnd(keys, key);
    entries[slot].value = value;
    entries[slot].node = keys.tail;
    HInsert(index, key, to_string(slot));
}
