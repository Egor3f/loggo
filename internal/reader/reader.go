package reader

type reader struct {
	strChan    chan string
	readerType Type
	onError    func(err error)
}

type Type = int64

const (
	TypeFile = Type(iota)
	TypePipe
	TypeGCP
)

// MakeReader builds a continues file/pipe streamer used to feed the logger. If
// fileName is not provided, it will attempt to consume the input from the stdin.
func MakeReader(fileName string, strChan chan string) Reader {
	if strChan == nil {
		strChan = make(chan string, 1)
	}
	if len(fileName) > 0 {
		return &fileStream{
			reader: reader{
				strChan:    strChan,
				readerType: TypeFile,
			},
			fileName: fileName,
		}
	}
	return &readPipeStream{
		reader: reader{
			strChan:    strChan,
			readerType: TypePipe,
		},
	}
}

func (s *reader) ChanReader() <-chan string {
	return s.strChan
}

func (s *reader) ErrorNotifier(onError func(err error)) {
	s.onError = onError
}

func (s *reader) Type() Type {
	return s.readerType
}

type Reader interface {
	// StreamInto feeds the strChan channel for every streamed line.
	StreamInto() error
	// Close finalises and invalidates this stream reader.
	Close()
	// ChanReader returns the outbound channel reader
	ChanReader() <-chan string
	// ErrorNotifier registers a callback func that's called upon fatal streaming log.
	ErrorNotifier(onError func(err error))
}
