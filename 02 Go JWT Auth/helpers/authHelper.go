package helper

import (
	"errors"
	"fmt"
	"net/http"
)

func CheckUserType(w http.ResponseWriter, role string) (err error) {
	userType := w.Header().Get("user_type")
	fmt.Printf("helllo user type is %v \n \n", userType)
	err = nil

	if userType != role {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	return err
}

func MatchUserTypeToUid(w http.ResponseWriter, userId string) (err error) {
	userType := w.Header().Get("user_type")
	uid := w.Header().Get("user_id")
	err = nil

	if userType == "USER" && uid != userId {
		err = errors.New("Unauthorized to access this resource")
		return err
	}

	err = CheckUserType(w, userType)
	return err
}
