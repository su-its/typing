package userimport

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/su-its/typing/typing-server/internal/domain/model"
)

const maxHandleNameLength = 36

func LoadUsersFromCSV(path string) ([]model.User, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open csv: %w", err)
	}
	defer file.Close()

	users, err := ParseUsersCSV(file)
	if err != nil {
		return nil, fmt.Errorf("parse csv %q: %w", path, err)
	}

	return users, nil
}

func ParseUsersCSV(r io.Reader) ([]model.User, error) {
	reader := csv.NewReader(r)
	reader.FieldsPerRecord = -1
	reader.TrimLeadingSpace = true

	var users []model.User
	seenStudentNumbers := make(map[string]int)
	rowNumber := 0

	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("row %d: %w", rowNumber+1, err)
		}

		rowNumber++
		if isBlankRecord(record) {
			continue
		}
		if len(record) < 2 {
			return nil, fmt.Errorf("row %d: expected at least 2 columns, got %d", rowNumber, len(record))
		}

		studentNumber := normalizeCSVCell(record[0])
		handleName := normalizeCSVCell(record[1])

		if isHeaderRow(studentNumber, handleName) {
			continue
		}
		if studentNumber == "" {
			return nil, fmt.Errorf("row %d: student number is empty", rowNumber)
		}
		if handleName == "" {
			return nil, fmt.Errorf("row %d: handle name is empty", rowNumber)
		}
		if utf8.RuneCountInString(handleName) > maxHandleNameLength {
			return nil, fmt.Errorf("row %d: handle name exceeds %d characters", rowNumber, maxHandleNameLength)
		}

		if firstSeenRow, exists := seenStudentNumbers[studentNumber]; exists {
			return nil, fmt.Errorf("row %d: duplicate student number %q (first seen at row %d)", rowNumber, studentNumber, firstSeenRow)
		}
		seenStudentNumbers[studentNumber] = rowNumber

		users = append(users, model.User{
			StudentNumber: studentNumber,
			HandleName:    handleName,
		})
	}

	if len(users) == 0 {
		return nil, errors.New("no users found in csv")
	}

	return users, nil
}

func normalizeCSVCell(value string) string {
	return strings.TrimSpace(strings.TrimPrefix(value, "\uFEFF"))
}

func isBlankRecord(record []string) bool {
	for _, cell := range record {
		if normalizeCSVCell(cell) != "" {
			return false
		}
	}
	return true
}

func isHeaderRow(studentNumber string, handleName string) bool {
	normalizedStudentNumber := normalizeHeaderCell(studentNumber)
	normalizedHandleName := normalizeHeaderCell(handleName)

	return isStudentNumberHeader(normalizedStudentNumber) && isHandleNameHeader(normalizedHandleName)
}

func normalizeHeaderCell(value string) string {
	replacer := strings.NewReplacer("_", "", "-", "", " ", "", "\u3000", "")
	return strings.ToLower(replacer.Replace(normalizeCSVCell(value)))
}

func isStudentNumberHeader(value string) bool {
	switch value {
	case "studentnumber", "学籍番号":
		return true
	default:
		return false
	}
}

func isHandleNameHeader(value string) bool {
	switch value {
	case "handlename", "name", "氏名", "名前":
		return true
	default:
		return false
	}
}
