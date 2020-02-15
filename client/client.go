package client

type Client interface {
	GetUser(accessToken string) (User, error)
	Login(email string, password string) (*string, error)
	Signup(email string, password string) error
}

type Options struct {
	Host string
}

type client struct {
	host string
}

func New(options *Options) Client {
	if options == nil {
		options = &Options{}
	}
	if options.Host == "" {
		options.Host = "localhost"
	}
	return &client{
		host: options.Host,
	}
}
