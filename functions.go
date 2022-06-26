package algobra

import "errors"

func NewMatrix[N Number](matr [][]N) (m *Matrix[N]) {
	m.RowsNumber = len(matr)
	m.ColumnsNumber = len(matr[0])
	m.Matr = matr
	return m
}

func NewEmptyMatrix[N Number](rowsNumber int, columnsNumber int) (m *Matrix[N]) {
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = columnsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
	}
	return m
}

func IsSquare[N Number](m *Matrix[N]) bool {
	return m.RowsNumber == m.ColumnsNumber
}

func NewIdentityMatrix[N Number](rowsNumber int) (m *Matrix[N]) {
	m.RowsNumber = rowsNumber
	m.ColumnsNumber = rowsNumber
	m.Matr = make([][]N, m.RowsNumber)
	for i := 0; i < m.RowsNumber; i++ {
		m.Matr[i] = make([]N, m.ColumnsNumber)
		m.Matr[i][i] = 1
	}
	return m
}

func MakeCopy[N Number](m *Matrix[N]) (ans *Matrix[N]) {
	ans = NewEmptyMatrix[N](m.RowsNumber, m.ColumnsNumber)
	for i, row := range m.Matr {
		for j, num := range row {
			ans.Matr[i][j] = num
		}
	}
	return ans
}

func Copy[N Number](from *Matrix[N], to *Matrix[N]) {
	to = MakeCopy[N](from)
}

func Negative[N Number](n interface{}) (interface{}, error) {
	switch obj := n.(type) {
	case N:
		var ans N
		ans = obj * -1
		return ans, nil

	case Matrix[N]:
		var ans Matrix[N]
		ans.Copy(&obj)
		for i := 0; i < ans.RowsNumber; i++ {
			for j := 0; j < ans.ColumnsNumber; j++ {
				ans.Matr[i][j] *= -1
			}
		}
		return ans, nil

	case *Matrix[N]:
		var ans Matrix[N]
		ans.Copy(obj)
		for i := 0; i < ans.RowsNumber; i++ {
			for j := 0; j < ans.ColumnsNumber; j++ {
				ans.Matr[i][j] *= -1
			}
		}
		return ans, nil

	default:
		return nil, invalidTypeError
	}
}

func Add[N Number](m *Matrix[N], n interface{}) (ans *Matrix[N], err error) {
	switch obj := n.(type) {
	case N:
		ans.Copy(m)
		for i := 0; i < m.RowsNumber; i++ {
			for j := 0; j < m.ColumnsNumber; j++ {
				ans.Matr[i][j] += obj
			}
		}

	case Matrix[N]:
		ans.Copy(m)
		for i := 0; i < m.RowsNumber; i++ {
			for j := 0; j < m.ColumnsNumber; j++ {
				ans.Matr[i][j] += obj.Matr[i][j]
			}
		}

	case *Matrix[N]:
		ans.Copy(m)
		for i := 0; i < m.RowsNumber; i++ {
			for j := 0; j < m.ColumnsNumber; j++ {
				ans.Matr[i][j] += obj.Matr[i][j]
			}
		}

	default:
		return nil, invalidTypeError
	}
	return ans, nil
}

func Substract[N Number](m *Matrix[N], n interface{}) (*Matrix[N], error) {
	neg, err := Negative[N](n)
	if err != nil {
		return nil, err
	}
	return Add(m, neg)
}

func Multiply[N Number](m1 *Matrix[N], m2 *Matrix[N]) (ans *Matrix[N], err error) {
	if m1.ColumnsNumber != m2.RowsNumber {
		return nil, errors.New("number of columns of first matrix" +
			" " + "has to be equal to number of rows of second matrix")
	}

	ans = NewEmptyMatrix[N](m1.GetRowsNumber(), m2.GetColumnsNumber())

	for h := 0; h < ans.RowsNumber; h++ {
		for i := 0; i < ans.ColumnsNumber; i++ {
			ans.Matr[h][i] = 0
			for j := 0; j < m1.ColumnsNumber; j++ {
				ans.Matr[h][i] += m1.Matr[h][j] * m2.Matr[j][i]
			}
		}
	}

	return ans, nil
}

func Power[N Number](m *Matrix[N], pow int) (ans *Matrix[N], err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}

	ans = NewIdentityMatrix[N](m.RowsNumber)

	for pow > 0 {
		if pow%2 == 0 {
			pow /= 2
			ans, err = ans.MultiplyTo(ans)
			if err != nil {
				return nil, err
			}
		} else {
			pow--
			ans, _ = ans.MultiplyTo(m)
		}
	}

	return ans, nil
}

func ToBinary[I Int](m *Matrix[I]) (ans *Matrix[I], err error) {
	ans.Copy(m)
	for i := 0; i < ans.RowsNumber; i++ {
		for j := 0; j < ans.ColumnsNumber; j++ {
			ans.Matr[i][j] = m.Matr[i][j] % 2
		}
	}
	return ans, err
}

func Det[N Number](m *Matrix[N]) (num N, err error) {
	if !m.IsSquare() {
		return nil, squareMatrixError
	}
	return nil, nil
}