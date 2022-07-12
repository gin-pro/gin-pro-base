package sqlx

import (
	"reflect"
	"xorm.io/xorm"
)

type Page struct {
	Page       int         `json:"page"`
	Pages      int         `json:"pages"`
	Data       interface{} `json:"data"`
	Total      int         `json:"total"`
	TotalPages int         `json:"totalPages"`
}

func FindPage(ses *xorm.Session, page, size int, ls interface{}) (*Page, error) {
	p := 0
	s := 20
	if page > 1 {
		p = page - 1
	} else {
		page = 1
	}
	if size > 0 {
		s = size
	}
	bean := reflect.Value{}
	of := reflect.TypeOf(ls)
	if of.Kind() == reflect.Ptr {
		of = of.Elem()
	}

	if of.Kind() == reflect.Slice {
		elem := of.Elem()
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		bean = reflect.New(elem)
	}
	err := ses.Limit(s, s*p).Find(ls)
	if err != nil {
		return nil, err
	}
	count, err := ses.Count(bean)
	if err != nil {
		return nil, err
	}
	var ps int64 = 1
	if s > 0 {
		ps = count / int64(s)
		if count%int64(s) > 0 {
			ps += 1
		}
	}
	pg := &Page{
		Page:       page,
		Pages:      s,
		Data:       ls,
		Total:      int(count),
		TotalPages: int(ps),
	}
	return pg, err
}
