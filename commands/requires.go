package commands

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"github.com/urfave/cli"
)

func Requires(c *cli.Context) error {
	name := c.Args().Get(0)

	if name == "" {
		return cli.NewExitError("missing name of environment variable that is required", 1)
	}

	value := os.Getenv(name)

	if value == "" {
		return cli.NewExitError(fmt.Sprintf("%s is required but not set or empty", name), 2)
	}

	message := "";
	switch c.String("validate") {
	case "int":
		intPattern, _ := regexp.Compile("^(-)?[0-9]+$")
		intValue, _ := strconv.ParseInt(value, 10, 64);

		if !intPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to be an int", name)
			break
		}

		if c.String("min") != "" {
			if intValue < c.Int64("min") {
				message = fmt.Sprintf("%s needs to be greater than or equal %d", name, c.Int64("min"))
				break
			}
		}

		if c.String("max") != "" {
			if intValue > c.Int64("max") {
				message = fmt.Sprintf("%s needs to be lower than or equal %d", name, c.Int64("max"))
				break
			}
		}
	case "float":
		floatPattern, _ := regexp.Compile("^(-)?[0-9]+(\\.[0-9]+)?$")
		floatValue, _ := strconv.ParseFloat(value, 64);

		if !floatPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to be an float", name)
			break
		}

		fmt.Print(c.String("min"))
		if c.String("min") != "" {
			if floatValue < c.Float64("min") {
				message = fmt.Sprintf("%s needs to be greater than or equal %f", name, c.Float64("min"))
				break
			}
		}

		if c.String("max") != "" {
			if floatValue > c.Float64("max") {
				message = fmt.Sprintf("%s needs to be lower than or equal %f", name, c.Float64("max"))
				break
			}
		}
	case "bool":
		if !(value == "true" || value == "false" || value == "1" || value == "0") {
			message = fmt.Sprintf("%s needs to be an boolean", name)
			break
		}
	case "regex":
		if c.String("pattern") == "" {
			message = "flag \"--pattern\" is needed for validator \"regex\""
			break
		}
		regexPattern, _ := regexp.Compile(c.String("pattern"))

		if !regexPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to match pattern \"%s\"", name, c.String("pattern"))
			break
		}
	}

	if message != "" {
		if c.String("example") != "" {
			return cli.NewExitError(fmt.Sprintf("%s (e.g. %s)", message, c.String("example")), 3)
		} else {
			return cli.NewExitError(message, 3)
		}
	}

	fmt.Printf("%s=%s", name, value)
	return nil
}
