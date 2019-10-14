//go:generate enumer -type=SPDX -output=spdx_type.go -yaml -trimprefix=SPDX -linecomment

package license

// SPDX is the License Identifier
type SPDX uint8

const (
	// ZeroBSD - BSD Zero Clause License
	ZeroBSD SPDX = iota // 0BSD

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
