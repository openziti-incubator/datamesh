package datamesh

import (
	"github.com/sirupsen/logrus"
	"io"
)

type NICAdapter struct {
	nic NIC
}

func NewNICAdapter(nic NIC) *NICAdapter {
	return &NICAdapter{nic}
}

func (na *NICAdapter) Read(p []byte) (n int, err error) {
	select {
	case buf, ok := <-na.nic.(*nicImpl).netq:
		if ok {
			n := copy(p, buf.Data[:buf.Used])
			logrus.Infof("read (%v)", n)
			return n, nil
		}
	}
	return 0, io.EOF
}

func (na *NICAdapter) Write(p []byte) (n int, err error) {
	if err := na.nic.(*nicImpl).Tx(p); err != nil {
		return 0, err
	}
	n = len(p)
	logrus.Infof("wrote (%v)", n)
	return n	, nil
}

func (na *NICAdapter) Close() error {
	return nil
}
