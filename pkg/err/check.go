package err

import "github.com/fatih/color"

func Check(err error, success string, args ...interface{}) error {
	if err != nil {
		color.Red("Error: %s", err)
		return err
	}
	color.Green(success, args...)
	return nil
}
