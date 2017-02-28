package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Screens struct {
	Id         int      `orm:"column(id);auto"`
	Type       string   `orm:"column(type);size(255)"`
	Position   int      `orm:"column(position)"`
	Text       string   `orm:"column(text);null"`
	VideoUrl   string   `orm:"column(video_url);size(255);null"`
	AudioUrl   string   `orm:"column(audio_url);size(255);null"`
	OptionsUrl string   `orm:"column(options_url);null"`
	ImagesUrl  string   `orm:"column(images_url);null"`
	LessonId   *Lessons `orm:"column(lesson_id);rel(fk)"`
}

func (t *Screens) TableName() string {
	return "screens"
}

func init() {
	orm.RegisterModel(new(Screens))
}

// AddScreens insert a new Screens into database and returns
// last inserted Id on success.
func AddScreens(m *Screens) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetScreensById retrieves Screens by Id. Returns error if
// Id doesn't exist
func GetScreensById(id int) (v *Screens, err error) {
	o := orm.NewOrm()
	v = &Screens{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetScreensByLessonId retrieves Screens by lesson_id. Returns error if
// Id doesn't exist
func GetScreensByLessonId(id int) (v *Screens, err error) {
	o := orm.NewOrm()
	l := &Lessons{Id: id}
	v = &Screens{LessonId: l}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetScreenByLessonAndScreenId retrieves Screens by lesson_id and screen_id. Returns error if
// Id doesn't exist
func GetScreenByLessonAndScreenId(lid int, sid int) (v *Screens, err error) {
	o := orm.NewOrm()
	l := &Lessons{Id: lid}
	v = &Screens{LessonId: l, Id: sid}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//GetScreenByLessonAndPositionId retrieves Screens by lesson_id and position. Returns error if
// Id doesn't exist
func GetScreenByLessonAndPositionId(lid int, position int) (v *Screens, err error) {
	o := orm.NewOrm()
	l := &Lessons{Id: lid}
	v = &Screens{LessonId: l, Position: position}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllScreens retrieves all Screens matches certain condition. Returns empty list if
// no records exist
func GetAllScreens(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Screens))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Screens
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateScreens updates Screens by Id and returns error if
// the record to be updated doesn't exist
func UpdateScreensById(m *Screens) (err error) {
	o := orm.NewOrm()
	v := Screens{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteScreens deletes Screens by Id and returns error if
// the record to be deleted doesn't exist
func DeleteScreens(id int) (err error) {
	o := orm.NewOrm()
	v := Screens{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Screens{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
