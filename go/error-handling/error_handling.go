package erratum

func Use(o ResourceOpener, input string) (err error) {
	var res Resource
	res, err = o()

	for err != nil {
		if _, ok := err.(TransientError); !ok {
			return
		}
		res, err = o()
	}

	defer res.Close()
	defer func() {
		if r := recover(); r != nil {
			switch errObj := r.(type) {
			case FrobError:
				res.Defrob(errObj.defrobTag)
				err = errObj.inner
			case error:
				err = errObj
			default:
				panic(errObj)
			}
		}
	}()

	res.Frob(input)
	return
}
