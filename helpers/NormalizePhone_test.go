package helpers

import "testing"

func TestNormalizePhone(t *testing.T) {
	tests := []struct {
		name    string
		phone   string
		want    string
		wantErr bool
	}{
		{"adds country code to 11 digits", "(11) 91234-5678", "5511912345678", false},
		{"adds country code to 8 digits", "1234-5678", "5512345678", false},
		{"keeps 13 digits", "+55 11 91234-5678", "5511912345678", false},
		{"rejects empty", "  ", "", true},
		{"rejects letters only", "abc", "", true},
		{"rejects 12 digits", "551191234567", "", true},
		{"rejects 7 digits", "1234567", "", true},
		{"rejects 14 digits", "55119123456789", "", true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := NormalizePhone(test.phone)

			if (err != nil) != test.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, test.wantErr)
			}

			if got != test.want {
				t.Errorf("got %q, want %q", got, test.want)
			}
		})
	}
}
