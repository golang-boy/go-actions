package reflect

import "reflect"

func IterateArrayOrSlice(entity any) ([]any, error) {
	val := reflect.ValueOf(entity)

	res := make([]any, 0, val.Len())
	for i := 0; i < val.Len(); i++ {
		e := val.Index(i) // 通过索引遍历
		res = append(res, e.Interface())
	}
	return res, nil
}

func IterateMap(entity any) ([]any, []any, error) {
	val := reflect.ValueOf(entity)

	keys := make([]any, 0, val.Len())
	vals := make([]any, 0, val.Len())

	// 方式一：
	// for _, key := range val.MapKeys() {
	// 	keys = append(keys, key.Interface())
	// 	vals = append(vals, val.MapIndex(key).Interface())
	// }

	// 方式二：
	itr := val.MapRange()
	for itr.Next() {
		keys = append(keys, itr.Key().Interface())
		vals = append(vals, itr.Value().Interface())
	}

	return keys, vals, nil
}
