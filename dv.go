package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"math"
)

type table map[string]map[string]int
type iteration map[string]table

const inf = math.MaxInt

func (i iteration) Hash() string {
	b, _ := json.Marshal(i)
	return fmt.Sprintf("%x", md5.Sum(b))
}

func (i iteration) Print() {
	for k, v := range i {
		println("Table:", k)
		for row, cols := range v {
			fmt.Printf("%s: ", row)
			for col, cost := range cols {
				if cost == math.MaxInt {
					fmt.Printf("(%s:∞) ", col)
				} else {
					fmt.Printf("(%s:%d) ", col, cost)
				}
			}
			println()
		}
	}
}

func (src iteration) deepCopyIter() iteration {
	copy := make(iteration)
	for k, table1 := range src {
		newTable := make(table)
		for row, cols := range table1 {
			newRow := make(map[string]int)
			for col, cost := range cols {
				newRow[col] = cost
			}
			newTable[row] = newRow
		}
		copy[k] = newTable
	}
	return copy
}

func minFromMap(m map[string]int) int {
	min := math.MaxInt
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	return min
}

func Calculate(iter0 iteration) {
	iters := []iteration{iter0.deepCopyIter()}
	iterCount := 1
	for {
		println("Calculating next iteration:", iterCount)
		lastIter := iters[len(iters)-1]
		newIter := lastIter.deepCopyIter()
		changed := false
		for tableName, table := range newIter {
			for row, cols := range table {
				for col, currentCost := range cols {
					//direkte nachbarn oder sich selbst skippen
					if row == col || row == tableName {
						continue
					}
					//nach row über col
					costToNextTable := table[col][col]
					costFromNextTableToDest := minFromMap(lastIter[col][row])
					//wir haben noch keinen weg zum "via" table oder wir kommen von da aus noch nicht weiter
					if costToNextTable == inf || costFromNextTableToDest == inf {
						continue
					}
					sum := costToNextTable + costFromNextTableToDest
					if sum < currentCost {
						newIter[tableName][row][col] = sum
						changed = true
					}
				}
			}
		}
		iters = append(iters, newIter)
		if !changed {
			break
		}
		iterCount++
	}
	for i, v := range iters {
		println("--------------")
		println("Iteration", i)
		v.Print()
		println("--------------")
	}
}
