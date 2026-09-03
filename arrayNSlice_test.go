package gobasic

import (
	"log"
	"testing"
	"unsafe"
)

func TestArrayNSlice(t *testing.T) {
	arrayNSlice()
}

func arrayNSlice() {
	//ARRAY
	arrNums := [6]int{}
	arrNums[0] = 1
	arrNums[1] = 2
	arrNums[2] = 3
	arrNums[3] = 4
	arrNums[4] = 5
	arrNums[5] = 6

	log.Printf("\nArray: %v", arrNums)

	// Memotong array
	cutArr1 := arrNums[:4] // dari awal sampai index ke-2
	log.Printf("\nPanjanga array %d, isi array adalah %+v", len(cutArr1), cutArr1)

	cutArr2 := arrNums[1:5] // dari index ke-1 sampai index ke-3
	log.Printf("\nPanjanga array %d, isi array adalah %v", len(cutArr2), cutArr2)

	cutArr3 := arrNums[1:] // dari index ke-2 sampai akhir array
	log.Printf("\nPanjanga array %d, isi array adalah %v", len(cutArr3), cutArr3)

	//SLICE

	var nums []int = arrNums[:5]
	log.Printf("\nSlice nums yang mereferensikan ke array arrNums %v", nums)

	nums[0] = 7
	log.Printf("\narray index ke-0 di variable arrNums akan berubah juga ke nilai baru \n arrNums[0] = %d \n arrNums = %v \n nums = %v ", arrNums[0], arrNums, nums)

	newSlice := make([]int, 0, 8)
	log.Printf("isi newSlice saat ini %v panjanganya %d, kapasitasnya %d,pointer slice %v ", newSlice, len(newSlice), cap(newSlice), unsafe.Pointer(&newSlice))

	newSlice = append(newSlice, 2, 3, 5, 61, 5, 2, 5)
	log.Printf("isi newSlice setelah di tambahkan beberapa element %v, panjanganya %d, kapasitasnya %d ,pointer %v, pointer element index ke-0 %v", newSlice, len(newSlice), cap(newSlice), unsafe.Pointer(&newSlice), unsafe.Pointer(&newSlice[0]))

	newSlice = append(newSlice, 5, 6)
	log.Printf("isi newSlice setelah di tambahkan beberapa element lagi yang melebihi kapasitas sehingga terjadi perubahan alamat memory dari element index ke-0 dan kapasitas dari slice %v, panjanganya %d, kapasitasnya %d pointer %v, pointer element index ke-0 %v", newSlice, len(newSlice), cap(newSlice), unsafe.Pointer(&newSlice), unsafe.Pointer(&newSlice[0]))

}
