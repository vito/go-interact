package interact_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/vito/go-interact/interact"
)

var _ = Describe("Resolving into bools", func() {
	Context("when the destination is false", func() {
		BeforeEach(func() {
			destination = boolDst(false)
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when 'y' is entered", Example{
				Prompt: "some prompt",

				Input: "y\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [yN]: y\r\n",
			}),

			Entry("when 'yes' is entered", Example{
				Prompt: "some prompt",

				Input: "yes\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [yN]: yes\r\n",
			}),

			Entry("when 'n' is entered", Example{
				Prompt: "some prompt",

				Input: "n\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [yN]: n\r\n",
			}),

			Entry("when 'no' is entered", Example{
				Prompt: "some prompt",

				Input: "no\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [yN]: no\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [yN]: \r\n",
			}),

			Entry("when a non-boolean is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: false,
				ExpectedErr:    io.EOF,
				ExpectedOutput: "some prompt [yN]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yN]: ",
			}),

			Entry("when a non-integer is entered, followed by 'y'", Example{
				Prompt: "some prompt",

				Input: "foo\ry\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [yN]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yN]: y\r\n",
			}),

			Entry("when a non-integer is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [yN]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yN]: \r\n",
			}),
		)

		Context("when required", func() {
			BeforeEach(func() {
				destination = interact.Required(destination)
			})

			DescribeTable("Resolve", (Example).Run,
				Entry("when 'y' is entered", Example{
					Prompt: "some prompt",

					Input: "y\r",

					ExpectedAnswer: true,
					ExpectedOutput: "some prompt [yn]: y\r\n",
				}),

				Entry("when 'yes' is entered", Example{
					Prompt: "some prompt",

					Input: "yes\r",

					ExpectedAnswer: true,
					ExpectedOutput: "some prompt [yn]: yes\r\n",
				}),

				Entry("when 'n' is entered", Example{
					Prompt: "some prompt",

					Input: "n\r",

					ExpectedAnswer: false,
					ExpectedOutput: "some prompt [yn]: n\r\n",
				}),

				Entry("when 'no' is entered", Example{
					Prompt: "some prompt",

					Input: "no\r",

					ExpectedAnswer: false,
					ExpectedOutput: "some prompt [yn]: no\r\n",
				}),

				Entry("when a blank line is entered, followed by EOF", Example{
					Prompt: "some prompt",

					Input: "\r",

					ExpectedAnswer: false,
					ExpectedErr:    io.EOF,
					ExpectedOutput: "some prompt [yn]: \r\nsome prompt [yn]: ",
				}),

				Entry("when a non-boolean is entered, followed by EOF", Example{
					Prompt: "some prompt",

					Input: "foo\r",

					ExpectedAnswer: false,
					ExpectedErr:    io.EOF,
					ExpectedOutput: "some prompt [yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yn]: ",
				}),

				Entry("when a non-integer is entered, followed by 'y'", Example{
					Prompt: "some prompt",

					Input: "foo\ry\r",

					ExpectedAnswer: true,
					ExpectedOutput: "some prompt [yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yn]: y\r\n",
				}),

				Entry("when a non-integer is entered, followed by a blank line, followed by EOF", Example{
					Prompt: "some prompt",

					Input: "foo\r\r",

					ExpectedAnswer: false,
					ExpectedErr:    io.EOF,
					ExpectedOutput: "some prompt [yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yn]: \r\nsome prompt [yn]: ",
				}),

				Entry("when a non-integer is entered, followed by a blank line, followed by 'y'", Example{
					Prompt: "some prompt",

					Input: "foo\r\ry\r",

					ExpectedAnswer: true,
					ExpectedOutput: "some prompt [yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [yn]: \r\nsome prompt [yn]: y\r\n",
				}),
			)
		})
	})

	Context("when the destination is true", func() {
		BeforeEach(func() {
			destination = boolDst(true)
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when 'y' is entered", Example{
				Prompt: "some prompt",

				Input: "y\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [Yn]: y\r\n",
			}),

			Entry("when 'yes' is entered", Example{
				Prompt: "some prompt",

				Input: "yes\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [Yn]: yes\r\n",
			}),

			Entry("when 'n' is entered", Example{
				Prompt: "some prompt",

				Input: "n\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [Yn]: n\r\n",
			}),

			Entry("when 'no' is entered", Example{
				Prompt: "some prompt",

				Input: "no\r",

				ExpectedAnswer: false,
				ExpectedOutput: "some prompt [Yn]: no\r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [Yn]: \r\n",
			}),

			Entry("when a non-boolean is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "foo\r",

				ExpectedAnswer: true,
				ExpectedErr:    io.EOF,
				ExpectedOutput: "some prompt [Yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [Yn]: ",
			}),

			Entry("when a non-integer is entered, followed by 'y'", Example{
				Prompt: "some prompt",

				Input: "foo\ry\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [Yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [Yn]: y\r\n",
			}),

			Entry("when a non-integer is entered, followed by a blank line", Example{
				Prompt: "some prompt",

				Input: "foo\r\r",

				ExpectedAnswer: true,
				ExpectedOutput: "some prompt [Yn]: foo\r\ninvalid input (not y, n, yes, or no)\r\nsome prompt [Yn]: \r\n",
			}),
		)
	})
})

func boolDst(dst bool) *bool {
	return &dst
}
