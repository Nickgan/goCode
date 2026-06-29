package dto

import (
	"encoding/json"
	"fmt"
)

type UserStatus int8

var (
	UserStatusNormal  UserStatus = 1
	UserStatusDeleted UserStatus = 0
	UserStatusBanned  UserStatus = -1
)

func (s UserStatus) MarshalJSON() (data []byte, err error) {

	var str string
	switch s {
	case UserStatusNormal:
		str = "normal"
	case UserStatusBanned:
		str = "banned"
	case UserStatusDeleted:
		str = "deleted"
	}

	return json.Marshal(map[string]any{
		"status":     int8(s),
		"statusName": str,
	})
}

func (s *UserStatus) UnmarshalJSON(data []byte) error {
	var obj map[string]any
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	status, ok := obj["status"]
	if !ok {
		return fmt.Errorf("missing status field")
	}

	switch v := status.(type) {
	case float64:
		*s = UserStatus(int8(v))
	case int8:
		*s = UserStatus(v)
	case int:
		*s = UserStatus(v)
	}

	return nil
}
