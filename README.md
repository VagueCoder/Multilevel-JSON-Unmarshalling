# Multilevel-JSON-Unmarshalling
The Multilevel (or) Unstructured (or) Nested JSON is nothing but any JSON whose structure (i.e., number of levels) is either variable or not known till the time of execution.

Data of this kind is generally read from an API which has the varying data structure but when the only useful part of that is key-values.

## What does this do?
The simple 10 lined (approx.) function in the code takes in the JSON as object and returns slice of all the leaf key-value pairs in the structure. This doesn't preserve details of the depth, but just KVs of all levels.

For instance, if your API was fed with the following JSON structure
```
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
```

The expected output from the function would be
```
map[key1:val1 key2:val2 key3:val3 key4:val4 key5:val5 key6:val6 key7:val7 key8:val8]
```

#### As simple as that!!

## My thoughts on this
1. This is a one-time development and isn't made for adding better versions of the same. Hope this gets captured in search for Newbies to go (golang) for their understanding.
1. This should handle JSON with any depth. The only expected limitation is the Out-Of-Memory-Exception.
1. Using goroutines, channels or receiver to function will definitely add value to execution. But the aim is to keep it simple for demonstration/execution and closer to native code.

When this is useful or enlightens someone, I'd be glad. But if you want to add somemore interesting features to this, use it. This is all upto you buddy!

Reach-out to me at: VagueCoder0t0.n@gmail.com

## Happy Coding !! :metal: