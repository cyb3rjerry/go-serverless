package validators

import "regexp"

func IsEmailValid(email string) bool {
	var regEmail = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")

	return !(len(email) < 3 || len(email) > 254 || !regEmail.MatchString(email))

}
