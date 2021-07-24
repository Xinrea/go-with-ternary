// Code generated by "stringer -type=Op -trimprefix=O node.go"; DO NOT EDIT.

package ir

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[OXXX-0]
	_ = x[ONAME-1]
	_ = x[ONONAME-2]
	_ = x[OTYPE-3]
	_ = x[OPACK-4]
	_ = x[OLITERAL-5]
	_ = x[ONIL-6]
	_ = x[OADD-7]
	_ = x[OSUB-8]
	_ = x[OOR-9]
	_ = x[OXOR-10]
	_ = x[OADDSTR-11]
	_ = x[OADDR-12]
	_ = x[OANDAND-13]
	_ = x[OAPPEND-14]
	_ = x[OBYTES2STR-15]
	_ = x[OBYTES2STRTMP-16]
	_ = x[ORUNES2STR-17]
	_ = x[OSTR2BYTES-18]
	_ = x[OSTR2BYTESTMP-19]
	_ = x[OSTR2RUNES-20]
	_ = x[OSLICE2ARRPTR-21]
	_ = x[OAS-22]
	_ = x[OAS2-23]
	_ = x[OAS2DOTTYPE-24]
	_ = x[OAS2FUNC-25]
	_ = x[OAS2MAPR-26]
	_ = x[OAS2RECV-27]
	_ = x[OASOP-28]
	_ = x[OCALL-29]
	_ = x[OCALLFUNC-30]
	_ = x[OCALLMETH-31]
	_ = x[OCALLINTER-32]
	_ = x[OCAP-33]
	_ = x[OCLOSE-34]
	_ = x[OCLOSURE-35]
	_ = x[OCOMPLIT-36]
	_ = x[OMAPLIT-37]
	_ = x[OSTRUCTLIT-38]
	_ = x[OARRAYLIT-39]
	_ = x[OSLICELIT-40]
	_ = x[OPTRLIT-41]
	_ = x[OCONV-42]
	_ = x[OCONVIFACE-43]
	_ = x[OCONVIDATA-44]
	_ = x[OCONVNOP-45]
	_ = x[OCOPY-46]
	_ = x[ODCL-47]
	_ = x[ODCLFUNC-48]
	_ = x[ODCLCONST-49]
	_ = x[ODCLTYPE-50]
	_ = x[ODELETE-51]
	_ = x[ODOT-52]
	_ = x[ODOTPTR-53]
	_ = x[ODOTMETH-54]
	_ = x[ODOTINTER-55]
	_ = x[OXDOT-56]
	_ = x[ODOTTYPE-57]
	_ = x[ODOTTYPE2-58]
	_ = x[OEQ-59]
	_ = x[ONE-60]
	_ = x[OLT-61]
	_ = x[OLE-62]
	_ = x[OGE-63]
	_ = x[OGT-64]
	_ = x[ODEREF-65]
	_ = x[OINDEX-66]
	_ = x[OINDEXMAP-67]
	_ = x[OKEY-68]
	_ = x[OSTRUCTKEY-69]
	_ = x[OLEN-70]
	_ = x[OMAKE-71]
	_ = x[OMAKECHAN-72]
	_ = x[OMAKEMAP-73]
	_ = x[OMAKESLICE-74]
	_ = x[OMAKESLICECOPY-75]
	_ = x[OMUL-76]
	_ = x[ODIV-77]
	_ = x[OMOD-78]
	_ = x[OLSH-79]
	_ = x[ORSH-80]
	_ = x[OAND-81]
	_ = x[OANDNOT-82]
	_ = x[ONEW-83]
	_ = x[ONOT-84]
	_ = x[OBITNOT-85]
	_ = x[OPLUS-86]
	_ = x[ONEG-87]
	_ = x[OOROR-88]
	_ = x[OPANIC-89]
	_ = x[OPRINT-90]
	_ = x[OPRINTN-91]
	_ = x[OPAREN-92]
	_ = x[OSEND-93]
	_ = x[OSLICE-94]
	_ = x[OSLICEARR-95]
	_ = x[OSLICESTR-96]
	_ = x[OSLICE3-97]
	_ = x[OSLICE3ARR-98]
	_ = x[OSLICEHEADER-99]
	_ = x[ORECOVER-100]
	_ = x[ORECOVERFP-101]
	_ = x[ORECV-102]
	_ = x[ORUNESTR-103]
	_ = x[OSELRECV2-104]
	_ = x[OIOTA-105]
	_ = x[OREAL-106]
	_ = x[OIMAG-107]
	_ = x[OCOMPLEX-108]
	_ = x[OALIGNOF-109]
	_ = x[OOFFSETOF-110]
	_ = x[OSIZEOF-111]
	_ = x[OUNSAFEADD-112]
	_ = x[OUNSAFESLICE-113]
	_ = x[OMETHEXPR-114]
	_ = x[OMETHVALUE-115]
	_ = x[OBLOCK-116]
	_ = x[OBREAK-117]
	_ = x[OCASE-118]
	_ = x[OCONTINUE-119]
	_ = x[ODEFER-120]
	_ = x[OFALL-121]
	_ = x[OFOR-122]
	_ = x[OFORUNTIL-123]
	_ = x[OGOTO-124]
	_ = x[OIF-125]
	_ = x[OLABEL-126]
	_ = x[OGO-127]
	_ = x[ORANGE-128]
	_ = x[ORETURN-129]
	_ = x[OSELECT-130]
	_ = x[OSWITCH-131]
	_ = x[OTYPESW-132]
	_ = x[OFUNCINST-133]
	_ = x[OTCHAN-134]
	_ = x[OTMAP-135]
	_ = x[OTSTRUCT-136]
	_ = x[OTINTER-137]
	_ = x[OTFUNC-138]
	_ = x[OTARRAY-139]
	_ = x[OTSLICE-140]
	_ = x[OINLCALL-141]
	_ = x[OEFACE-142]
	_ = x[OITAB-143]
	_ = x[OIDATA-144]
	_ = x[OSPTR-145]
	_ = x[OCFUNC-146]
	_ = x[OCHECKNIL-147]
	_ = x[OVARDEF-148]
	_ = x[OVARKILL-149]
	_ = x[OVARLIVE-150]
	_ = x[ORESULT-151]
	_ = x[OINLMARK-152]
	_ = x[OLINKSYMOFFSET-153]
	_ = x[OTAILCALL-154]
	_ = x[OGETG-155]
	_ = x[OGETCALLERPC-156]
	_ = x[OGETCALLERSP-157]
	_ = x[OEND-158]
}

const _Op_name = "XXXNAMENONAMETYPEPACKLITERALNILADDSUBORXORADDSTRADDRANDANDAPPENDBYTES2STRBYTES2STRTMPRUNES2STRSTR2BYTESSTR2BYTESTMPSTR2RUNESSLICE2ARRPTRASAS2AS2DOTTYPEAS2FUNCAS2MAPRAS2RECVASOPCALLCALLFUNCCALLMETHCALLINTERCAPCLOSECLOSURECOMPLITMAPLITSTRUCTLITARRAYLITSLICELITPTRLITCONVCONVIFACECONVIDATACONVNOPCOPYDCLDCLFUNCDCLCONSTDCLTYPEDELETEDOTDOTPTRDOTMETHDOTINTERXDOTDOTTYPEDOTTYPE2EQNELTLEGEGTDEREFINDEXINDEXMAPKEYSTRUCTKEYLENMAKEMAKECHANMAKEMAPMAKESLICEMAKESLICECOPYMULDIVMODLSHRSHANDANDNOTNEWNOTBITNOTPLUSNEGORORPANICPRINTPRINTNPARENSENDSLICESLICEARRSLICESTRSLICE3SLICE3ARRSLICEHEADERRECOVERRECOVERFPRECVRUNESTRSELRECV2IOTAREALIMAGCOMPLEXALIGNOFOFFSETOFSIZEOFUNSAFEADDUNSAFESLICEMETHEXPRMETHVALUEBLOCKBREAKCASECONTINUEDEFERFALLFORFORUNTILGOTOIFLABELGORANGERETURNSELECTSWITCHTYPESWFUNCINSTTCHANTMAPTSTRUCTTINTERTFUNCTARRAYTSLICEINLCALLEFACEITABIDATASPTRCFUNCCHECKNILVARDEFVARKILLVARLIVERESULTINLMARKLINKSYMOFFSETTAILCALLGETGGETCALLERPCGETCALLERSPEND"

var _Op_index = [...]uint16{0, 3, 7, 13, 17, 21, 28, 31, 34, 37, 39, 42, 48, 52, 58, 64, 73, 85, 94, 103, 115, 124, 136, 138, 141, 151, 158, 165, 172, 176, 180, 188, 196, 205, 208, 213, 220, 227, 233, 242, 250, 258, 264, 268, 277, 286, 293, 297, 300, 307, 315, 322, 328, 331, 337, 344, 352, 356, 363, 371, 373, 375, 377, 379, 381, 383, 388, 393, 401, 404, 413, 416, 420, 428, 435, 444, 457, 460, 463, 466, 469, 472, 475, 481, 484, 487, 493, 497, 500, 504, 509, 514, 520, 525, 529, 534, 542, 550, 556, 565, 576, 583, 592, 596, 603, 611, 615, 619, 623, 630, 637, 645, 651, 660, 671, 679, 688, 693, 698, 702, 710, 715, 719, 722, 730, 734, 736, 741, 743, 748, 754, 760, 766, 772, 780, 785, 789, 796, 802, 807, 813, 819, 826, 831, 835, 840, 844, 849, 857, 863, 870, 877, 883, 890, 903, 911, 915, 926, 937, 940}

func (i Op) String() string {
	if i >= Op(len(_Op_index)-1) {
		return "Op(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Op_name[_Op_index[i]:_Op_index[i+1]]
}
