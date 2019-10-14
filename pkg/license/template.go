//go:generate enumer -type=Template -output=template_type.go -yaml -trimprefix=Template -linecomment

package license

import "fmt"

// Template is the License Identifier
type Template uint8

const (
	// ZeroBSD - BSD Zero Clause License
	ZeroBSD Template = iota // 0BSD

	// AFLv3 - Academic Free License v3.0
	AFLv3 // AFLv3

	// AGPLv3 - GNU AFFERO GENERAL PUBLIC LICENSE version 3.0
	AGPLv3 // AGPLv3

	// APACHEv2 - Apache License Version 2.0
	APACHEv2 // APACHEv2

	// ARTISTICv2 - The Artistic License 2.0
	ARTISTICv2 // ARTISTICv2

	// BSD2Clause - BSD 2-Clause License
	BSD2Clause // BSD2-CLAUSE

	// BSD3Clause - BSD 3-Clause License
	BSD3Clause // BSD3-CLAUSE

	// BSD3ClauseClear - The Clear BSD License
	BSD3ClauseClear // BSD3-CLEAR-CLEAR

	// BSL - Boost Software License
	BSL // BSL

	// CCv4 - Creative Commons Attribution 4.0 International
	CCv4 // CCv4

	// CCSAv4 - Creative Commons Attribution Share Alike 4.0 International
	CCSAv4 // CCSAv4

	// CC0v1 - Creative Commons Zero v1.0 Universal
	CC0v1 // CC0v1

	// CECILLv21 - CeCILL Free Software License Agreement v2.1
	CECILLv21 // CECILLv21

	// ECLv2 - Educational Community License v2.0
	ECLv2 // ECLv2

	// EPLv1 - Eclipse Public License 1.0
	EPLv1 // EPLv1

	// EPLv2 - Eclipse Public License 2.0
	EPLv2 // EPLv2

	// EUPLv11 - European Union Public License 1.1
	EUPLv11 // EUPLv11

	// EUPLv12 - European Union Public License 1.2
	EUPLv12 // EUPLv12

	// GPLv3 - GNU GENERAL PUBLIC LICENSE Version 3
	GPLv3 // GPLv3

	// ISC - ISC License
	ISC // ISC

	// LGPLv21 - GNU Lesser General Public License v2.1
	LGPLv21 // LGPLv21

	// LGPLv3 - GNU Lesser General Public License v3.0
	LGPLv3 // LGPLv3

	// LPPLv13 - LaTeX Project Public License v1.3c
	LPPLv13 // LPPLv13

	// MIT - MIT License
	MIT // MIT

	// MPLv2 - Mozilla Public License 2.0
	MPLv2 // MPLv2

	// MSPL - Microsoft Public License
	MSPL // MSPL

	// MSRL - Microsoft Reciprocal License
	MSRL // MSRL

	// NCSA - University of Illinois/NCSA Open Source License
	NCSA // NCSA

	// ODBLv1 - ODC Open Database License v1.0
	ODBLv1 // ODBLv1

	// OFLv11 - SIL Open Font License 1.1
	OFLv11 // OFLv11

	// OSLv3 - Open Software License 3.0
	OSLv3 // OSLv3

	// TODO: POSTGRESQL - Requires Template Fix

	// UNLICENSE - The Unlicense
	UNLICENSE // UNLICENSE

	// UPLv1 - Universal Permissive License v1.0
	UPLv1 // UPLv1

	// WTFPL - Do What The F*ck You Want To Public License
	WTFPL // WTFPL

	// ZLIB - ZLIB License
	ZLIB // ZLIB
)

var (
	// TemplateDescription defines the descriptions of an embedded License
	TemplateDescription map[Template]string = make(map[Template]string, 0)
)

func init() {
	TemplateDescription[ZeroBSD] = "BSD Zero Clause License"
	TemplateDescription[AFLv3] = "Academic Free License v3.0"
	TemplateDescription[AGPLv3] = "GNU AFFERO GENERAL PUBLIC LICENSE version 3.0"
	TemplateDescription[APACHEv2] = "Apache License Version 2.0"
	TemplateDescription[ARTISTICv2] = "The Artistic License 2.0"
	TemplateDescription[BSD2Clause] = "BSD 2-Clause License"
	TemplateDescription[BSD3Clause] = "BSD 3-Clause License"
	TemplateDescription[BSD3ClauseClear] = "The Clear BSD License"
	TemplateDescription[BSL] = "Boost Software License"
	TemplateDescription[CCv4] = "Creative Commons Attribution 4.0 International"
	TemplateDescription[CCSAv4] = "Creative Commons Attribution Share Alike 4.0 International"
	TemplateDescription[CC0v1] = "Creative Commons Zero v1.0 Universal"
	TemplateDescription[CECILLv21] = "CeCILL Free Software License Agreement v2.1"
	TemplateDescription[ECLv2] = "Educational Community License v2.0"
	TemplateDescription[EPLv1] = "Eclipse Public License 1.0"
	TemplateDescription[EPLv2] = "Eclipse Public License 2.0"
	TemplateDescription[EUPLv11] = "European Union Public License 1.1"
	TemplateDescription[EUPLv12] = "European Union Public License 1.2"
	TemplateDescription[GPLv3] = "GNU GENERAL PUBLIC LICENSE Version 3"
	TemplateDescription[ISC] = "ISC License"
	TemplateDescription[LGPLv21] = "GNU Lesser General Public License v2.1"
	TemplateDescription[LGPLv3] = "GNU Lesser General Public License v3.0"
	TemplateDescription[LPPLv13] = "LaTeX Project Public License v1.3c"
	TemplateDescription[MIT] = "MIT License"
	TemplateDescription[MPLv2] = "Mozilla Public License 2.0"
	TemplateDescription[MSPL] = "Microsoft Public License"
	TemplateDescription[MSRL] = "Microsoft Reciprocal License"
	TemplateDescription[NCSA] = "University of Illinois/NCSA Open Source License"
	TemplateDescription[ODBLv1] = "ODC Open Database License v1.0"
	TemplateDescription[OFLv11] = "SIL Open Font License 1.1"
	TemplateDescription[OSLv3] = "Open Software License 3.0"
	TemplateDescription[UNLICENSE] = "The Unlicense"
	TemplateDescription[UPLv1] = "Universal Permissive License v1.0"
	TemplateDescription[WTFPL] = "Do What The F*ck You Want To Public License"
	TemplateDescription[ZLIB] = "ZLIB License"
}

// List all embedded licenses
func List() []string {
	var l []string

	for _, v := range TemplateValues() {
		l = append(l, v.String())
	}

	return l
}

// ListWithDescription will list all the embedded licenses with their fullname
func ListWithDescription() []string {
	var l []string

	for _, v := range TemplateValues() {
		l = append(l, fmt.Sprintf("%-16s - %s", v.String(), TemplateDescription[v]))
	}

	return l
}
