package env

import (
	"errors"
	"os"
	"testing"
)

func TestStr(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		envs     map[string]string
		fallback string
		want     string
	}{
		{
			name: "valid key and value",
			key:  "STR_1",
			envs: map[string]string{
				"STR_1": "VAL_1",
			},
			fallback: "fallback",
			want:     "VAL_1",
		},
		{
			name: "invalid key",
			key:  "STR_2",
			envs: map[string]string{
				"STR_1": "VAL_1",
			},
			fallback: "fallback",
			want:     "fallback",
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)

		setupEnv(test.envs, t)
		got := Str(test.key, test.fallback)

		if got != test.want {
			t.Fatalf("got %q, want %q", got, test.want)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		envs     map[string]string
		fallback int
		want     int
	}{
		{
			name: "valid key and value",
			key:  "INT_1",
			envs: map[string]string{
				"INT_1": "1",
			},
			fallback: 0,
			want:     1,
		},
		{
			name: "invalid key",
			key:  "INT_2",
			envs: map[string]string{
				"INT_1": "1",
			},
			fallback: 0,
			want:     0,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)

		setupEnv(test.envs, t)
		got := Int(test.key, test.fallback)

		if got != test.want {
			t.Fatalf("got %d, want %d", got, test.want)
		}
	}
}
func TestBool(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		envs     map[string]string
		fallback bool
		want     bool
	}{
		{
			name: "valid key and 'true'",
			key:  "BOOL_1",
			envs: map[string]string{
				"BOOL_1": "true",
			},
			fallback: false,
			want:     true,
		},
		{
			name: "valid key and 'false'",
			key:  "BOOL_1",
			envs: map[string]string{
				"BOOL_1": "false",
			},
			fallback: true,
			want:     false,
		},
		{
			name: "invalid key",
			key:  "BOOL_2",
			envs: map[string]string{
				"BOOL_1": "true",
			},
			fallback: false,
			want:     false,
		},
		{
			name: "case insensitive and 'true'",
			key:  "BOOL_1",
			envs: map[string]string{
				"BOOL_1": "TrUe",
			},
			fallback: false,
			want:     true,
		},
		{
			name: "case insensitive and 'false'",
			key:  "BOOL_1",
			envs: map[string]string{
				"BOOL_1": "fAlSe",
			},
			fallback: true,
			want:     false,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)

		setupEnv(test.envs, t)
		got := Bool(test.key, test.fallback)

		if got != test.want {
			t.Fatalf("got %v, want %v", got, test.want)
		}
	}
}

func TestFloat(t *testing.T) {
	tests := []struct {
		name     string
		key      string
		envs     map[string]string
		fallback float32
		want     float32
	}{
		{
			name: "valid key and value",
			key:  "FLOAT_1",
			envs: map[string]string{
				"FLOAT_1": "1.15",
			},
			fallback: 0.0,
			want:     1.15,
		},
		{
			name: "invalid key",
			key:  "FLOAT_2",
			envs: map[string]string{
				"FLOAT_1": "1.15",
			},
			fallback: 2.00,
			want:     2.00,
		},
		{
			name: "no decimals",
			key:  "FLOAT_1",
			envs: map[string]string{
				"FLOAT_1": "1",
			},
			fallback: 0.00,
			want:     1.00,
		},
		{
			name: "trailing decimail point",
			key:  "FLOAT_1",
			envs: map[string]string{
				"FLOAT_1": "1.",
			},
			fallback: 0.00,
			want:     1.00,
		},
		{
			name: "out of range",
			key:  "FLOAT_1",
			envs: map[string]string{
				"FLOAT_1": "1.500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000500000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
			},
			fallback: 0.0,
			want:     1.5,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)

		setupEnv(test.envs, t)
		got := Float32(test.key, test.fallback)

		if got != test.want {
			t.Fatalf("got %.4f, want %.4f", got, test.want)
		}
	}
}

func TestSetStr(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  string
		want error
	}{
		{
			name: "valid key and value",
			key:  "KEY_1",
			val:  "VAL_1",
			want: nil,
		},
		{
			name: "empty key, valid value",
			key:  "",
			val:  "VAL_1",
			want: ErrInvalidKey,
		},
		{
			name: "quote in key, valid value",
			key:  `KEY_1"`,
			val:  "VAL_1",
			want: ErrInvalidKey,
		},
		{
			name: "valid key, quote in value",
			key:  "KEY_1",
			val:  `VAL_1"`,
			want: ErrInvalidValue,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)
		os.Clearenv()
		gotErr := SetStr(test.key, test.val)

		if !errors.Is(gotErr, test.want) {
			t.Fatalf("got %q, want %q", gotErr, test.want)
		}
	}
}

func TestSetInt(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  int
		want error
	}{
		{
			name: "valid key and value",
			key:  "KEY_1",
			val:  1,
			want: nil,
		},
		{
			name: "empty key, valid value",
			key:  "",
			val:  1,
			want: ErrInvalidKey,
		},
		{
			name: "quote in key, valid value",
			key:  `KEY_1"`,
			val:  1,
			want: ErrInvalidKey,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)
		os.Clearenv()
		gotErr := SetInt(test.key, test.val)

		if !errors.Is(gotErr, test.want) {
			t.Fatalf("got %q, want %q", gotErr, test.want)
		}
	}
}

func TestSetBool(t *testing.T) {
	tests := []struct {
		name string
		key  string
		val  bool
		want error
	}{
		{
			name: "valid key and value 'true'",
			key:  "KEY_1",
			val:  true,
			want: nil,
		},
		{
			name: "valid key and value 'false'",
			key:  "KEY_1",
			val:  false,
			want: nil,
		},
		{
			name: "empty key, valid value 'true'",
			key:  "",
			val:  true,
			want: ErrInvalidKey,
		},
		{
			name: "empty key, valid value 'false'",
			key:  "",
			val:  false,
			want: ErrInvalidKey,
		},
		{
			name: "quote in key, valid value 'true'",
			key:  `KEY_1"`,
			val:  true,
			want: ErrInvalidKey,
		},
		{
			name: "quote in key, valid value 'false'",
			key:  `KEY_1"`,
			val:  false,
			want: ErrInvalidKey,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)
		os.Clearenv()
		gotErr := SetBool(test.key, test.val)

		if !errors.Is(gotErr, test.want) {
			t.Fatalf("got %v, want %v", gotErr, test.want)
		}
	}
}

func TestSetFloat32(t *testing.T) {
	tests := []struct {
		name    string
		key     string
		val     float32
		envVal  string
		wantErr error
	}{
		{
			name:    "valid key and value ",
			key:     "KEY_1",
			val:     1.55,
			envVal:  "1.55",
			wantErr: nil,
		},
		{
			name:    "empty key, valid value",
			key:     "",
			val:     1.55,
			envVal:  "",
			wantErr: ErrInvalidKey,
		},
		{
			name:    "quote in key, valid value ",
			key:     `KEY_1"`,
			val:     1.55,
			envVal:  "",
			wantErr: ErrInvalidKey,
		},
	}

	for _, test := range tests {
		t.Logf("case %q", test.name)
		os.Clearenv()
		gotErr := SetFloat32(test.key, test.val)

		if !errors.Is(gotErr, test.wantErr) {
			t.Fatalf("got %f, want %f", gotErr, test.wantErr)
		}

		val, ok := os.LookupEnv(test.key)
		if !ok && test.envVal != "" {
			t.Fatalf("the environment variable was not to be found")
		}

		if val != test.envVal {
			t.Fatalf("want %q, got %q", test.envVal, val)
		}
	}
}

func setupEnv(envs map[string]string, t *testing.T) {
	os.Clearenv()
	for key, val := range envs {
		t.Setenv(key, val)
	}
}
