package core

type HttpRequester interface {
	//get request
	Get(path string, queries map[string]string, auth bool, response Response) error

	//post request
	Post(path string, body []byte, queries map[string]string, auth bool, response Response) error
}

type Response interface {
	//check ret field
	CheckRet() error
}
