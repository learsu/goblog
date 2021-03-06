package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

//标签表
type Tag struct {
	Id    int64
	Name  string `orm:"size(20);index"`
	Count int64
}

func (m *Tag) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Tag) Link() string {
	return fmt.Sprintf("<a class=\"category\" href=\"/category/%s\">%s</a>", Rawurlencode(m.Name), m.Name)
}
