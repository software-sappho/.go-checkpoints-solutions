package main

func PrintMemory(arr [10]byte) {
	lastNonZero := getLastNonZeroIndex(arr)

	for i := 0; i < 10; i += 4 {
		for j := 0; j < 4 && i+j < 10; j++ {
			if j > 0 {
				print(" ")
			}
			if i+j <= lastNonZero || i == lastNonZero/4*4+4 {
				hex := "0123456789abcdef"
				print(string(hex[arr[i+j]>>4]))
				print(string(hex[arr[i+j]&0xf]))
			}
		}
		if i <= lastNonZero || i == lastNonZero/4*4+4 {
			print("\n")
		}
	}

	for i := 0; i < 10; i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			print(string(arr[i]))
		} else {
			print(".")
		}
	}
	print("\n")
}

func getLastNonZeroIndex(arr [10]byte) int {
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 0 {
			return i
		}
	}
	return -1
}

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}
