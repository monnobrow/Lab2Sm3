CXX = clang++
CXXFLAGS = -std=c++11 -O2 -Wall -Wextra
STRUCT_CPP = $(wildcard struct_cplas/*/*.cpp)
TASK_CPP = $(wildcard zadachi/*.cpp)
UTIL_CPP = $(wildcard utils/*.cpp)
HEADERS = $(wildcard struct_cplas/*/*.h) $(wildcard zadachi/*.h) $(wildcard utils/*.h)
TARGET = build/lab2

all: $(TARGET)

$(TARGET): main.cpp $(STRUCT_CPP) $(TASK_CPP) $(UTIL_CPP) $(HEADERS)
	@mkdir -p build
	$(CXX) $(CXXFLAGS) main.cpp $(TASK_CPP) $(STRUCT_CPP) $(UTIL_CPP) -o $(TARGET)

clean:
	rm -rf build

.PHONY: all clean
