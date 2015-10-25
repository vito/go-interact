package interact_test

import (
	"io"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	"github.com/vito/go-interact/interact"
)

var _ = Describe("Resolving into passwords", func() {
	Context("when the destination is empty", func() {
		BeforeEach(func() {
			destination = passDst("")
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when a string is entered", Example{
				Prompt: "some prompt",

				Input: "forty two\r",

				ExpectedAnswer: interact.Password("forty two"),
				ExpectedOutput: "some prompt (): \r\n",
			}),

			Entry("when a blank line is entered, followed by EOF", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: interact.Password(""),
				ExpectedOutput: "some prompt (): \r\n",
			}),
		)

		Context("when required", func() {
			BeforeEach(func() {
				destination = interact.Required(destination)
			})

			DescribeTable("Resolve", (Example).Run,
				Entry("when a string is entered", Example{
					Prompt: "some prompt",

					Input: "forty two\r",

					ExpectedAnswer: interact.Password("forty two"),
					ExpectedOutput: "some prompt: \r\n",
				}),

				Entry("when a blank line is entered, followed by EOF", Example{
					Prompt: "some prompt",

					Input: "\r",

					ExpectedAnswer: interact.Password(""),
					ExpectedErr:    io.EOF,
					ExpectedOutput: "some prompt: \r\nsome prompt: ",
				}),

				Entry("when a blank line is entered, followed by a string", Example{
					Prompt: "some prompt",

					Input: "\rforty two\r",

					ExpectedAnswer: interact.Password("forty two"),
					ExpectedOutput: "some prompt: \r\nsome prompt: \r\n",
				}),
			)
		})
	})

	Context("when the destination is not empty", func() {
		BeforeEach(func() {
			destination = passDst("some default")
		})

		DescribeTable("Resolve", (Example).Run,
			Entry("when a string is entered", Example{
				Prompt: "some prompt",

				Input: "forty two\r",

				ExpectedAnswer: interact.Password("forty two"),
				ExpectedOutput: "some prompt (has default): \r\n",
			}),

			Entry("when a blank line is entered", Example{
				Prompt: "some prompt",

				Input: "\r",

				ExpectedAnswer: interact.Password("some default"),
				ExpectedOutput: "some prompt (has default): \r\n",
			}),
		)
	})
})

func passDst(dst interact.Password) *interact.Password {
	return &dst
}
