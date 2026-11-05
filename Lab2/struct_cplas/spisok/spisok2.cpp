#include <iostream>
#include <string>
#include "spisok2.h"
using namespace std;

void LInitList(LList& list){
    list.head = nullptr;
    list.tail = nullptr;
}

LNode* LSearch(LNode* flist, string chislo){
    while (flist != nullptr && flist->chislo != chislo){
        flist = flist->next;
    }
    return flist;
}
// вставка рядом с узлом и удаление узла - O(1), есть prev
void LLinkBefore(LList& list, LNode* flist, string chislo){
    LNode* new_node = new LNode{chislo, flist, flist->prev};
    if (flist->prev != nullptr){
        flist->prev->next = new_node;
    } else {
        list.head = new_node;
    }
    flist->prev = new_node;
}
void LLinkAfter(LList& list, LNode* flist, string chislo){
    LNode* new_node = new LNode{chislo, flist->next, flist};
    if (flist->next != nullptr){
        flist->next->prev = new_node;
    } else {
        list.tail = new_node;
    }
    flist->next = new_node;
}
void LDeleteNode(LList& list, LNode* flist){
    if (flist->prev != nullptr){
        flist->prev->next = flist->next;
    } else {
        list.head = flist->next;
    }
    if (flist->next != nullptr){
        flist->next->prev = flist->prev;
    } else {
        list.tail = flist->prev;
    }
    delete flist;
}

void LAddFirst(LList& list, string chislo){
    LNode* new_node = new LNode{chislo, list.head, nullptr};
    if (list.head != nullptr){
        list.head->prev = new_node;
    } else {
        list.tail = new_node;
    }
    list.head = new_node;
}
void LAddEnd(LList& list, string chislo){
    LNode* new_node = new LNode{chislo, nullptr, list.tail};
    if (list.tail != nullptr){
        list.tail->next = new_node;
    } else {
        list.head = new_node;
    }
    list.tail = new_node;
}
void LAddAfter(LList& list, string target, string chislo){
    LNode* cur_node = LSearch(list.head, target);
    if (cur_node == nullptr){
        cout << "Ошибка: " << target << " не найден\n";
        return;
    }
    LLinkAfter(list, cur_node, chislo);
    cout << "-> " << chislo << "\n";
}
void LAddBefore(LList& list, string target, string chislo){
    LNode* cur_node = LSearch(list.head, target);
    if (cur_node == nullptr){
        cout << "Ошибка: " << target << " не найден\n";
        return;
    }
    LLinkBefore(list, cur_node, chislo);
    cout << "-> " << chislo << "\n";
}
// index == длина - вставка в конец
void LAddIndex(LList& list, string chislo, int index){
    LNode* cur_node = list.head;
    int i = 0;
    while (i < index && cur_node != nullptr){
        cur_node = cur_node->next;
        i++;
    }
    if (index < 0 || i < index){ // список короче индекса
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    if (cur_node == nullptr){
        LAddEnd(list, chislo);
    } else {
        LLinkBefore(list, cur_node, chislo);
    }
    cout << "-> " << chislo << "\n";
}

void LDeleteFirst(LList& list){
    if (list.head == nullptr){
        cout << "Ошибка: список пуст\n";
        return;
    }
    LDeleteNode(list, list.head);
}
void LDeleteEnd(LList& list){
    if (list.tail == nullptr){
        cout << "Ошибка: список пуст\n";
        return;
    }
    LDeleteNode(list, list.tail);
}
void LDeleteAfter(LList& list, string target){
    LNode* cur_node = LSearch(list.head, target);
    if (cur_node == nullptr || cur_node->next == nullptr){
        cout << "Ошибка: после " << target << " нечего удалять\n";
        return;
    }
    LDeleteNode(list, cur_node->next);
}
void LDeleteBefore(LList& list, string target){
    LNode* cur_node = LSearch(list.head, target);
    if (cur_node == nullptr || cur_node->prev == nullptr){
        cout << "Ошибка: перед " << target << " нечего удалять\n";
        return;
    }
    LDeleteNode(list, cur_node->prev);
}
void LDeleteIndex(LList& list, int index){
    LNode* cur_node = list.head;
    for (int i = 0; i < index && cur_node != nullptr; ++i){
        cur_node = cur_node->next;
    }
    if (index < 0 || cur_node == nullptr){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    LDeleteNode(list, cur_node);
}
void LDeleteValue(LList& list, string chislo){
    LNode* cur_node = LSearch(list.head, chislo);
    if (cur_node == nullptr){
        cout << "Ошибка: " << chislo << " не найден\n";
        return;
    }
    LDeleteNode(list, cur_node);
}

void LGet(const LList& list, int index){
    LNode* cur_node = list.head;
    for (int i = 0; i < index && cur_node != nullptr; ++i){
        cur_node = cur_node->next;
    }
    if (index < 0 || cur_node == nullptr){
        cout << "Ошибка: нет такого индекса\n";
        return;
    }
    cout << "-> " << cur_node->chislo << "\n";
}
void LFind(const LList& list, string chislo){
    if (LSearch(list.head, chislo) != nullptr){
        cout << "-> TRUE\n";
    } else {
        cout << "-> FALSE\n";
    }
}

void LPrint(const LList& list){
    for (LNode* cur_node = list.head; cur_node != nullptr; cur_node = cur_node->next){
        cout << cur_node->chislo << "\n";
    }
}
void LPrintBack(const LList& list){
    for (LNode* cur_node = list.tail; cur_node != nullptr; cur_node = cur_node->prev){
        cout << cur_node->chislo << "\n";
    }
}
void LPrintAddresses(const LList& list){
    for (LNode* cur_node = list.head; cur_node != nullptr; cur_node = cur_node->next){
        cout << cur_node->chislo << " | адрес: " << cur_node << " | next: " << cur_node->next << " | prev: " << cur_node->prev << '\n';
    }
    cout << '\n';
}
void LClear(LList& list){
    while (list.head != nullptr){
        LNode* delete_node = list.head;
        list.head = list.head->next;
        delete delete_node;
    }
    list.tail = nullptr;
}
