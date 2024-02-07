package functional

import "fmt"

func SortByName() {
	Usr := getData("users.txt")
	sort(Usr, 0, len(Usr)-1, true)
	for i := 0; i < len(Usr); i++ {
		fmt.Println(Usr[i])
	}
}

func SortByData() {
	Usr := getData("users.txt")
	sort(Usr, 0, len(Usr)-1, false)
	for i := 0; i < len(Usr); i++ {
		fmt.Println(Usr[i])
	}
}

func sort(arr Users, low int, high int, name bool) {
	if name {
		if low < high {
			index := partitionName(arr, low, high)
			sort(arr, low, index-1, name)
			sort(arr, index+1, high, name)
		}
	} else {
		if low < high {
			index := partitionData(arr, low, high)
			sort(arr, low, index-1, name)
			sort(arr, index+1, high, name)
		}
	}

}

func partitionName(arr Users, low int, high int) int {
	i := low - 1
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j].Name < pivot.Name {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		} else if arr[j].Name == pivot.Name {
			if arr[j].Surname < pivot.Surname {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}

func partitionData(arr Users, low int, high int) int {
	i := low - 1
	pivot := arr[high]
	for j := low; j < high; j++ {
		if arr[j].Year < pivot.Year {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		} else if arr[j].Year == pivot.Year {
			if arr[j].Month < pivot.Month {
				i++
				arr[i], arr[j] = arr[j], arr[i]
			} else if arr[j].Month == pivot.Month {
				if arr[j].Day < pivot.Day {
					i++
					arr[i], arr[j] = arr[j], arr[i]
				}
			}
		}
	}
	arr[high], arr[i+1] = arr[i+1], arr[high]
	return i + 1
}

