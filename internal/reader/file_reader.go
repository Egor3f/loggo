package reader

import (
	"fmt"

	"github.com/nxadm/tail"
)

type fileStream struct {
	reader
	fileName string
	tail     *tail.Tail
}

func (s *fileStream) StreamInto() error {
	var err error
	s.tail, err = tail.TailFile(s.fileName, tail.Config{Follow: true, Poll: true})
	if err != nil {
		return err
	}

	go func() {
		for line := range s.tail.Lines {
			s.strChan <- line.Text
		}
	}()
	return nil
}

func (s *fileStream) Close() {
	s.tail.Kill(fmt.Errorf("stopped by Close method"))
	close(s.strChan)
}
