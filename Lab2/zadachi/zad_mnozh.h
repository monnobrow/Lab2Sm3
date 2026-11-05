#pragma once
#include "../struct_cplas/Mnozh/mnozh.h"

bool ReadPartitionSet(std::string text, Set& set);
bool EqualPartition(Array& elements, int target, int*& groups, int& group_count);
void RunPartition();
