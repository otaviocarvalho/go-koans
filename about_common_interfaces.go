package go_koans

import "bytes"
import "io"

func aboutCommonInterfaces() {
	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)

		/*
		   Your code goes here.
		   Hint, use these resources:

		   $ godoc -http=:8080
		   $ open http://localhost:8080/pkg/io/
		   $ open http://localhost:8080/pkg/bytes/
		*/

        // CopyBuffer version
        buf := make([]byte, 8)
        io.CopyBuffer(out, in, buf)

        // Copy version
        io.Copy(out, in)

		assert(out.String() == "hello world") // get data from the io.Reader to the io.Writer
	}

	{
		in := new(bytes.Buffer)
		in.WriteString("hello world")

		out := new(bytes.Buffer)
        io.CopyN(out, in, 5)

		assert(out.String() == "hello") // duplicate only a portion of the io.Reader
	}
}
