package dataloader

func FillErrorSlice(fillWith error, slice []error) []error {

	for i := 0; i < len(slice); i++ {
		slice[i] = fillWith
	}

	return slice

}
