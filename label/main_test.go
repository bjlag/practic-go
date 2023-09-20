package label

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestLabel(t *testing.T) {
    testCases := []struct {
        name     string
        str      string
        labels   []Label
        expected string
    }{
        {
            "one",
            "somesuperlongstring",
            []Label{
                {1, 2, "la"},
                {4, 4, "d"},
                {10, 13, "tiny"},
            },
            "lamdsupertinystring",
        },
        {
            "two",
            "somesuperlongstring",
            []Label{
                {1, 2, "la"},
                {4, 4, "d"},
                {10, 13, "tiny"},
                {4, 5, "ed"},
            },
            "lamedupertinystring",
        },
        {
            "no labels",
            "somesuperlongstring",
            []Label{},
            "somesuperlongstring",
        },
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := label(tc.str, tc.labels)
            require.Equal(t, tc.expected, result)
        })
    }
}
