package erratum

func Use(opener ResourceOpener, input string) (err error) {
	newResource, err := opener()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(opener, input)
		}
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			if frob, ok := r.(FrobError); ok {
				newResource.Defrob(frob.defrobTag)
			}
			err = r.(error)
		}
		newResource.Close()
	}()

	newResource.Frob(input)
	return
}
