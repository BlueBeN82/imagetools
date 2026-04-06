package commands

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/urfave/cli/v3"
)

func Requires(ctx context.Context, cmd *cli.Command) error {
	name := cmd.Args().Get(0)

	if name == "" {
		return cli.Exit("missing name of environment variable that is required", 1)
	}

	value := os.Getenv(name)

	if value == "" {
		return cli.Exit(fmt.Sprintf("%s is required but not set or empty", name), 2)
	}

	message := ""
	switch cmd.String("validate") {
	case "int":
		intPattern, _ := regexp.Compile("^(-)?[0-9]+$")
		intValue, _ := strconv.ParseInt(value, 10, 64)

		if !intPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to be an int", name)
			break
		}

		if cmd.String("min") != "" {
			if intValue < cmd.Int64("min") {
				message = fmt.Sprintf("%s needs to be greater than or equal %d", name, cmd.Int64("min"))
				break
			}
		}

		if cmd.String("max") != "" {
			if intValue > cmd.Int64("max") {
				message = fmt.Sprintf("%s needs to be lower than or equal %d", name, cmd.Int64("max"))
				break
			}
		}
	case "float":
		floatPattern, _ := regexp.Compile("^(-)?[0-9]+(\\.[0-9]+)?$")
		floatValue, _ := strconv.ParseFloat(value, 64)

		if !floatPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to be an float", name)
			break
		}

		fmt.Print(cmd.String("min"))
		if cmd.String("min") != "" {
			if floatValue < cmd.Float64("min") {
				message = fmt.Sprintf("%s needs to be greater than or equal %f", name, cmd.Float64("min"))
				break
			}
		}

		if cmd.String("max") != "" {
			if floatValue > cmd.Float64("max") {
				message = fmt.Sprintf("%s needs to be lower than or equal %f", name, cmd.Float64("max"))
				break
			}
		}
	case "bool":
		if !(value == "true" || value == "false" || value == "1" || value == "0") {
			message = fmt.Sprintf("%s needs to be an boolean", name)
			break
		}
	case "regex":
		if cmd.String("pattern") == "" {
			message = "flag \"--pattern\" is needed for validator \"regex\""
			break
		}
		regexPattern, _ := regexp.Compile(cmd.String("pattern"))

		if !regexPattern.MatchString(value) {
			message = fmt.Sprintf("%s needs to match pattern \"%s\"", name, cmd.String("pattern"))
			break
		}
	}

	if message != "" {
		if cmd.String("example") != "" {
			return cli.Exit(fmt.Sprintf("%s (e.g. %s)", message, cmd.String("example")), 3)
		} else {
			return cli.Exit(message, 3)
		}
	}

	fmt.Printf("%s=%s", name, value)
	return nil
}
