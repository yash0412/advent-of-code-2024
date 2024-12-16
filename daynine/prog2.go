package daynine

import (
	"adventofcode/utils"
	"log"
)

type FileDetails struct {
	FileSize int64
	FileID   int64
	Index    int64
}

func Solve2() {
	inputString := readInputFile("./daynine/input.txt")
	diskMap, idMap := createDiskMapAndIdMap(inputString)
	// clearOutputFile("output-init.txt")
	// printLayout("output-init.txt", diskMap)
	newInput := defragmentDisk2(diskMap, idMap)
	log.Println("Checksum: ", calculateChecksum(newInput))
	// clearOutputFile("output.txt")
	// printLayout("output.txt", newInput)
}

func createDiskMapAndIdMap(input string) ([]int64, []FileDetails) {
	fileId := int64(0)
	diskMap := make([]int64, 0)
	idSizeData := make([]FileDetails, len(input))
	for i := range input {
		char := string(input[i])
		charNum := int64(utils.StringToInt(char))
		idSizeData[i] = FileDetails{FileID: fileId, FileSize: charNum, Index: int64(len(diskMap))}
		if i%2 != 0 {
			for k := int64(0); k < charNum; k++ {
				diskMap = append(diskMap, -1)
			}
		} else {
			for k := int64(0); k < charNum; k++ {
				diskMap = append(diskMap, fileId)
			}
			fileId++
		}
	}
	return diskMap, idSizeData
}

func defragmentDisk2(input []int64, idSizeData []FileDetails) []int64 {
	movedFilesMap := map[int64]bool{}
	// clearOutputFile("steps.txt")
	for m := 0; ; m += 2 {
		i := int64(len(idSizeData) - m - 1)
		if i < 0 {
			break
		}
		fileId, fileSize, fileIndex := idSizeData[i].FileID, idSizeData[i].FileSize, idSizeData[i].Index
		if movedFilesMap[fileId] {
			continue
		}
		movedFilesMap[fileId] = true
		// fmt.Println(len(idSizeData), i)
		for j := int64(1); j < i; j += 2 {
			spaceSize := idSizeData[j].FileSize
			if spaceSize >= fileSize {
				spaceIndex := idSizeData[j].Index
				for k := int64(0); k < fileSize; k++ {
					input[spaceIndex+k], input[fileIndex+k] = input[fileIndex+k], input[spaceIndex+k]
				}
				idSizeData = removeIndexFromSlice(idSizeData, j)
				idSizeData = addElemetToSliceAtIndex(idSizeData, j, FileDetails{FileID: idSizeData[j-1].FileID,
					FileSize: 0, Index: findSumOfFileSizesTillIndex(idSizeData, j-1)})
				idSizeData = addElemetToSliceAtIndex(idSizeData, j+1, FileDetails{FileID: fileId,
					FileSize: fileSize, Index: findSumOfFileSizesTillIndex(idSizeData, j)})
				idSizeData = addElemetToSliceAtIndex(idSizeData, j+2, FileDetails{FileID: fileId,
					FileSize: spaceSize - fileSize, Index: findSumOfFileSizesTillIndex(idSizeData, j+1)})
				break
			}
		}
		// printLayout("steps.txt", input)
		// printLayout("steps.txt", getFileSizesFromFileDetails(idSizeData))
		// printLayout("steps.txt", getIndicesFromFileDetails(idSizeData))
		// printLayout("steps.txt", []int64{})
	}
	return input
}

func removeIndexFromSlice(idSizeData []FileDetails, index int64) []FileDetails {
	newInput := cloneList(idSizeData)
	return append(newInput[:index], newInput[index+1:]...)
}

func addElemetToSliceAtIndex(idSizeData []FileDetails, index int64, element FileDetails) []FileDetails {
	newInput := cloneList(idSizeData)
	return append(newInput[:index], append([]FileDetails{element}, newInput[index:]...)...)
}

func cloneList(list []FileDetails) []FileDetails {
	newList := make([]FileDetails, len(list))
	copy(newList, list)
	return newList
}

func getFileSizesFromFileDetails(idSizeData []FileDetails) []int64 {
	fileSizes := make([]int64, len(idSizeData))
	for i, v := range idSizeData {
		fileSizes[i] = v.FileSize
	}
	return fileSizes
}

func getIndicesFromFileDetails(idSizeData []FileDetails) []int64 {
	indices := make([]int64, len(idSizeData))
	for i, v := range idSizeData {
		indices[i] = v.Index
	}
	return indices
}

func findSumOfFileSizesTillIndex(idSizeData []FileDetails, index int64) int64 {
	sum := int64(0)
	for i := int64(0); i <= index; i++ {
		sum += idSizeData[i].FileSize
	}
	return sum
}
