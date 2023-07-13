package handler

import (
	"errors"
	"testing"
	"text-cdn/internal/core"
)

func (h *Handler) validateMessage_test(t *testing.T) {

	type TestCase struct {
		core.Message
		Expected error
	}

	testCases := []TestCase{
		{
			core.Message{
				Type:     "rule",
				Language: "ru",
			},
			nil,
		},
		{
			core.Message{
				Type:     "rule",
				Language: "",
			},
			errors.New(SomeJSONFiledsEmpty),
		},
		{
			core.Message{
				Type:     "",
				Language: "ru",
			},
			errors.New(SomeJSONFiledsEmpty),
		},
		{
			core.Message{
				Type:     "",
				Language: "",
			},
			errors.New(SomeJSONFiledsEmpty),
		},
	}

	for _, test := range testCases {
		err := h.validateMessage(test.Message)
		if err != test.Expected {
			t.Errorf("incorrect result: expected: %v, received: %v", test.Expected, err)
		}
	}
}
