// Copyright (c) 2019.
// Author: Quan TRAN

package utils

func Map(list []string, f func(string) string) []string {
	result := make([]string, len(list))
	for i, item := range list {
		result[i] = f(item)
	}
	return result
}

func Unique(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
