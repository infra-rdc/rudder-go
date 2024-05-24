package components

import (
	"encoding/json"
	"fmt"
)

type OsType string

const (
	OsTypeLinux   OsType = "linux"
	OsTypeWindows OsType = "windows"
	OsTypeSolaris OsType = "solaris"
	OsTypeAix     OsType = "aix"
	OsTypeFreebsd OsType = "freebsd"
	OsTypeUnknown OsType = "unknown"
)

func (e OsType) ToPointer() *OsType {
	return &e
}
func (e *OsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linux":
		fallthrough
	case "windows":
		fallthrough
	case "solaris":
		fallthrough
	case "aix":
		fallthrough
	case "freebsd":
		fallthrough
	case "unknown":
		*e = OsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for OsType: %v", v)
	}
}

// Name - For Linux, a distribution, for Windows, the commercial name
type Name string

const (
	NameDebian                 Name = "debian"
	NameKali                   Name = "kali"
	NameUbuntu                 Name = "ubuntu"
	NameRedhat                 Name = "redhat"
	NameCentos                 Name = "centos"
	NameFedora                 Name = "fedora"
	NameSuse                   Name = "suse"
	NameOracle                 Name = "oracle"
	NameScientific             Name = "scientific"
	NameSlackware              Name = "slackware"
	NameXp                     Name = "xp"
	NameVista                  Name = "vista"
	NameSeven                  Name = "seven"
	NameTen                    Name = "10"
	NameTwoThousand            Name = "2000"
	NameTwoThousandAndThree    Name = "2003"
	NameTwoThousandAndEightR2  Name = "2008 r2"
	NameTwoThousandAndTwelve   Name = "2012"
	NameTwoThousandAndTwelveR2 Name = "2012 r2"
	NameTwoThousandAndSixteen  Name = "2016"
)

func (e Name) ToPointer() *Name {
	return &e
}
func (e *Name) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "debian":
		fallthrough
	case "kali":
		fallthrough
	case "ubuntu":
		fallthrough
	case "redhat":
		fallthrough
	case "centos":
		fallthrough
	case "fedora":
		fallthrough
	case "suse":
		fallthrough
	case "oracle":
		fallthrough
	case "scientific":
		fallthrough
	case "slackware":
		fallthrough
	case "xp":
		fallthrough
	case "vista":
		fallthrough
	case "seven":
		fallthrough
	case "10":
		fallthrough
	case "2000":
		fallthrough
	case "2003":
		fallthrough
	case "2008 r2":
		fallthrough
	case "2012":
		fallthrough
	case "2012 r2":
		fallthrough
	case "2016":
		*e = Name(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Name: %v", v)
	}
}

type Os struct {
	Type OsType `json:"type"`
	// For Linux, a distribution, for Windows, the commercial name
	Name Name `json:"name"`
	// A string representation of the version
	Version string `json:"version"`
	// The long description name of the os
	FullName string `json:"fullName"`
	// a service pack informationnal string
	ServicePack *string `json:"servicePack,omitempty"`
}

func (o *Os) GetType() OsType {
	if o == nil {
		return OsType("")
	}
	return o.Type
}

func (o *Os) GetName() Name {
	if o == nil {
		return Name("")
	}
	return o.Name
}

func (o *Os) GetVersion() string {
	if o == nil {
		return ""
	}
	return o.Version
}

func (o *Os) GetFullName() string {
	if o == nil {
		return ""
	}
	return o.FullName
}

func (o *Os) GetServicePack() *string {
	if o == nil {
		return nil
	}
	return o.ServicePack
}
