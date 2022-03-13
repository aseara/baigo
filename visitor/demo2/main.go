package main

import "fmt"

// VisitorFunc func
type VisitorFunc func(info *Info, err error) error

// Visitor intf
type Visitor interface {
	Visit(VisitorFunc) error
}

// Info info
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

// Visit Visitor intf impl
func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

// NameVisitor name visitor
type NameVisitor struct {
	visitor Visitor
}

// Visit name visitor impl
func (v NameVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("NameVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> Name=%s, Namespace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("NameVisitor() after call function")
		return err
	})
}

// OtherThingsVisitor other things visitor
type OtherThingsVisitor struct {
	visitor Visitor
}

// Visit other things visitor impl
func (v OtherThingsVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("OtherThingsVisitor() before call function")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
		}
		fmt.Println("OtherThingsVisitor() after call function")
		return err
	})
}

// LogVisitor log visitor
type LogVisitor struct {
	visitor Visitor
}

// Visit other things visitor impl
func (v LogVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		fmt.Println("LogVisitor() before call function")
		err = fn(info, err)
		fmt.Println("LogVisitor() after call function")
		return err
	})
}

// DecoratedVisitor decorated visitor
type DecoratedVisitor struct {
	Visitor   Visitor
	decorates []VisitorFunc
}

// NewDecoratedVisitor factory method
func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit Visitor impl
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.Visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err = fn(info, nil); err != nil {
			return err
		}
		for _, d := range v.decorates {
			if err = d(info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	var info Info
	var v Visitor = &info
	v = LogVisitor{v}
	v = NameVisitor{v}
	v = OtherThingsVisitor{v}

	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}
	_ = v.Visit(loadFile)
	fmt.Println()

	d := NewDecoratedVisitor(&info, func(info *Info, err error) error {
		fmt.Println("log after call function")
		return err
	})

	_ = d.Visit(loadFile)
}
