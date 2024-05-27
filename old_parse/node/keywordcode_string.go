// Code generated by "stringer -type KeywordCode keyword.go"; DO NOT EDIT.

package node

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[KeywordUnspecified-0]
	_ = x[KeywordAll-1]
	_ = x[KeywordAnd-2]
	_ = x[KeywordAny-3]
	_ = x[KeywordArray-4]
	_ = x[KeywordAs-5]
	_ = x[KeywordAsc-6]
	_ = x[KeywordAssertRowsModified-7]
	_ = x[KeywordAt-8]
	_ = x[KeywordBetween-9]
	_ = x[KeywordBy-10]
	_ = x[KeywordCase-11]
	_ = x[KeywordCast-12]
	_ = x[KeywordCollate-13]
	_ = x[KeywordContains-14]
	_ = x[KeywordCreate-15]
	_ = x[KeywordCross-16]
	_ = x[KeywordCube-17]
	_ = x[KeywordCurrent-18]
	_ = x[KeywordDefault-19]
	_ = x[KeywordDefine-20]
	_ = x[KeywordDesc-21]
	_ = x[KeywordDistinct-22]
	_ = x[KeywordElse-23]
	_ = x[KeywordEnd-24]
	_ = x[KeywordEnum-25]
	_ = x[KeywordEscape-26]
	_ = x[KeywordExcept-27]
	_ = x[KeywordExclude-28]
	_ = x[KeywordExists-29]
	_ = x[KeywordExtract-30]
	_ = x[KeywordFalse-31]
	_ = x[KeywordFetch-32]
	_ = x[KeywordFollowing-33]
	_ = x[KeywordFor-34]
	_ = x[KeywordFrom-35]
	_ = x[KeywordFull-36]
	_ = x[KeywordGroup-37]
	_ = x[KeywordGrouping-38]
	_ = x[KeywordGroups-39]
	_ = x[KeywordHash-40]
	_ = x[KeywordHaving-41]
	_ = x[KeywordIf-42]
	_ = x[KeywordIgnore-43]
	_ = x[KeywordIn-44]
	_ = x[KeywordInner-45]
	_ = x[KeywordIntersect-46]
	_ = x[KeywordInterval-47]
	_ = x[KeywordInto-48]
	_ = x[KeywordIs-49]
	_ = x[KeywordJoin-50]
	_ = x[KeywordLateral-51]
	_ = x[KeywordLeft-52]
	_ = x[KeywordLike-53]
	_ = x[KeywordLimit-54]
	_ = x[KeywordLookup-55]
	_ = x[KeywordMerge-56]
	_ = x[KeywordNatural-57]
	_ = x[KeywordNew-58]
	_ = x[KeywordNo-59]
	_ = x[KeywordNot-60]
	_ = x[KeywordNull-61]
	_ = x[KeywordNulls-62]
	_ = x[KeywordOf-63]
	_ = x[KeywordOn-64]
	_ = x[KeywordOr-65]
	_ = x[KeywordOrder-66]
	_ = x[KeywordOuter-67]
	_ = x[KeywordOver-68]
	_ = x[KeywordPartition-69]
	_ = x[KeywordPreceding-70]
	_ = x[KeywordProto-71]
	_ = x[KeywordRange-72]
	_ = x[KeywordRecursive-73]
	_ = x[KeywordRespect-74]
	_ = x[KeywordRight-75]
	_ = x[KeywordRollup-76]
	_ = x[KeywordRows-77]
	_ = x[KeywordSelect-78]
	_ = x[KeywordSet-79]
	_ = x[KeywordSome-80]
	_ = x[KeywordStruct-81]
	_ = x[KeywordTablesample-82]
	_ = x[KeywordThen-83]
	_ = x[KeywordTo-84]
	_ = x[KeywordTreat-85]
	_ = x[KeywordTrue-86]
	_ = x[KeywordUnbound-87]
	_ = x[KeywordUnion-88]
	_ = x[KeywordUnnest-89]
	_ = x[KeywordUsing-90]
	_ = x[KeywordWhen-91]
	_ = x[KeywordWhere-92]
	_ = x[KeywordWindow-93]
	_ = x[KeywordWith-94]
	_ = x[KeywordWithin-95]
}

const _KeywordCode_name = "KeywordUnspecifiedKeywordAllKeywordAndKeywordAnyKeywordArrayKeywordAsKeywordAscKeywordAssertRowsModifiedKeywordAtKeywordBetweenKeywordByKeywordCaseKeywordCastKeywordCollateKeywordContainsKeywordCreateKeywordCrossKeywordCubeKeywordCurrentKeywordDefaultKeywordDefineKeywordDescKeywordDistinctKeywordElseKeywordEndKeywordEnumKeywordEscapeKeywordExceptKeywordExcludeKeywordExistsKeywordExtractKeywordFalseKeywordFetchKeywordFollowingKeywordForKeywordFromKeywordFullKeywordGroupKeywordGroupingKeywordGroupsKeywordHashKeywordHavingKeywordIfKeywordIgnoreKeywordInKeywordInnerKeywordIntersectKeywordIntervalKeywordIntoKeywordIsKeywordJoinKeywordLateralKeywordLeftKeywordLikeKeywordLimitKeywordLookupKeywordMergeKeywordNaturalKeywordNewKeywordNoKeywordNotKeywordNullKeywordNullsKeywordOfKeywordOnKeywordOrKeywordOrderKeywordOuterKeywordOverKeywordPartitionKeywordPrecedingKeywordProtoKeywordRangeKeywordRecursiveKeywordRespectKeywordRightKeywordRollupKeywordRowsKeywordSelectKeywordSetKeywordSomeKeywordStructKeywordTablesampleKeywordThenKeywordToKeywordTreatKeywordTrueKeywordUnboundKeywordUnionKeywordUnnestKeywordUsingKeywordWhenKeywordWhereKeywordWindowKeywordWithKeywordWithin"

var _KeywordCode_index = [...]uint16{0, 18, 28, 38, 48, 60, 69, 79, 104, 113, 127, 136, 147, 158, 172, 187, 200, 212, 223, 237, 251, 264, 275, 290, 301, 311, 322, 335, 348, 362, 375, 389, 401, 413, 429, 439, 450, 461, 473, 488, 501, 512, 525, 534, 547, 556, 568, 584, 599, 610, 619, 630, 644, 655, 666, 678, 691, 703, 717, 727, 736, 746, 757, 769, 778, 787, 796, 808, 820, 831, 847, 863, 875, 887, 903, 917, 929, 942, 953, 966, 976, 987, 1000, 1018, 1029, 1038, 1050, 1061, 1075, 1087, 1100, 1112, 1123, 1135, 1148, 1159, 1172}

func (i KeywordCode) String() string {
	if i < 0 || i >= KeywordCode(len(_KeywordCode_index)-1) {
		return "KeywordCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _KeywordCode_name[_KeywordCode_index[i]:_KeywordCode_index[i+1]]
}
