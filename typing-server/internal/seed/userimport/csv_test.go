package userimport

import (
	"strings"
	"testing"
)

func TestParseUsersCSV(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		csv         string
		wantCount   int
		wantFirstID string
		wantFirstHN string
		wantErr     string
	}{
		{
			name:        "headerなしの2列CSVを読み込める",
			csv:         "\uFEFF725A0701,AUNG BHONE PYAE KYAW\n726A0001,青山　哲\n",
			wantCount:   2,
			wantFirstID: "725A0701",
			wantFirstHN: "AUNG BHONE PYAE KYAW",
		},
		{
			name:        "header行を読み飛ばせる",
			csv:         "student_number,handle_name\n725A0701,AUNG BHONE PYAE KYAW\n",
			wantCount:   1,
			wantFirstID: "725A0701",
			wantFirstHN: "AUNG BHONE PYAE KYAW",
		},
		{
			name:    "列不足はエラー",
			csv:     "725A0701\n",
			wantErr: "expected at least 2 columns",
		},
		{
			name:    "重複学籍番号はエラー",
			csv:     "725A0701,AUNG BHONE PYAE KYAW\n725A0701,青山　哲\n",
			wantErr: "duplicate student number",
		},
		{
			name:    "空ファイルはエラー",
			csv:     "\n",
			wantErr: "no users found in csv",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := ParseUsersCSV(strings.NewReader(tt.csv))
			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("ParseUsersCSV() error = nil, want substring %q", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("ParseUsersCSV() error = %q, want substring %q", err.Error(), tt.wantErr)
				}
				return
			}

			if err != nil {
				t.Fatalf("ParseUsersCSV() error = %v", err)
			}
			if len(got) != tt.wantCount {
				t.Fatalf("ParseUsersCSV() count = %d, want %d", len(got), tt.wantCount)
			}
			if got[0].StudentNumber != tt.wantFirstID {
				t.Fatalf("ParseUsersCSV() first studentNumber = %q, want %q", got[0].StudentNumber, tt.wantFirstID)
			}
			if got[0].HandleName != tt.wantFirstHN {
				t.Fatalf("ParseUsersCSV() first handleName = %q, want %q", got[0].HandleName, tt.wantFirstHN)
			}
		})
	}
}
