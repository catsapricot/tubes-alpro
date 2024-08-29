package main

func nilaiAscend(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].nilai < A[idx].nilai {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func nilaiDescend(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].nilai > A[idx].nilai {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func jurusan(A *arrMahasiswa, n *int) { // selection sort
	var pass, idx, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		idx = pass - 1
		i = pass
		for i < *n {
			if A[i].jurusan < A[idx].jurusan {
				idx = i
			}
			i++
		}
		temp = A[pass - 1]
		A[pass - 1] = A[idx]
		A[idx] = temp
		pass++
	}
}

func namaAscend(A *arrMahasiswa, n *int) { // insertion sort
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama < A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}

func namaDescend(A *arrMahasiswa, n *int) { // insertion sort
	var pass, i int
	var temp mahasiswa

	pass = 1
	for pass <= *n - 1 {
		i = pass
		temp = A[pass]
		for i > 0 && temp.nama > A[i-1].nama {
			A[i] = A[i-1]
			i--
		}
		A[i] = temp
		pass++
	}
}