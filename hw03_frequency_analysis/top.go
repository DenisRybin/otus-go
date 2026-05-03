package hw03frequencyanalysis

import (
	"slices"
	"sort"
	"strings"
)

func Top10(source string) []string {
	var result []string
	if len(source) == 0 { // если подали пустой слайс, то и возвращаем пустой
		return result
	}
	x := strings.Fields(source) // разбиваем на слова
	slices.Sort(x)
	m := make(map[string]int)
	var prev string
	count := 1
	for k, v := range x { // получаем словарь с ко-вом
		if k == 0 { // first pass
			prev = v
			continue
		}
		if v == prev {
			count++
		} else {
			m[prev] = count
			prev = v
			count = 1
		}
	}
	m[prev] = count

	type kv struct {
		Key   string
		Value int
	}

	ss := make([]kv, 0, len(m))
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.SliceStable(ss, func(i, j int) bool { // сортируем по ко-ву
		return ss[i].Value > ss[j].Value
	})

	var resultgroup []string
	count = 0
	var prevvalue int

	for _, pair := range ss { // осталось отсортировать внутри групп
		if count == 0 { // first pass
			prevvalue = pair.Value
			resultgroup = append(resultgroup, pair.Key)
			count++
		} else {
			if prevvalue != pair.Value {
				slices.Sort(resultgroup)
				result = append(result, resultgroup...) // сортируем группу и добавляем в результат
				resultgroup = nil
				prevvalue = pair.Value
			}
			resultgroup = append(resultgroup, pair.Key)
			count++
			if count > 10 {
				break
			}
		}
	}

	result = append(result, resultgroup...)
	sliceLen := (len(result))
	if sliceLen > 10 {
		return result[0:10] // исключаем 11-ю, не отсортированную в группе
	} else {
		return result[0:sliceLen]
	}
}
