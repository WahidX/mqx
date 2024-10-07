package errs

type xError struct {
  err error
  detail string
  severity int
}

