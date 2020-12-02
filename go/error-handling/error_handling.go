package erratum

func Use(opener func() (Resource, error), s string) (err error) {
	resource, err := opener()
	for {
		if err != nil {
			if _, ok := err.(TransientError); !ok {
				return err
			} else {
				resource, err = opener()
				continue
			}
		} else {
			break
		}
	}

	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if _, ok = r.(FrobError); !ok {
				err = r.(error)
			} else {
				resource.Defrob(r.(FrobError).defrobTag)
				err = r.(FrobError)
			}
		}
		resource.Close()
	}()

	resource.Frob(s)
	return
}
