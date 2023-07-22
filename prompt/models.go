package prompt

import "os"

type noBellStdout struct{}

func (bs *noBellStdout) Write(b []byte) (int, error) {
	const charBell = 7 // c.f. readline.CharBell
	if len(b) == 1 && b[0] == charBell {
		return 0, nil
	}
	return os.Stderr.Write(b)
}

// Close implements an io.WriterCloser over os.Stderr.
func (bs *noBellStdout) Close() error {
	return os.Stderr.Close()
}
