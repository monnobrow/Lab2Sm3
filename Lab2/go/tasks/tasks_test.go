package tasks

import (
	"bytes"
	"fmt"
	"lab2ann/structures"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func runTask(task, input string) (int, string, string) {
	var out, err bytes.Buffer
	status := Run([]string{task}, strings.NewReader(input), &out, &err)
	return status, out.String(), err.String()
}
func TestBaseball(t *testing.T) {
	for _, c := range []struct {
		ops   string
		score int
		ok    bool
	}{{"[5,2,C,D,+]", 30, true}, {"[5,4,D,+,C,3]", 20, true}, {"", 0, true}, {"[1,C]", 0, true}, {"+", 0, false}, {"D", 0, false}, {"C", 0, false}, {"0", 0, false}, {"-1", 0, false}, {"hello", 0, false}, {"2147483648", 0, false}, {"1073741824 D", 0, false}, {"2147483647 1 +", 0, false}, {"2147483647 1", 0, false}} {
		score, ok := BaseballScore(c.ops)
		if ok != c.ok || (ok && score != c.score) {
			t.Fatal(c, score, ok)
		}
	}
	for _, s := range []string{"", "-1", "+1", "1.0", "2147483648", "0"} {
		if _, ok := PositiveInt(s); ok {
			t.Fatal("invalid positive", s)
		}
	}
	if n, ok := PositiveInt("00019"); !ok || n != 19 {
		t.Fatal("normalization")
	}
}
func partitionOracle(values []int, target int) bool {
	total := 0
	for _, v := range values {
		total += v
	}
	if total%target != 0 {
		return false
	}
	full := (1 << len(values)) - 1
	var visit func(int) bool
	visit = func(remaining int) bool {
		if remaining == 0 {
			return true
		}
		first := remaining & -remaining
		for subset := remaining; subset > 0; subset = (subset - 1) & remaining {
			if subset&first == 0 {
				continue
			}
			sum := 0
			for i, v := range values {
				if subset&(1<<i) != 0 {
					sum += v
				}
			}
			if sum == target && visit(remaining^subset) {
				return true
			}
		}
		return false
	}
	return visit(full)
}
func TestPartition(t *testing.T) {
	random := rand.New(rand.NewSource(19))
	for trial := 0; trial < 300; trial++ {
		set := structures.NewSet()
		values := []int{}
		for n := random.Intn(8) + 1; len(values) < n; {
			value := random.Intn(12) + 1
			if set.Add(strconv.Itoa(value)) {
				values = append(values, value)
			}
		}
		target := random.Intn(25) + 1
		elements := set.Elements()
		groups, count, ok := EqualPartition(elements, target)
		if ok != partitionOracle(values, target) {
			t.Fatal("partition differs", values, target, ok)
		}
		if ok {
			sums := make([]int, count)
			for i, g := range groups {
				if g < 0 || g >= count {
					t.Fatal("invalid group")
				}
				sums[g] += number(elements, i)
			}
			for _, sum := range sums {
				if sum != target {
					t.Fatal("bad sum", sum, target)
				}
			}
		}
	}
	set := structures.NewSet()
	if !ReadPartitionSet("{01,1,2}", set) || set.Length() != 2 {
		t.Fatal("set normalization")
	}
	for _, s := range []string{"{}", "{-1,2}", "{0}", "{abc}"} {
		if ReadPartitionSet(s, structures.NewSet()) {
			t.Fatal("bad set", s)
		}
	}
	empty := structures.NewArray(1)
	if _, _, ok := EqualPartition(empty, 0); ok {
		t.Fatal("empty partition")
	}
	large := structures.NewArray(2)
	large.AddEnd("2147483647")
	large.AddEnd("1")
	if _, _, ok := EqualPartition(large, 2147483647); ok {
		t.Fatal("total overflow")
	}
}
func TestSubsets(t *testing.T) {
	a := structures.NewArray(3)
	for _, v := range []string{"z", "x", "y"} {
		a.AddEnd(v)
	}
	var out bytes.Buffer
	PrintAllSubsets(&out, a)
	want := "[{}, {x}, {y}, {z}, {x, y}, {x, z}, {y, z}, {x, y, z}]\n"
	if out.String() != want {
		t.Fatal(out.String())
	}
	a = structures.NewArray(3)
	for _, v := range []string{"a", "a", "b"} {
		a.AddEnd(v)
	}
	out.Reset()
	PrintAllSubsets(&out, a)
	if out.String() != "[{}, {a}, {b}, {a, a}, {a, b}, {a, a, b}]\n" {
		t.Fatal("duplicate subsets", out.String())
	}
	random := rand.New(rand.NewSource(19))
	for trial := 0; trial < 100; trial++ {
		n := random.Intn(8)
		values := make([]string, n)
		a = structures.NewArray(n)
		for i := range values {
			values[i] = strconv.Itoa(random.Intn(3))
			a.AddEnd(values[i])
		}
		model := map[string]bool{}
		for mask := 0; mask < (1 << n); mask++ {
			subset := []string{}
			for i, v := range values {
				if mask&(1<<i) != 0 {
					subset = append(subset, v)
				}
			}
			sort.Strings(subset)
			model["{"+strings.Join(subset, ", ")+"}"] = true
		}
		out.Reset()
		PrintAllSubsets(&out, a)
		got := strings.TrimSuffix(strings.TrimPrefix(out.String(), "["), "]\n")
		// Boundaries are unambiguous because test elements contain only digits.
		chunks := strings.Split(got, "}, {")
		if len(chunks) != len(model) {
			t.Fatal("subset count", values, len(chunks), len(model))
		}
		for i, chunk := range chunks {
			if i > 0 {
				chunk = "{" + chunk
			}
			if i < len(chunks)-1 {
				chunk += "}"
			}
			if !model[chunk] {
				t.Fatal("unexpected subset", chunk)
			}
			delete(model, chunk)
		}
	}
}
func TestSetQueries(t *testing.T) {
	path := filepath.Join(t.TempDir(), "set.txt")
	for _, c := range []struct {
		q, want string
		status  int
	}{{"SETADD 5", "-> 5\n", 0}, {"SETADD 5", "5 уже есть в множестве\n", 0}, {"SET_AT 5", "-> TRUE\n", 0}, {"SETDEL 5", "-> 5 удалён\n", 0}, {"SET_AT 5", "-> FALSE\n", 0}, {"SETDEL 5", "5 нет в множестве\n", 0}, {"SETPRINT", "{}\n", 0}} {
		var out, err bytes.Buffer
		status := Run([]string{"--file", path, "--query", c.q}, strings.NewReader(""), &out, &err)
		if status != c.status || out.String() != c.want || err.Len() != 0 {
			t.Fatal(c, status, out.String(), err.String())
		}
	}
	for _, args := range [][]string{{"--file"}, {"unknown"}, {"--file", path}, {"--file", path, "--query", ""}, {"--file", path, "--query", "WRONG"}, {"--file", path, "--query", "SETADD"}, {"--file", path, "--query", "SETPRINT extra"}, {"--file", path, "--query", "SETADD a b"}, {"--file", filepath.Join(path, "missing"), "--query", "SETADD a"}} {
		var out, err bytes.Buffer
		if status := Run(args, strings.NewReader(""), &out, &err); status != 1 || err.Len() == 0 {
			t.Fatal("bad arguments accepted", args, status)
		}
	}
	if err := os.WriteFile(path, []byte("a\r\nb\n\na\n"), 0600); err != nil {
		t.Fatal(err)
	}
	set := structures.NewSet()
	if err := LoadSet(set, path); err != nil || set.Length() != 2 {
		t.Fatal("load CRLF", err)
	}
	if LoadSet(structures.NewSet(), t.TempDir()) == nil {
		t.Fatal("directory should not be read as a set")
	}
}
func TestTaskErrorsAndMenu(t *testing.T) {
	for _, c := range []struct{ task, input, part string }{{"1", "+\n", "Ошибка:"}, {"3", "{}\n", "Ошибка:"}, {"3", "{1,2}\n0\n", "Ошибка:"}, {"3", "{1,2}\nx\n", "Ошибка:"}, {"4", "-1\n", "Ошибка:"}, {"4", "2\na\n", "Ошибка:"}, {"5", "\n", "Строка пустая"}, {"5", "aaaa\n1\n", "Ошибка:"}, {"5", "aaaa\n00\n", "Результат: aa"}, {"6", "1\nbad\n", "Ошибка:"}, {"6", "3\n11 22 bad 33\n", "Пропущен неверный ключ: bad"}, {"6", "9\n", "Нет такого режима"}, {"7", "-1 1\n", "Ошибка:"}, {"7", "2 -1\n", "Ошибка:"}, {"7", "x\n", "Ошибка:"}, {"7", "2 1\nSET a\n", "Ошибка:"}, {"7", "2 1\nGET\n", "Ошибка:"}, {"7", "2 1\nWRONG a\n", "Ошибка:"}} {
		_, out, _ := runTask(c.task, c.input)
		if !strings.Contains(out, c.part) {
			t.Fatal(c, out)
		}
	}
	path := filepath.Join(t.TempDir(), "set.txt")
	input := fmt.Sprintf("1\n[5,2,C,D,+]\n2\n%s\nSETADD 19\n3\n{4,10,5,1,3,7}\n15\n4\n3\nx y z\n5\nabracadabra\n\n6\n1\n523456795\n7\n2 2\nSET 1 2 GET 1\n0\n", path)
	var out, err bytes.Buffer
	if status := Run(nil, strings.NewReader(input), &out, &err); status != 0 || err.Len() != 0 {
		t.Fatal(status, err.String())
	}
	for _, part := range []string{"Сумма очков: 30", "-> 19", "Подмножество 1: {10, 5}", "{x, y, z}", "Расшифровка: abracadabra", "523+456+795 = 1774", "Результат: 2"} {
		if !strings.Contains(out.String(), part) {
			t.Fatal("menu missed task", part, out.String())
		}
	}
	out.Reset()
	Run([]string{"--help"}, strings.NewReader(""), &out, &err)
	if !strings.Contains(out.String(), "lab2 1..7") {
		t.Fatal("help")
	}
	a := NewApp(strings.NewReader(""), &out, &err)
	if a.RunTask(9) != 1 {
		t.Fatal("invalid task")
	}
	_, output, _ := runTask("6", "2\n")
	for _, n := range []int{1000, 5000, 10000} {
		if !strings.Contains(output, fmt.Sprintf("N = %d", n)) {
			t.Fatal("benchmark missing N")
		}
	}
	if strings.Contains(output, "Ошибка:") {
		t.Fatal("benchmark consistency", output)
	}
}
