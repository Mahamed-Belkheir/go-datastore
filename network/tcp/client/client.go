package tcp

type TCPClient struct {
	host string
	username string
	password string
}

func Client(host, username, password string) *TCPClient {
	return &TCPClient{
		host: host,
		username: username,
		password: password,
	}
}


func (c TCPClient) Get(key string) (value interface{}, err error) {

	return
}

func (c TCPClient) Set(key string, value interface{}) (err error) {

	return
}