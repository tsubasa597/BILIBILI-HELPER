package ecode

type APIErr struct {
	E   string
	Msg string
}

var (
	_ error = (*APIErr)(nil)
)

func (err APIErr) Error() string {
	if err.Msg != "" {
		return err.E + ": " + err.Msg
	}

	return err.E
}
