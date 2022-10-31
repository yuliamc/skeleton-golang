package enum

import (
	"modalrakyat/skeleton-golang/pkg/utils/constant"
	"strings"
)

// Education is connected with Loan model struct #Registration.Education
type Education int64

//Scan for converting byte to string for fetching/read
func (s *Education) Scan(value interface{}) error {
	val := value.([]uint8)
	key := string(val)
	for i, v := range educationKey {
		if v == key {
			*s = i
		}
	}
	return nil
}

func NewEducation(value string) Education {
	for i, v := range educationKey {
		if v == value {
			return i
		}
	}
	panic("enum not found")
}

const (
	REGISTRATION_EDUCATION_TIDAK_SEKOLAH Education = iota + 1
	REGISTRATION_EDUCATION_SD
	REGISTRATION_EDUCATION_SMP
	REGISTRATION_EDUCATION_SMA
	REGISTRATION_EDUCATION_DIPLOMA
	REGISTRATION_EDUCATION_SARJANA
	REGISTRATION_EDUCATION_MAGISTER
	REGISTRATION_EDUCATION_DOKTOR
)

var educationKey = map[Education]string{
	REGISTRATION_EDUCATION_TIDAK_SEKOLAH: "Tidak Sekolah",
	REGISTRATION_EDUCATION_SD:            "SD",
	REGISTRATION_EDUCATION_SMP:           "SMP",
	REGISTRATION_EDUCATION_SMA:           "SMA",
	REGISTRATION_EDUCATION_DIPLOMA:       "Diploma",
	REGISTRATION_EDUCATION_SARJANA:       "Sarjana",
	REGISTRATION_EDUCATION_MAGISTER:      "Magister",
	REGISTRATION_EDUCATION_DOKTOR:        "Doktor",
}

//String for stringify Education
func (s Education) String() string {
	return educationKey[s]
}

// Construct array of key value pairs of jobs
func GetEducationKeyValuePairs() []constant.KeyValue {
	arr := []constant.KeyValue{}
	for _, v := range educationKey {
		arr = append(arr, constant.KeyValue{
			Key:   v,                // act as key
			Value: strings.Title(v), // act as label
		})
	}
	return arr
}
