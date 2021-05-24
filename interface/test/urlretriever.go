package test

type Retriever struct {}

func (Retriever) Get(url string) string  {
	return string(" I am faker :" + url)
}
