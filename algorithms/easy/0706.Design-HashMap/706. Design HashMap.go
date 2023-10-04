package main

type MyHashMap struct {
	data map[int]int
}

func Constructor() MyHashMap {
	return MyHashMap{data: make(map[int]int)}
}

func (hash *MyHashMap) Put(key int, value int) {
	hash.data[key] = value
}

func (hash *MyHashMap) Get(key int) int {
	if val, ok := hash.data[key]; ok {
		return val
	}
	return -1
}

func (hash *MyHashMap) Remove(key int) {
	delete(hash.data, key)
}

func main() {
	mm := Constructor()
	mm.Put(1, 10)
	mm.Get(1)
	mm.Remove(1)

}
