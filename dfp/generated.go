package dfp

type Package struct {
	SchemaLocation            string      `xml:"schemaLocation,attr"`
	Xs                        string      `xml:"xs,attr"`
	SchemaVersion             float64     `xml:"schemaVersion,attr"`
	NoNamespaceSchemaLocation string      `xml:"noNamespaceSchemaLocation,attr"`
	Keywords                  Keywords    `xml:"keywords"`
	Devices                   Devices     `xml:"devices"`
	Conditions                Conditions  `xml:"conditions"`
	Description               Description `xml:"description"`
	Releases                  Releases    `xml:"releases"`
	Vendor                    Vendor      `xml:"vendor"`
	Url                       Url         `xml:"url"`
	Name                      Name        `xml:"name"`
	Components                Components  `xml:"components"`
}

type Keywords struct {
	Keywords []Keyword `xml:"keyword"`
}

type Keyword struct {
	Content string `xml:",chardata"`
}

type Devices struct {
	Family Family `xml:"family"`
}

type Family struct {
	Dfamily     string      `xml:"Dfamily,attr"`
	Dvendor     string      `xml:"Dvendor,attr"`
	Environment Environment `xml:"environment"`
	Devices     []Device    `xml:"device"`
}

type Device struct {
	Dname        string        `xml:"Dname,attr"`
	Books        []Book        `xml:"book"`
	Processor    Processor     `xml:"processor"`
	Compile      Compile       `xml:"compile"`
	Debug        Debug         `xml:"debug"`
	Memorys      []Memory      `xml:"memory"`
	Algorithms   []Algorithm   `xml:"algorithm"`
	Environments []Environment `xml:"environment"`
	Description  Description   `xml:"description"`
}

type Book struct {
	Name  string `xml:"name,attr"`
	Title string `xml:"title,attr"`
}

type Processor struct {
	Dcore   string `xml:"Dcore,attr"`
	Dendian string `xml:"Dendian,attr"`
	Dmpu    string `xml:"Dmpu,attr"`
	Dfpu    string `xml:"Dfpu,attr"`
}

type Compile struct {
	Header string `xml:"header,attr"`
	Define string `xml:"define,attr"`
}

type Debug struct {
	Svd string `xml:"svd,attr"`
}

type Algorithm struct {
	Start   string `xml:"start,attr"`
	Size    string `xml:"size,attr"`
	Default int64  `xml:"default,attr"`
	Name    string `xml:"name,attr"`
}

type Environment struct {
	Name      string    `xml:"name,attr"`
	Extension Extension `xml:"extension"`
}

type Extension struct {
	At            string         `xml:"at,attr"`
	SchemaVersion float64        `xml:"schemaVersion,attr"`
	Mchp          string         `xml:"mchp,attr"`
	Atdf          Atdf           `xml:"atdf"`
	Memorys       []Memory       `xml:"memory"`
	Tools         []Tool         `xml:"tool"`
	Projects      []Project      `xml:"project"`
	Prerequisites []Prerequisite `xml:"prerequisite"`
	Properties    []Property     `xml:"property"`
	Variants      []Variant      `xml:"variant"`
	Interface     Interface      `xml:"interface"`
	Pic           Pic            `xml:"pic"`
}

type Atdf struct {
	Name string `xml:"name,attr"`
}

type Memory struct {
	Exec         bool   `xml:"exec,attr"`
	AddressSpace string `xml:"address-space,attr"`
	Name         string `xml:"name,attr"`
	Start        string `xml:"start,attr"`
	Size         string `xml:"size,attr"`
	Type         string `xml:"type,attr"`
	Pagesize     string `xml:"pagesize,attr"`
	Rw           string `xml:"rw,attr"`
	Id           string `xml:"id,attr"`
	Default      int64  `xml:"default,attr"`
	Startup      int64  `xml:"startup,attr"`
}

type Tool struct {
	Id string `xml:"id,attr"`
}

type Project struct {
	Name       string      `xml:"name,attr"`
	Components []Component `xml:"component"`
}

type Template struct {
	Select string `xml:"select,attr"`
}

type Prerequisite struct {
	Context   string `xml:"context,attr"`
	Tcompiler string `xml:"Tcompiler,attr"`
	Component string `xml:"component,attr"`
	Version   string `xml:"version,attr"`
}

type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

type Variant struct {
	Vccmin    float64 `xml:"vccmin,attr"`
	Vccmax    float64 `xml:"vccmax,attr"`
	Tempmin   int64   `xml:"tempmin,attr"`
	Tempmax   int64   `xml:"tempmax,attr"`
	Ordercode string  `xml:"ordercode,attr"`
}

type Interface struct {
	Type string `xml:"type,attr"`
	Name string `xml:"name,attr"`
}

type Pic struct {
	Name string `xml:"name,attr"`
}

type Conditions struct {
	Conditions []Condition `xml:"condition"`
}

type Condition struct {
	Id       string    `xml:"id,attr"`
	Accepts  []Accept  `xml:"accept"`
	Requires []Require `xml:"require"`
}

type Accept struct {
	Tcompiler string `xml:"Tcompiler,attr"`
	Toutput   string `xml:"Toutput,attr"`
}

type Require struct {
	Dvendor string `xml:"Dvendor,attr"`
	Dname   string `xml:"Dname,attr"`
	Cclass  string `xml:"Cclass,attr"`
	Cgroup  string `xml:"Cgroup,attr"`
}

type Releases struct {
	Releases []Release `xml:"release"`
}

type Release struct {
	Version string `xml:"version,attr"`
	Date    string `xml:"date,attr"`
	Content string `xml:",chardata"`
}

type Vendor struct {
	Content string `xml:",chardata"`
}

type Url struct {
	Content string `xml:",chardata"`
}

type Name struct {
	Content string `xml:",chardata"`
}

type Components struct {
	Components []Component `xml:"component"`
}

type Component struct {
	Cgroup      string      `xml:"Cgroup,attr"`
	Cversion    string      `xml:"Cversion,attr"`
	Condition   string      `xml:"condition,attr"`
	Cvendor     string      `xml:"Cvendor,attr"`
	Cclass      string      `xml:"Cclass,attr"`
	Description Description `xml:"description"`
	Files       Files       `xml:"files"`
	Template    Template    `xml:"template"`
}

type Description struct {
	Content string `xml:",chardata"`
}

type Files struct {
	Files []File `xml:"file"`
}

type File struct {
	Condition string `xml:"condition,attr"`
	Category  string `xml:"category,attr"`
	Name      string `xml:"name,attr"`
	Attr      string `xml:"attr,attr"`
	Select    string `xml:"select,attr"`
}
