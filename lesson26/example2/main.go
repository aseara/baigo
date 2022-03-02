package main

// Result type
type Result struct {
	Count int
}

// Int return count
func (r Result) Int() int { return r.Count }

// Rows type
type Rows []struct{}

// Stmt intf
type Stmt interface {
	Close() error
	NumInput() int
	Exec(stmt string, args ...string) (Result, error)
	Query(args []string) (Rows, error)
}

// MaleCount count male
func MaleCount(s Stmt) (int, error) {
	result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")
	if err != nil {
		return 0, err
	}
	return result.Int(), nil
}
