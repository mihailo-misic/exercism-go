// Package erratum has tools for throwing errors.
package erratum

const testVersion = 2

// I've extracted the types and helper methods in another file (types.go).

// Use opens a resource and adds input to it, returning the error if any.
func Use(o ResourceOpener, input string) (err error) {
	r, err := o()
	_, ok := err.(TransientError)

	// If the error is Transient error try again, until it's not.
	for ok {
		r, err = o()
		_, ok = err.(TransientError)
	}
	// If there's an error opening and it's not Transient return it.
	if err != nil {
		return err
	}

	defer func() {
		// Recover from a panic (attack)!
		if rec := recover(); rec != nil {
			// Check if the error is a FrobError. If it is use Defrob.
			if v, forbErr := rec.(FrobError); forbErr {
				r.Defrob(v.defrobTag)
			}
			err, _ = rec.(error)
		}
		r.Close()
	}()

	r.Frob(input)

	return err
}
