package main

import "sort"

//911. 在线选举
type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons, times []int) TopVotedCandidate {
	tops := []int{}
	top := -1
	voteCounts := map[int]int{-1: -1}
	for _, p := range persons {
		voteCounts[p]++
		if voteCounts[p] >= voteCounts[top] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{tops, times}
}

func (c *TopVotedCandidate) Q(t int) int {
	return c.tops[sort.SearchInts(c.times, t+1)-1]
}

// TopVotedCandidate define
type TopVotedCandidate2 struct {
	persons []int
	times   []int
}

func Constructor911(persons []int, times []int) TopVotedCandidate2 {
	leaders, votes := make([]int, len(persons)), make([]int, len(persons))
	leader := persons[0]
	for i := 0; i < len(persons); i++ {
		p := persons[i]
		votes[p]++
		if votes[p] >= votes[leader] {
			leader = p
		}
		leaders[i] = leader
	}
	return TopVotedCandidate2{persons: leaders, times: times}
}

// Q define
func (tvc *TopVotedCandidate2) Q2(t int) int {
	i := sort.Search(len(tvc.times), func(p int) bool { return tvc.times[p] > t })
	return tvc.persons[i-1]
}
