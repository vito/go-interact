package interact_test

import (
	"io"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/vito/go-interact/interact"
)

var _ = Describe("Resolving from a set of choices", func() {
	BeforeEach(func() {
		choices = []interact.Choice{
			{Display: "Uno", Value: arbitrary{"uno"}},
			{Display: "Dos", Value: arbitrary{"dos"}},
			{Display: "Tres", Value: arbitrary{"tres"}},
		}
	})

	Context("when the destination is zero-valued", func() {
		BeforeEach(func() {
			destination = arbDst(arbitrary{})
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when '0' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "0\r",

				ExpectedAnswer: arbitrary{},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 0\r\ninvalid selection (must be 1-3)\r\nsome prompt: ",
			}),

			Entry("when '0' is entered, followed by '1'", Example{
				Prompt: "some prompt",

				Input: "0\r1\r",

				ExpectedAnswer: arbitrary{"uno"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 0\r\ninvalid selection (must be 1-3)\r\nsome prompt: 1\r\n",
			}),

			Entry("when '1' is entered", Example{
				Prompt: "some prompt",

				Input: "1\r",

				ExpectedAnswer: arbitrary{"uno"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 1\r\n",
			}),

			Entry("when '2' is entered", Example{
				Prompt: "some prompt",

				Input: "2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 2\r\n",
			}),

			Entry("when '3' is entered", Example{
				Prompt: "some prompt",

				Input: "3\r",

				ExpectedAnswer: arbitrary{"tres"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 3\r\n",
			}),

			Entry("when '4' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "4\r",

				ExpectedAnswer: arbitrary{},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 4\r\ninvalid selection (must be 1-3)\r\nsome prompt: ",
			}),

			Entry("when '4' is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "4\r2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: 4\r\ninvalid selection (must be 1-3)\r\nsome prompt: 2\r\n",
			}),

			Entry("when a blank line is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: arbitrary{},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: \r\nsome prompt: ",
			}),

			Entry("when a blank line is entered, followed by '3'", Example{
				Prompt: "some prompt",

				Input: "\r3\r",

				ExpectedAnswer: arbitrary{"tres"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: \r\nsome prompt: 3\r\n",
			}),

			Entry("when a non-selection is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: arbitrary{},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: foo\r\ninvalid selection (not a number)\r\nsome prompt: ",
			}),

			Entry("when a non-selection is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "foo\r2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: foo\r\ninvalid selection (not a number)\r\nsome prompt: 2\r\n",
			}),

			Entry("when a non-integer is entered, followed by a blank line, followed by '3'", Example{
				Prompt: "some prompt",

				Input: "foo\r\r3\r",

				ExpectedAnswer: arbitrary{"tres"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt: foo\r\ninvalid selection (not a number)\r\nsome prompt: \r\nsome prompt: 3\r\n",
			}),
		)

		Context("when an unassignable choice is configured", func() {
			BeforeEach(func() {
				choices = append(choices, interact.Choice{
					Display: "Bogus",
					Value:   "bogus",
				})
			})

			DescribeTable("Resolve", (Example).Run,
				Entry("when the unassignable choice is chosen", Example{
					Prompt: "some prompt",

					Input: "4\r",

					ExpectedAnswer: arbitrary{},
					ExpectedErr: interact.NotAssignableError{
						Destination: reflect.TypeOf(arbitrary{}),
						Value:       reflect.TypeOf("bogus"),
					},
					ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: Bogus\r\nsome prompt: 4\r\n",
				}),
			)
		})
	})

	Context("when the destination is one of the choices", func() {
		BeforeEach(func() {
			destination = arbDst(arbitrary{"dos"})
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when '0' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "0\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 0\r\ninvalid selection (must be 1-3)\r\nsome prompt (2): ",
			}),

			Entry("when '0' is entered, followed by '1'", Example{
				Prompt: "some prompt",

				Input: "0\r1\r",

				ExpectedAnswer: arbitrary{"uno"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 0\r\ninvalid selection (must be 1-3)\r\nsome prompt (2): 1\r\n",
			}),

			Entry("when '1' is entered", Example{
				Prompt: "some prompt",

				Input: "1\r",

				ExpectedAnswer: arbitrary{"uno"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 1\r\n",
			}),

			Entry("when '2' is entered", Example{
				Prompt: "some prompt",

				Input: "2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 2\r\n",
			}),

			Entry("when '3' is entered", Example{
				Prompt: "some prompt",

				Input: "3\r",

				ExpectedAnswer: arbitrary{"tres"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 3\r\n",
			}),

			Entry("when '4' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "4\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 4\r\ninvalid selection (must be 1-3)\r\nsome prompt (2): ",
			}),

			Entry("when '4' is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "4\r2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): 4\r\ninvalid selection (must be 1-3)\r\nsome prompt (2): 2\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): \r\n",
			}),

			Entry("when a non-selection is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): ",
			}),

			Entry("when a non-selection is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "foo\r2\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): 2\r\n",
			}),

			Entry("when a non-integer is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: arbitrary{"dos"},
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): \r\n",
			}),
		)
	})

	Context("when the destination is nil and one of the choices is nil", func() {
		BeforeEach(func() {
			var emptyDst *arbitrary
			destination = &emptyDst

			choices = []interact.Choice{
				{Display: "Uno", Value: &arbitrary{"uno"}},
				{Display: "Dos", Value: &arbitrary{"dos"}},
				{Display: "Tres", Value: &arbitrary{"tres"}},
				{Display: "none", Value: nil},
			}
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when '0' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "0\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): ",
			}),

			Entry("when '0' is entered, followed by '1'", Example{
				Prompt: "some prompt",

				Input: "0\r1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): 1\r\n",
			}),

			Entry("when '1' is entered", Example{
				Prompt: "some prompt",

				Input: "1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 1\r\n",
			}),

			Entry("when '2' is entered", Example{
				Prompt: "some prompt",

				Input: "2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when '3' is entered", Example{
				Prompt: "some prompt",

				Input: "3\r",

				ExpectedAnswer: arbAns(arbitrary{"tres"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 3\r\n",
			}),

			Entry("when '4' is entered", Example{
				Prompt: "some prompt",

				Input: "4\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 4\r\n",
			}),

			Entry("when '5' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "5\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): ",
			}),

			Entry("when '5' is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "5\r2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): \r\n",
			}),

			Entry("when a non-selection is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): ",
			}),

			Entry("when a non-selection is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "foo\r2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when a non-selection is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): \r\n",
			}),
		)
	})

	Context("when the destination is nil and one of the choices is a typed nil", func() {
		BeforeEach(func() {
			var emptyDst *arbitrary
			destination = &emptyDst

			var nilDst *arbitrary

			choices = []interact.Choice{
				{Display: "Uno", Value: &arbitrary{"uno"}},
				{Display: "Dos", Value: &arbitrary{"dos"}},
				{Display: "Tres", Value: &arbitrary{"tres"}},
				{Display: "none", Value: nilDst},
			}
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when '0' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "0\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): ",
			}),

			Entry("when '0' is entered, followed by '1'", Example{
				Prompt: "some prompt",

				Input: "0\r1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): 1\r\n",
			}),

			Entry("when '1' is entered", Example{
				Prompt: "some prompt",

				Input: "1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 1\r\n",
			}),

			Entry("when '2' is entered", Example{
				Prompt: "some prompt",

				Input: "2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when '3' is entered", Example{
				Prompt: "some prompt",

				Input: "3\r",

				ExpectedAnswer: arbAns(arbitrary{"tres"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 3\r\n",
			}),

			Entry("when '4' is entered", Example{
				Prompt: "some prompt",

				Input: "4\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 4\r\n",
			}),

			Entry("when '5' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "5\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): ",
			}),

			Entry("when '5' is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "5\r2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): \r\n",
			}),

			Entry("when a non-selection is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: noArbAns(),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): ",
			}),

			Entry("when a non-selection is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "foo\r2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): 2\r\n",
			}),

			Entry("when a non-selection is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (4): foo\r\ninvalid selection (not a number)\r\nsome prompt (4): \r\n",
			}),
		)
	})

	Context("when the destination is by reference and one of the choices is nil", func() {
		BeforeEach(func() {
			dosDst := &arbitrary{"dos"}
			destination = &dosDst

			choices = []interact.Choice{
				{Display: "Uno", Value: &arbitrary{"uno"}},
				{Display: "Dos", Value: &arbitrary{"dos"}},
				{Display: "Tres", Value: &arbitrary{"tres"}},
				{Display: "none", Value: nil},
			}
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when '0' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "0\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (2): ",
			}),

			Entry("when '0' is entered, followed by '1'", Example{
				Prompt: "some prompt",

				Input: "0\r1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 0\r\ninvalid selection (must be 1-4)\r\nsome prompt (2): 1\r\n",
			}),

			Entry("when '1' is entered", Example{
				Prompt: "some prompt",

				Input: "1\r",

				ExpectedAnswer: arbAns(arbitrary{"uno"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 1\r\n",
			}),

			Entry("when '2' is entered", Example{
				Prompt: "some prompt",

				Input: "2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 2\r\n",
			}),

			Entry("when '3' is entered", Example{
				Prompt: "some prompt",

				Input: "3\r",

				ExpectedAnswer: arbAns(arbitrary{"tres"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 3\r\n",
			}),

			Entry("when '4' is entered", Example{
				Prompt: "some prompt",

				Input: "4\r",

				ExpectedAnswer: noArbAns(),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 4\r\n",
			}),

			Entry("when '5' is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "5\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (2): ",
			}),

			Entry("when '5' is entered, followed by '3'", Example{
				Prompt: "some prompt",

				Input: "5\r3\r",

				ExpectedAnswer: arbAns(arbitrary{"tres"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): 5\r\ninvalid selection (must be 1-4)\r\nsome prompt (2): 3\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): \r\n",
			}),

			Entry("when a non-selection is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedErr:    io.EOF,
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): ",
			}),

			Entry("when a non-selection is entered, followed by '2'", Example{
				Prompt: "some prompt",

				Input: "foo\r2\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): 2\r\n",
			}),

			Entry("when a non-selection is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: arbAns(arbitrary{"dos"}),
				ExpectedOutput: "1: Uno\r\n2: Dos\r\n3: Tres\r\n4: none\r\nsome prompt (2): foo\r\ninvalid selection (not a number)\r\nsome prompt (2): \r\n",
			}),
		)
	})
})

type arbitrary struct {
	value string
}

func arbDst(dst arbitrary) *arbitrary {
	return &dst
}

func arbAns(dst arbitrary) *arbitrary {
	return &dst
}

func noArbAns() *arbitrary {
	return nil
}
