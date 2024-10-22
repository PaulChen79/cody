package handler

func Hello() (int, error) {
	r, err := TestFunc(1, 4)
	if err != nil {
		return 0, err
	}

	return r, nil
}

func TestFunc(a, b int) (int, error) {
	return a + b, nil
}
