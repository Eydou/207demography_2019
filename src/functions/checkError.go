//
// EPITECH PROJECT, 2020
// 202unsold_2019
// File description:
// function
//

package functions

import (
	"errors"
)

//ErrorArgs check error
func ErrorArgs(args []string) (int, error) {
	if len(args) == 1 {
		return 84, errors.New("invalid arguments")
	}
	return 0, nil
}
