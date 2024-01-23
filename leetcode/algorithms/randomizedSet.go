package algorithms

import "math/rand"

// map stores values and indexes
// nums stores values
type RandomizedSet struct {
	numsMap map[int]int
	nums    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		numsMap: make(map[int]int),
		nums:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, found := this.numsMap[val]; found {
		return false
	}
	this.numsMap[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, found := this.numsMap[val]
	if !found {
		return false
	}
	last := this.nums[len(this.nums)-1]
	this.nums[index] = last
	this.numsMap[last] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.numsMap, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.nums))
	return this.nums[index]
}
