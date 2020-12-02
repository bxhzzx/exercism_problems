package erratum

func Use(opener func() (Resource, error), s string) (err error) { // 这里一定需要使用named return value，如果把err定义在函数里面，然后`return err`的话，虽然defer函数里面设置了err，但是在Use函数返回后err还是nil。
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
