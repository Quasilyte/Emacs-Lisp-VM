package bcode

// Opcode names are almost identical to Emacs Lisp bytecode specs.
// There are some differences in opcode suffixes:
// uint8  arg   => "B"
// uint16 arg   => "W"
// implicit arg => "0"..."5"
const (
	OpExt byte = 0 // Replaces "stack_ref+0"

	OpStackRef1 byte = 1
	OpStackRef2 byte = 2
	OpStackRef3 byte = 3
	OpStackRef4 byte = 4
	OpStackRef5 byte = 5
	OpStackRefB byte = 6
	OpStackRefW byte = 7

	// Issue#3
	OpVarRef0 byte = 010
	OpVarRef1      = OpVarRef0 + 1
	OpVarRef2      = OpVarRef0 + 2
	OpVarRef3      = OpVarRef0 + 3
	OpVarRef4      = OpVarRef0 + 4
	OpVarRef5      = OpVarRef0 + 5
	OpVarRefB      = OpVarRef0 + 6
	OpVarRefW      = OpVarRef0 + 7

	// Issue#3
	OpVarSet0 byte = 020
	OpVarSet1      = OpVarSet0 + 1
	OpVarSet2      = OpVarSet0 + 2
	OpVarSet3      = OpVarSet0 + 3
	OpVarSet4      = OpVarSet0 + 4
	OpVarSet5      = OpVarSet0 + 5
	OpVarSetB      = OpVarSet0 + 6
	OpVarSetW      = OpVarSet0 + 7

	// Issue#3
	OpVarBind0 byte = 030
	OpVarBind1      = OpVarBind0 + 1
	OpVarBind2      = OpVarBind0 + 2
	OpVarBind3      = OpVarBind0 + 3
	OpVarBind4      = OpVarBind0 + 4
	OpVarBind5      = OpVarBind0 + 5
	OpVarBindB      = OpVarBind0 + 6
	OpVarBindW      = OpVarBind0 + 7

	// Issue#5
	OpCall0 byte = 040
	OpCall1      = OpCall0 + 1
	OpCall2      = OpCall0 + 2
	OpCall3      = OpCall0 + 3
	OpCall4      = OpCall0 + 4
	OpCall5      = OpCall0 + 5
	OpCallB      = OpCall0 + 6
	OpCallW      = OpCall0 + 7

	// Issue#3
	OpUnbind0 byte = 050
	OpUnbind1      = OpUnbind0 + 1
	OpUnbind2      = OpUnbind0 + 2
	OpUnbind3      = OpUnbind0 + 3
	OpUnbind4      = OpUnbind0 + 4
	OpUnbind5      = OpUnbind0 + 5
	OpUnbindB      = OpUnbind0 + 6
	OpUnbindW      = OpUnbind0 + 7

	OpPopHandler            byte = 060 // Issue#7
	OpPushConditionCase     byte = 061 // Issue#7
	OpPushCatch             byte = 062 // Issue#7
	OpNth                   byte = 070
	OpSymbolp               byte = 071
	OpConsp                 byte = 072
	OpStringp               byte = 073
	OpListp                 byte = 074
	OpEq                    byte = 075
	OpMemq                  byte = 076
	OpNot                   byte = 077
	OpCar                   byte = 0100
	OpCdr                   byte = 0101
	OpCons                  byte = 0102
	OpList1                 byte = 0103
	OpList2                 byte = 0104
	OpList3                 byte = 0105
	OpList4                 byte = 0106
	OpLength                byte = 0107
	OpAref                  byte = 0110
	OpAset                  byte = 0111
	OpSet                   byte = 0114 // Issue#4
	OpFset                  byte = 0115 // Issue#5
	OpGet                   byte = 0116 // Issue#4
	OpSubstring             byte = 0117
	OpConcat2               byte = 0120
	OpConcat3               byte = 0121
	OpConcat4               byte = 0122
	OpSub1                  byte = 0123
	OpAdd1                  byte = 0124
	OpEqlsign               byte = 0125 // `=`
	OpGtr                   byte = 0126 // `>`
	OpLss                   byte = 0127 // `<`
	OpLeq                   byte = 0130 // `<=`
	OpGeq                   byte = 0131 // `>=`
	OpDiff                  byte = 0132 // `-`
	OpNegate                byte = 0133 // Negate number sign
	OpPlus                  byte = 0134
	OpMax                   byte = 0135
	OpMin                   byte = 0136
	OpMult                  byte = 0137 // `*`
	OpPoint                 byte = 0140 // Issue#4
	OpSaveCurrentBuffer     byte = 0141 // Issue#4
	OpGotoChar              byte = 0142 // Issue#4
	OpInsert                byte = 0143 // Issue#4
	OpPointMax              byte = 0144 // Issue#4
	OpPointMin              byte = 0145 // Issue#4
	OpCharAfter             byte = 0146 // Issue#4
	OpFollowingChar         byte = 0147 // Issue#4
	OpPrecedingChar         byte = 0150 // Issue#4
	OpCurrentColumn         byte = 0151 // Issue#4
	OpIndentTo              byte = 0152 // Issue#4
	OpEolp                  byte = 0154 // EOL; Issue#4
	OpEobp                  byte = 0155 // End of buffer; Issue#4
	OpBolp                  byte = 0156 // Beginning of line; Issue#4
	OpBobp                  byte = 0157 // Beginning of buffer; Issue#4
	OpCurrentBuffer         byte = 0160 // Issue#4
	OpSetBuffer             byte = 0161 // Issue#4
	OpSaveCurrentBuffer2    byte = 0162 // Issue#4
	OpInteractivep          byte = 0164 // Issue#5
	OpForwardChar           byte = 0165 // Issue#4
	OpForwardWord           byte = 0166 // Issue#4
	OpSkipCharsForward      byte = 0167 // Issue#4
	OpSkipCharsBackward     byte = 0170 // Issue#4
	OpForwardLine           byte = 0171 // Issue#4
	OpCharSyntax            byte = 0172 // Issue#6
	OpBufferSubstring       byte = 0173 // Issue#4
	OpDeleteRegion          byte = 0174 // Issue#4
	OpNarrowToRegion        byte = 0175 // Issue#4
	OpWiden                 byte = 0176 // Issue#4
	OpEndOfLine             byte = 0177 // Issue#4
	OpConstantW             byte = 0201
	OpGotoW                 byte = 0202
	OpGotoIfNilW            byte = 0203
	OpGotoIfNonNilW         byte = 0204
	OpGotoIfNilElsePopW     byte = 0205
	OpGotoIfNonNilElsePopW  byte = 0206
	OpReturn                byte = 0207
	OpDiscard               byte = 0210
	OpDup                   byte = 0211
	OpSaveExcursion         byte = 0212 // Issue#4
	OpSaveWindowExcursion   byte = 0213 // Issue#4
	OpSaveRestriction       byte = 0214 // Issue#4
	OpCatch                 byte = 0215 // Issue#7
	OpUnwindProtect         byte = 0216 // Issue#7
	OpConditionCase         byte = 0217 // Issue#7
	OpTempOutputBufferSetup byte = 0220 // Issue#4
	OpTempOutputBufferShow  byte = 0221 // Issue#4
	OpUnbindAll             byte = 0222 // Issue#3
	OpSetMarker             byte = 0223 // Issue#4
	OpMatchBeginning        byte = 0224 // Issue#4
	OpMatchEnd              byte = 0225 // Issue#4
	OpUpcase                byte = 0226
	OpDowncase              byte = 0227
	OpStringEqlsign         byte = 0230 // `string=`
	OpStringLss             byte = 0231 // `string<`
	OpEqual                 byte = 0232
	OpNthCdr                byte = 0233
	OpElt                   byte = 0234
	OpMember                byte = 0235
	OpAssq                  byte = 0236
	OpNreverse              byte = 0237
	OpSetcar                byte = 0240
	OpSetcdr                byte = 0241
	OpCarSafe               byte = 0242
	OpCdrSafe               byte = 0243
	OpNconc                 byte = 0244
	OpQuo                   byte = 0245 // `/`
	OpRem                   byte = 0246
	OpNumberp               byte = 0247
	OpIntegerp              byte = 0250
	OpRgotoB                byte = 0252 // Issue#8
	OpRgotoIfNilB           byte = 0253 // Issue#8
	OpRgotoIfNonNilB        byte = 0254 // Issue#8
	OpRgotoIfNilElsePopB    byte = 0255 // Issue#8
	OpRgotoIfNonNilElsePopB byte = 0256 // Issue#8
	OpListB                 byte = 0257
	OpConcatB               byte = 0260
	OpInsertB               byte = 0261
	OpStackSetB             byte = 0262
	OpStackSetW             byte = 0263
	OpDiscardB              byte = 0266

	OpConstant0  byte = 0300
	OpConstant1       = OpConstant0 + 1
	OpConstant2       = OpConstant0 + 2
	OpConstant3       = OpConstant0 + 3
	OpConstant4       = OpConstant0 + 4
	OpConstant5       = OpConstant0 + 5
	OpConstant6       = OpConstant0 + 6
	OpConstant7       = OpConstant0 + 7
	OpConstant8       = OpConstant0 + 8
	OpConstant9       = OpConstant0 + 9
	OpConstant10      = OpConstant0 + 10
	OpConstant11      = OpConstant0 + 11
	OpConstant12      = OpConstant0 + 12
	OpConstant13      = OpConstant0 + 13
	OpConstant14      = OpConstant0 + 14
	OpConstant15      = OpConstant0 + 15
	OpConstant16      = OpConstant0 + 16
	OpConstant17      = OpConstant0 + 17
	OpConstant18      = OpConstant0 + 18
	OpConstant19      = OpConstant0 + 19
	OpConstant20      = OpConstant0 + 20
	OpConstant21      = OpConstant0 + 21
	OpConstant22      = OpConstant0 + 22
	OpConstant23      = OpConstant0 + 23
	OpConstant24      = OpConstant0 + 24
	OpConstant25      = OpConstant0 + 25
	OpConstant26      = OpConstant0 + 26
	OpConstant27      = OpConstant0 + 27
	OpConstant28      = OpConstant0 + 28
	OpConstant29      = OpConstant0 + 29
	OpConstant30      = OpConstant0 + 30
	OpConstant31      = OpConstant0 + 31
	OpConstant32      = OpConstant0 + 32
	OpConstant33      = OpConstant0 + 33
	OpConstant34      = OpConstant0 + 34
	OpConstant35      = OpConstant0 + 35
	OpConstant36      = OpConstant0 + 36
	OpConstant37      = OpConstant0 + 37
	OpConstant38      = OpConstant0 + 38
	OpConstant39      = OpConstant0 + 39
	OpConstant40      = OpConstant0 + 40
	OpConstant41      = OpConstant0 + 41
	OpConstant42      = OpConstant0 + 42
	OpConstant43      = OpConstant0 + 43
	OpConstant44      = OpConstant0 + 44
	OpConstant45      = OpConstant0 + 45
	OpConstant46      = OpConstant0 + 46
	OpConstant47      = OpConstant0 + 47
	OpConstant48      = OpConstant0 + 48
	OpConstant49      = OpConstant0 + 49
	OpConstant50      = OpConstant0 + 50
	OpConstant51      = OpConstant0 + 51
	OpConstant52      = OpConstant0 + 52
	OpConstant53      = OpConstant0 + 53
	OpConstant54      = OpConstant0 + 54
	OpConstant55      = OpConstant0 + 55
	OpConstant56      = OpConstant0 + 56
	OpConstant57      = OpConstant0 + 57
	OpConstant58      = OpConstant0 + 58
	OpConstant59      = OpConstant0 + 59
	OpConstant60      = OpConstant0 + 60
	OpConstant61      = OpConstant0 + 61
	OpConstant62      = OpConstant0 + 62
	OpConstant63      = OpConstant0 + 63
)

// Opcodes that are valid after OpExt prefix.
const (
	OpExtStop byte = 0

	OpExtGoCall0 byte = 1
	OpExtGoCall1      = OpExtGoCall0 + 1
	OpExtGoCall2      = OpExtGoCall0 + 2
	OpExtGoCall3      = OpExtGoCall0 + 3
	OpExtGoCall4      = OpExtGoCall0 + 4
	OpExtGoCall5      = OpExtGoCall0 + 5
	OpExtGoCallB      = OpExtGoCall0 + 6
	OpExtGoCallW      = OpExtGoCall0 + 7

	// OpExtShagit is school-pensioner; Pudge from katka (gaem).
	OpExtShagit byte = 0xFF
)

// extOpWidth specifies PC advancement for enumerated opcodes.
var extOpWidth = [256]uint32{
	// OpExtStop does not modify PC as it stops evaluation when executed.
	OpExtStop: 0,

	OpExtGoCall0: 2,
	OpExtGoCall1: 2,
	OpExtGoCall2: 2,
	OpExtGoCall3: 2,
	OpExtGoCall4: 2,
	OpExtGoCall5: 2,
	OpExtGoCallB: 2,
	OpExtGoCallW: 2,
}
