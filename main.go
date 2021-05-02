package main

import (
	"fmt"
	"reflect"
)

// Sample Unstructured/Multilevel/Nested JSON
/*
	{
		"map1": {
					"key1": "val1",
					"key2": "val2",
					"key3": "val3"
				},
		"map2": {
					"key4": "val4",
					"map3": {
								"key5": "val5",
								"key6": "val6"
							},
					"map4": {
								"key7": "val7",
								"map5": {
											"key8": "val8"
										}
							}
				}
	}
*/

func main() {

	// Go map[string]interface{} representation of the sample JSON structure mentioned in the comments above
	// The map[string]interface{} is the same structure that can be ready from json.Unmarshal function
	// The interface{} in map[string]interface{} has not limited to levels. i.e., any number of sub-levels can exists.
	// Our aim with code snippet in keyValuePairs() is to handle JSON with any number of such sub-levels.
	// Handling doesn't preserve the JSON structure details, but just leaf key-value pairs.
	// Note: Same structure can be read from API. But to make the demonstration easier, static data is used.
	map0 := map[string]interface{}{
		"map1": map[string]string{"key1": "val1", "key2": "val2", "key3": "val3"},
		"map2": map[string]interface{}{
			"key4": "val4",
			"map3": map[string]string{"key5": "val5", "key6": "val6"},
			"map4": map[string]interface{}{
				"key7": "val7",
				"map5": map[string]string{"key8": "val8"},
			},
		},
	}

	// This should return all the key value pairs in JSON as a map.
	kvs := keyValuePairs(map0)

	fmt.Println(kvs)
}

// keyValuePairs takes in map[string]interface{} as interface{} and returns map[string]string
// In simpler words, it takes map of uncertain levels and returns the leaf key-value pairs.
func keyValuePairs(m interface{}) map[string]string {
	// Using goroutines, channels or receiver here will definitely add value to execution.
	// But the aim here is to keep the demonstration simple and closer to native code.
	kvs := make(map[string]string)

	if reflect.ValueOf(m).Kind() == reflect.Map {
		// When interface type is a map.
		// Eg.: map0, map1, map2, map3, map4, map5

		mp, ok := m.(map[string]interface{})
		if ok {
			// When map type is map[string]interface{}.
			// i.e., it may have more levels inside.
			// Eg.: map0, map2, map4

			for k, v := range mp {
				if reflect.ValueOf(v).Kind() == reflect.String {
					// When value type is string.
					// i.e., only leaf item.
					// Eg.: key1-8 (all)

					kvs[k] = v.(string)
				} else {
					// When value type is map[string]interface{}.
					// i.e., if may have more maps inside.
					// Eg.: map0, map2, map4

					// keyValuePairs calls itself recursively till leaf items
					for nk, nv := range keyValuePairs(v) {

						// Add key-values pairs of inner levels to outer
						kvs[nk] = nv
					}
				}
			}
		} else {
			// When map type is not map[string]interface{}, but just map[string]string.
			// i.e., has leaf items but no sub-levels.
			// Eg.: map1, map3, map5

			for k, v := range m.(map[string]string) {

				// Add key-values pairs of map to result
				kvs[k] = v
			}
		}
	}

	return kvs
}
