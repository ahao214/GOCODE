package main

//278. 第一个错误的版本
import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	return sort.Search(n, func(version int) bool { return isBadVersion(version) })
}
