package cron

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const scheduleComponentsSize = 5

var scheduleParts []string = []string{
	"minute",
	"hour",
	"date",
	"month",
	"dayOfWeek",
}

var schedulePartRanges map[string][2]int = map[string][2]int{
	"minute":    {0, 59},
	"hour":      {0, 23},
	"date":      {1, 31},
	"month":     {1, 12},
	"dayOfWeek": {1, 7},
}

// ParseSchedule converts a schedule string into its possible date/time component options
func ParseSchedule(schedule string) (map[string][]int, error) {
	scheduleComponents := strings.Fields(strings.TrimSpace(schedule))

	if len(scheduleComponents) != scheduleComponentsSize {
		return nil, fmt.Errorf("Invalid schedule components size. Expected %v components for the cron schedule", scheduleComponentsSize)
	}

	var result map[string][]int = map[string][]int{}
	for index, component := range scheduleComponents {
		switch {
		case component == "*":
			// all possible values option
			result[scheduleParts[index]] = parseContinuous(scheduleParts[index])
		default:
			// range of values option
			match, _ := regexp.MatchString("\\d{1,2}-\\d{1,2}", component)
			if match {
				possibleValues, err := parseRange(component, scheduleParts[index])

				if err != nil {
					fmt.Println("error has happened", err)
					return nil, err
				}
				result[scheduleParts[index]] = possibleValues
				continue
			}

			// list of values option
			match, _ = regexp.MatchString("\\d{1,2}(,\\d{1,2})+", component)
			if match {
				possibleValues, err := parseList(component, scheduleParts[index])

				if err != nil {
					return nil, err
				}
				result[scheduleParts[index]] = possibleValues
				continue
			}

			match, _ = regexp.MatchString("\\d{1,2}/\\d{1,2}", component)
			if match {
				possibleValues, err := parseRecurring(component, scheduleParts[index])

				if err != nil {
					return nil, err
				}
				result[scheduleParts[index]] = possibleValues
				continue
			}

			match, _ = regexp.MatchString("\\*/\\d{1,2}", component)
			if match {
				possibleValues, err := parseRecurring(component, scheduleParts[index])

				if err != nil {
					return nil, err
				}
				result[scheduleParts[index]] = possibleValues
				continue
			}

			match, _ = regexp.MatchString("\\d{1,2}", component)
			if match {
				value, err := strconv.Atoi(strings.TrimSpace(component))
				if err != nil {
					return nil, err
				}
				var possibleValues []int = []int{value}
				result[scheduleParts[index]] = possibleValues
				continue
			}

		}
	}

	return result, nil
}

func parseContinuous(schedulePart string) []int {
	var min = schedulePartRanges[schedulePart][0]
	var max = schedulePartRanges[schedulePart][1]
	var possibleValues []int = []int{}
	for i := min; i <= max; i++ {
		possibleValues = append(possibleValues, i)
	}

	return possibleValues
}

func parseRange(component string, schedulePart string) ([]int, error) {
	var componentParts = strings.Split(component, "-")
	start, err := strconv.Atoi(componentParts[0])
	if err != nil {
		return nil, err
	}
	end, err := strconv.Atoi(componentParts[1])
	if err != nil {
		return nil, err
	}

	if start > end {
		return nil, fmt.Errorf("Start value %v cannot be greater than end value %v", start, end)
	}

	var min = schedulePartRanges[schedulePart][0]
	var max = schedulePartRanges[schedulePart][1]

	if start < min || start > max {
		return nil, fmt.Errorf("Value %v is out of range for %s ", start, schedulePart)
	}

	if end < min || end > max {
		return nil, fmt.Errorf("Value %v is out of range for %s ", end, schedulePart)
	}

	var possibleValues []int = []int{}

	for i := start; i <= end; i++ {
		possibleValues = append(possibleValues, i)
	}

	return possibleValues, nil
}

func parseList(component string, schedulePart string) ([]int, error) {
	var possibleValues []int = []int{}
	var min = schedulePartRanges[schedulePart][0]
	var max = schedulePartRanges[schedulePart][1]

	var componentParts = strings.Split(component, ",")

	for _, option := range componentParts {
		optionValue, err := strconv.Atoi(strings.TrimSpace(option))

		if err != nil {
			return nil, err
		}

		if optionValue < min || optionValue > max {
			return nil, fmt.Errorf("Value %v is out of range for %s component", optionValue, schedulePart)
		}

		possibleValues = append(possibleValues, optionValue)
	}

	return possibleValues, nil
}

func parseRecurring(component string, schedulePart string) ([]int, error) {
	var possibleValues []int = []int{}
	var componentParts = strings.Split(component, "/")
	var max = schedulePartRanges[schedulePart][1]
	var min = schedulePartRanges[schedulePart][0]

	start := min
	var err error
	if strings.TrimSpace(componentParts[0]) != "*" {
		start, err = strconv.Atoi(strings.TrimSpace(componentParts[0]))
		if err != nil {
			return nil, err
		}
	}

	incrementer, err := strconv.Atoi(strings.TrimSpace(componentParts[1]))
	if err != nil {
		return nil, err
	}

	for i := start; i <= max; i += incrementer {
		possibleValues = append(possibleValues, i)
	}

	return possibleValues, nil
}
