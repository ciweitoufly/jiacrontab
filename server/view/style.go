package view

// cssPath 匹配样式文件
type StyleFinder struct {
	route string
	hashMap map[string]string
}

func NewStyleFinder() *StyleFinder{
	s := &StyleFinder{}
	return s
}

func Style(route string) (headerStyle string, foot string){

	return "", ""
} 






