// Package dateparse parses date-strings without knowing the format
// in advance, using a fast lex based approach to eliminate shotgun
// attempts.  It leans towards US style dates when there is a conflict.
package dateparse

import (
	"fmt"
	"time"
)

// func init() {
// 	gou.SetupLogging("debug")
// 	gou.SetColorOutput()
// }

var days = []string{
	"mon",
	"tue",
	"wed",
	"thu",
	"fri",
	"sat",
	"sun",
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

var months = []string{
	"january",
	"february",
	"march",
	"april",
	"may",
	"june",
	"july",
	"august",
	"september",
	"october",
	"november",
	"december",
}

type dateState uint8
type timeState uint8

const (
	dateStart dateState = iota // 0
	dateDigit
	dateDigitSt
	dateYearDash
	dateYearDashAlphaDash
	dateYearDashDash
	dateYearDashDashWs // 5
	dateYearDashDashT
	dateYearDashDashOffset
	dateDigitDash
	dateDigitDashAlpha
	dateDigitDashAlphaDash // 10
	dateDigitDot
	dateDigitDotDot
	dateDigitSlash
	dateDigitYearSlash
	dateDigitSlashAlpha // 15
	dateDigitColon
	dateDigitChineseYear
	dateDigitChineseYearWs
	dateDigitWs
	dateDigitWsMoYear // 20
	dateDigitWsMolong
	dateAlpha
	dateAlphaWs
	dateAlphaWsDigit
	dateAlphaWsDigitMore // 25
	dateAlphaWsDigitMoreWs
	dateAlphaWsDigitMoreWsYear
	dateAlphaWsMonth
	dateAlphaWsDigitYearmaybe
	dateAlphaWsMonthMore
	dateAlphaWsMonthSuffix
	dateAlphaWsMore
	dateAlphaWsAtTime
	dateAlphaWsAlpha
	dateAlphaWsAlphaYearmaybe // 35
	dateAlphaPeriodWsDigit
	dateWeekdayComma
	dateWeekdayAbbrevComma
)
const (
	// Time state
	timeIgnore timeState = iota // 0
	timeStart
	timeWs
	timeWsAlpha
	timeWsAlphaWs
	timeWsAlphaZoneOffset // 5
	timeWsAlphaZoneOffsetWs
	timeWsAlphaZoneOffsetWsYear
	timeWsAlphaZoneOffsetWsExtra
	timeWsAMPMMaybe
	timeWsAMPM // 10
	timeWsOffset
	timeWsOffsetWs // 12
	timeWsOffsetColonAlpha
	timeWsOffsetColon
	timeWsYear // 15
	timeOffset
	timeOffsetColon
	timeAlpha
	timePeriod
	timePeriodOffset // 20
	timePeriodOffsetColon
	timePeriodOffsetColonWs
	timePeriodWs
	timePeriodWsAlpha
	timePeriodWsOffset // 25
	timePeriodWsOffsetWs
	timePeriodWsOffsetWsAlpha
	timePeriodWsOffsetColon
	timePeriodWsOffsetColonAlpha
	timeZ
	timeZDigit
)

var (
	// ErrAmbiguousMMDD for date formats such as 04/02/2014 the mm/dd vs dd/mm are
	// ambiguous, so it is an error for strict parse rules.
	ErrAmbiguousMMDD = fmt.Errorf("This date has ambiguous mm/dd vs dd/mm type format")
)

func unknownErr(datestr string) error { _ = "STUB: not implemented"; return nil }

// ParseAny parse an unknown date format, detect the layout.
// Normal parse.  Equivalent Timezone rules as time.Parse().
// NOTE:  please see readme on mmdd vs ddmm ambiguous dates.
func ParseAny(datestr string, opts ...ParserOption) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ParseIn with Location, equivalent to time.ParseInLocation() timezone/offset
// rules.  Using location arg, if timezone/offset info exists in the
// datestring, it uses the given location rules for any zone interpretation.
// That is, MST means one thing when using America/Denver and something else
// in other locations.
func ParseIn(datestr string, loc *time.Location, opts ...ParserOption) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// ParseLocal Given an unknown date format, detect the layout,
// using time.Local, parse.
//
// Set Location to time.Local.  Same as ParseIn Location but lazily uses
// the global time.Local variable for Location argument.
//
//	denverLoc, _ := time.LoadLocation("America/Denver")
//	time.Local = denverLoc
//
//	t, err := dateparse.ParseLocal("3/1/2014")
//
// Equivalent to:
//
//	t, err := dateparse.ParseIn("3/1/2014", denverLoc)
func ParseLocal(datestr string, opts ...ParserOption) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

// MustParse  parse a date, and panic if it can't be parsed.  Used for testing.
// Not recommended for most use-cases.
func MustParse(datestr string, opts ...ParserOption) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// ParseFormat parse's an unknown date-time string and returns a layout
// string that can parse this (and exact same format) other date-time strings.
//
//	layout, err := dateparse.ParseFormat("2013-02-01 00:00:00")
//	// layout = "2006-01-02 15:04:05"
func ParseFormat(datestr string, opts ...ParserOption) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ParseStrict parse an unknown date format.  IF the date is ambigous
// mm/dd vs dd/mm then return an error. These return errors:   3.3.2014 , 8/8/71 etc
func ParseStrict(datestr string, opts ...ParserOption) (time.Time, error) {
	_ = "STUB: not implemented"
	return *new(time.Time), nil
}

func parseTime(datestr string, loc *time.Location, opts ...ParserOption) (p *parser, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// month out of range signifies that a day/month swap is the correct solution to an ambiguous date
// this is because it means that a day is being interpreted as a month and overflowing the valid value for that
// by retrying in this case, we can fix a common situation with no assumptions

// if it errors out with the following error, swap before we
// get out of this function to reduce scope it needs to be applied on

// create the option to reverse the preference

// turn off the retry to avoid endless recursion

// General strategy is to read rune by rune through the date looking for
// certain hints of what type of date we are dealing with.
// Hopefully we only need to read about 5 or 6 bytes before
// we figure it out and then attempt a parse

//r := rune(datestr[i])

// gou.Debugf("i=%d r=%s state=%d   %s", i, string(r), p.stateDate, datestr)

// 2006-01-02
// 2013-Feb-03
// 13-Feb-03
// 29-Jun-2016

// 08/May/2005
// 03/31/2005
// 2014/02/24

// 2014/02/24  -  Year first /
// since it was start of datestr, i=len

// Either Ambiguous dd/mm vs mm/dd  OR dd/month/yy
// 08/May/2005
// 03/31/2005
// 31/03/2005

// 08/May/2005

// Ambiguous dd/mm vs mm/dd the bane of date-parsing
// 03/31/2005
// 31/03/2005

// 03/31/2005

// 03/31/2005
// 2014/02/24

// 3.31.2014
// 08.21.71
// 2014.05

// 18 January 2018
// 8 January 2018
// 8 jan 2018
// 02 Jan 2018 23:59
// 02 Jan 2018 23:59:34
// 12 Feb 2006, 19:17
// 12 Feb 2006, 19:17:22

// Chinese Year

// dateYearDashDashT
//  2006-01-02T15:04:05Z07:00
//  2020-08-17T17:00:00:000+0100
// dateYearDashDashWs
//  2013-04-01 22:43:22
// dateYearDashAlphaDash
//   2013-Feb-03

// dateYearDashDashT
//  2006-01-02T15:04:05Z07:00
// dateYearDashDashWs
//  2013-04-01 22:43:22
// dateYearDashDashOffset
//  2020-07-20+00:00

// dateYearDashDashT
//  2006-01-02T15:04:05Z07:00
//  2020-08-17T17:00:00:000+0100

//  2020-07-20+00:00

// case ' ':
// 	return nil, unknownErr(datestr)

// 2013-Feb-03

// 13-Feb-03
// 29-Jun-2016

// 13-Feb-03
// 28-Feb-03
// 29-Jun-2016

// 13-Feb-03   ambiguous
// 28-Feb-03   ambiguous
// 29-Jun-2016  dd-month(alpha)-yyyy

// we need to find if this was 4 digits, aka year
// or 2 digits which makes it ambiguous year/day

// We now also know that part1 was the day

// We have no idea if this is
// yy-mon-dd   OR  dd-mon-yy
//
// We are going to ASSUME (bad, bad) that it is dd-mon-yy  which is a horible assumption

// We now also know that part1 was the day

// 2014/07/10 06:55:38.156283
// I honestly don't know if this format ever shows up as yyyy/

// 06/May/2008

//       |
// 06/May/2008

// We aren't breaking because we are going to re-use this case
// to find where the date starts, and possible time begins

// 03/19/2012 10:11:59
// 04/2/2014 03:00:37
// 3/1/2012 10:11:59
// 4/8/2014 22:05
// 3/1/2014
// 10/13/2014
// 01/02/2006
// 1/2/06

// This is the 2nd / so now we should know start pts of all of the dd, mm, yy

// Note no break, we are going to pass by and re-enter this dateDigitSlash
// and look for ending (space) or not (just date)

// 2014:07:10 06:55:38.156283
// 03:19:2012 10:11:59
// 04:2:2014 03:00:37
// 3:1:2012 10:11:59
// 4:8:2014 22:05
// 3:1:2014
// 10:13:2014
// 01:02:2006
// 1:2:06

// 2014:07:10 06:55:38.156283

// 18 January 2018
// 8 January 2018
// 8 jan 2018
// 1 jan 18
// 02 Jan 2018 23:59
// 02 Jan 2018 23:59:34
// 12 Feb 2006, 19:17
// 12 Feb 2006, 19:17:22

//p.yearlen = 4

//  November etc
// If len greather than space + 3 it must be full month

// If len=3, the might be Feb or May?  Ie ambigous abbreviated but
// we can parse may with either.  BUT, that means the
// format may not be correct?
// mo := strings.ToLower(datestr[p.daylen+1 : i])

// 8 jan 2018
// 02 Jan 2018 23:59
// 02 Jan 2018 23:59:34
// 12 Feb 2006, 19:17
// 12 Feb 2006, 19:17:22

// 18 January 2018
// 8 January 2018

// dateDigitChineseYear
//   2014年04月08日
//               weekday  %Y年%m月%e日 %A %I:%M %p
// 2013年07月18日 星期四 10:27 上午

// This is the 2nd period
// 3.31.2014
// 08.21.71
// 2014.05
// 2018.09.30

// 3.31.2014

// 2018.09.30
//p.molen = 2

// iterate all the way through

// dateAlphaWS
//  Mon Jan _2 15:04:05 2006
//  Mon Jan _2 15:04:05 MST 2006
//  Mon Jan 02 15:04:05 -0700 2006
//  Mon Aug 10 15:44:11 UTC+0100 2015
//  Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)
//  dateAlphaWSDigit
//    May 8, 2009 5:57:51 PM
//    oct 1, 1970
//  dateAlphaWsMonth
//    April 8, 2009
//  dateAlphaWsMore
//    dateAlphaWsAtTime
//      January 02, 2006 at 3:04pm MST-07
//
//  dateAlphaPeriodWsDigit
//    oct. 1, 1970
// dateWeekdayComma
//   Monday, 02 Jan 2006 15:04:05 MST
//   Monday, 02-Jan-06 15:04:05 MST
//   Monday, 02 Jan 2006 15:04:05 -0700
//   Monday, 02 Jan 2006 15:04:05 +0100
// dateWeekdayAbbrevComma
//   Mon, 02 Jan 2006 15:04:05 MST
//   Mon, 02 Jan 2006 15:04:05 -0700
//   Thu, 13 Jul 2017 08:58:40 +0100
//   Tue, 11 Jul 2017 16:28:13 +0200 (CEST)
//   Mon, 02-Jan-06 15:04:05 MST

//      X
// April 8, 2009

// Check to see if the alpha is name of month?  or Day?

// len(" 31, 2018")   = 9

// April 8, 2009

// This is possibly ambiguous?  May will parse as either though.
// So, it could return in-correct format.
// dateAlphaWs
//   May 05, 2005, 05:05:05
//   May 05 2005, 05:05:05
//   Jul 05, 2005, 05:05:05
//   May 8 17:57:51 2009
//   May  8 17:57:51 2009
// skip & return to dateStart
//   Tue 05 May 2020, 05:05:05
//   Mon Jan  2 15:04:05 2006

// using skip throws off indices used by other code; saner to restart

// Mon, 02 Jan 2006

// TODO:  lets just make this "skip" as we don't need
// the mon, monday, they are all superfelous and not needed
// just lay down the skip, no need to fill and then skip

// sept. 28, 2017
// jan. 28, 2017

// gross

// dateAlphaWsAlpha
//   Mon Jan _2 15:04:05 2006
//   Mon Jan _2 15:04:05 MST 2006
//   Mon Jan 02 15:04:05 -0700 2006
//   Fri Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)
//   Mon Aug 10 15:44:11 UTC+0100 2015
// dateAlphaWsDigit
//   May 8, 2009 5:57:51 PM
//   May 8 2009 5:57:51 PM
//   May 8 17:57:51 2009
//   May  8 17:57:51 2009
//   May 08 17:57:51 2009
//   oct 1, 1970
//   oct 7, '70

// May 8, 2009 5:57:51 PM
// May 8 2009 5:57:51 PM
// oct 1, 1970
// oct 7, '70
// oct. 7, 1970
// May 8 17:57:51 2009
// May  8 17:57:51 2009
// May 08 17:57:51 2009

//       x
// May 8 2009 5:57:51 PM
// May 8 17:57:51 2009
// May  8 17:57:51 2009
// May 08 17:57:51 2009
// Jul 03 2015 18:04:07 GMT+0100 (GMT Daylight Time)

// Guessed wrong; was not a year

// must be year format, not 15:04

//       x
// May 8, 2009 5:57:51 PM
// May 05, 2005, 05:05:05
// May 05 2005, 05:05:05
// oct 1, 1970
// oct 7, '70

//            x
// May 8, 2009 5:57:51 PM
// May 05, 2005, 05:05:05
// oct 1, 1970
// oct 7, '70

//            x
// May 8, 2009 5:57:51 PM
//            x
// May 8, 2009, 5:57:51 PM

// April 8, 2009
// April 8 2009

//       x
// June 8, 2009
//       x
// June 8 2009

// st, rd, nd, st

//                  X
// January 02, 2006, 15:04:05
// January 02 2006, 15:04:05
// January 02, 2006 15:04:05
// January 02 2006 15:04:05

//        x
// April 8th, 2009
// April 8th 2009

// January 02, 2006, 15:04:05
// January 02 2006, 15:04:05
// January 2nd, 2006, 15:04:05
// January 2nd 2006, 15:04:05
// September 17, 2012 at 5:00pm UTC-05

//           x
// January 02, 2006, 15:04:05

//           x
// January 02 2006, 15:04:05

//         XX
// January 02, 2006, 15:04:05

//          X
// January 2nd, 2006, 15:04:05

//    oct. 7, '70

// continue

// Monday, 02 Jan 2006 15:04:05 MST
// Monday, 02 Jan 2006 15:04:05 -0700
// Monday, 02 Jan 2006 15:04:05 +0100
// Monday, 02-Jan-06 15:04:05 MST

// Mon, 02 Jan 2006 15:04:05 MST
// Mon, 02 Jan 2006 15:04:05 -0700
// Thu, 13 Jul 2017 08:58:40 +0100
// Thu, 4 Jan 2018 17:53:36 +0000
// Tue, 11 Jul 2017 16:28:13 +0200 (CEST)
// Mon, 02-Jan-06 15:04:05 MST

// increment first one, since the i++ occurs at end of loop

// ensure we skip any whitespace prefix

// gou.Debugf("i=%d r=%s state=%d iterTimeRunes  %s %s", i, string(r), p.stateTime, p.ds(), p.ts())

// 22:43:22
// 22:43
// timeComma
//   08:20:13,787
// timeWs
//   05:24:37 PM
//   06:20:00 UTC
//   06:20:00 UTC-05
//   00:12:00 +0000 UTC
//   22:18:00 +0000 UTC m=+0.000000001
//   15:04:05 -0700
//   15:04:05 -07:00
//   15:04:05 2008
// timeOffset
//   03:21:51+00:00
//   19:55:00+0100
// timePeriod
//   17:24:37.3186369
//   00:07:31.945167
//   18:31:59.257000000
//   00:00:00.000
//   timePeriodOffset
//     19:55:00.799+0100
//     timePeriodOffsetColon
//       15:04:05.999-07:00
//   timePeriodWs
//     timePeriodWsOffset
//       00:07:31.945167 +0000
//       00:00:00.000 +0000
//     timePeriodWsOffsetAlpha
//       00:07:31.945167 +0000 UTC
//       22:18:00.001 +0000 UTC m=+0.000000001
//       00:00:00.000 +0000 UTC
//     timePeriodWsAlpha
//       06:20:00.000 UTC

// hm, lets just swap out comma for period.  for some reason go
// won't parse it.
// 2014-05-11 08:20:13,787

//   03:21:51+00:00

// 22:18+0530

// (Z)ulu time

//                    x
// September 17, 2012 at 5:00pm UTC-05
// skip t

//                      x
// September 17, 2012 at 5:00pm UTC-05
// skip '
// reset hour

// Could be AM/PM

// 18:31:59:257    ms uses colon, wtf

// gross, gross, gross.   manipulating the datestr is horrible.
// https://github.com/araddon/dateparse/issues/117
// Could not get the parsing to work using golang time.Parse() without
// replacing that colon with period.

// 19:55:00+0100
// timeOffsetColon
//   15:04:05+07:00
//   15:04:05-07:00

// timeWsAlpha
//   06:20:00 UTC
//   06:20:00 UTC-05
//   15:44:11 UTC+0100 2015
//   18:04:07 GMT+0100 (GMT Daylight Time)
//   17:57:51 MST 2009
//   timeWsAMPMMaybe
//     05:24:37 PM
// timeWsOffset
//   15:04:05 -0700
//   00:12:00 +0000 UTC
//   timeWsOffsetColon
//     15:04:05 -07:00
//     17:57:51 -0700 2009
//     timeWsOffsetColonAlpha
//       00:12:00 +00:00 UTC
// timeWsYear
//     00:12:00 2008
// timeZ
//   15:04:05.99Z

// Could be AM/PM or could be PST or similar

// 06:20:00 UTC
// 06:20:00 UTC-05
// 15:44:11 UTC+0100 2015
// 17:57:51 MST 2009

// 00:12:00 2008

// 06:20:00 UTC
// 06:20:00 UTC-05
// timeWsAlphaWs
//   17:57:51 MST 2009
// timeWsAlphaZoneOffset
// timeWsAlphaZoneOffsetWs
//   timeWsAlphaZoneOffsetWsExtra
//     18:04:07 GMT+0100 (GMT Daylight Time)
//   timeWsAlphaZoneOffsetWsYear
//     15:44:11 UTC+0100 2015

// 17:57:51 MST 2009
// 17:57:51 MST

//   17:57:51 MST 2009

// 06:20:00 UTC-05
// timeWsAlphaZoneOffset
// timeWsAlphaZoneOffsetWs
//   timeWsAlphaZoneOffsetWsExtra
//     18:04:07 GMT+0100 (GMT Daylight Time)
//   timeWsAlphaZoneOffsetWsYear
//     15:44:11 UTC+0100 2015

// timeWsAlphaZoneOffsetWs
//   timeWsAlphaZoneOffsetWsExtra
//     18:04:07 GMT+0100 (GMT Daylight Time)
//   timeWsAlphaZoneOffsetWsYear
//     15:44:11 UTC+0100 2015

// 15:44:11 UTC+0100 2015

// timeWsAMPMMaybe
//   timeWsAMPM
//     05:24:37 PM
//   timeWsAlpha
//     00:12:00 PST
//     15:44:11 UTC+0100 2015

//return parse("2006-01-02 03:04:05 PM", datestr, loc)

// timeWsOffset
//   15:04:05 -0700
//   timeWsOffsetWsOffset
//     17:57:51 -0700 -07
//   timeWsOffsetWs
//     17:57:51 -0700 2009
//     00:12:00 +0000 UTC
//   timeWsOffsetColon
//     15:04:05 -07:00
//     timeWsOffsetColonAlpha
//       00:12:00 +00:00 UTC

// 17:57:51 -0700 2009
// 00:12:00 +0000 UTC
// 22:18:00.001 +0000 UTC m=+0.000000001
// w Extra
//   17:57:51 -0700 -07

// eff you golang

// This really doesn't seem valid, but for some reason when round-tripping a go date
// their is an extra +03 printed out.  seems like go bug to me, but, parsing anyway.
// 00:00:00 +0300 +03
// 00:00:00 +0300 +0300

// 15:04:05 -0700 MST

// timeWsOffsetColon
//   15:04:05 -07:00
//   timeWsOffsetColonAlpha
//     2015-02-18 00:12:00 +00:00 UTC

// 2015-02-18 00:12:00 +00:00 UTC

// 15:04:05.999999999+07:00
// 15:04:05.999999999-07:00
// 15:04:05.999999+07:00
// 15:04:05.999999-07:00
// 15:04:05.999+07:00
// 15:04:05.999-07:00
// timePeriod
//   17:24:37.3186369
//   00:07:31.945167
//   18:31:59.257000000
//   00:00:00.000
//   timePeriodOffset
//     19:55:00.799+0100
//     timePeriodOffsetColon
//       15:04:05.999-07:00
//   timePeriodWs
//     timePeriodWsOffset
//       00:07:31.945167 +0000
//       00:00:00.000 +0000
//       With Extra
//         00:00:00.000 +0300 +03
//     timePeriodWsOffsetAlpha
//       00:07:31.945167 +0000 UTC
//       00:00:00.000 +0000 UTC
//       22:18:00.001 +0000 UTC m=+0.000000001
//     timePeriodWsAlpha
//       06:20:00.000 UTC

// This really shouldn't happen

// 06:20:00.000 UTC

// timePeriodOffset
//   19:55:00.799+0100
//   timePeriodOffsetColon
//     15:04:05.999-07:00
//     13:31:51.999-07:00 MST

// timePeriodOffset
//   timePeriodOffsetColon
//     15:04:05.999-07:00
//     13:31:51.999 -07:00 MST

// continue

// timePeriodWs
//   timePeriodWsOffset
//     00:07:31.945167 +0000
//     00:00:00.000 +0000
//   timePeriodWsOffsetAlpha
//     00:07:31.945167 +0000 UTC
//     00:00:00.000 +0000 UTC
//   timePeriodWsOffsetColon
//     13:31:51.999 -07:00 MST
//   timePeriodWsAlpha
//     06:20:00.000 UTC

//     00:07:31.945167 +0000 UTC
//     00:00:00.000 +0000 UTC

// timePeriodWs
//   timePeriodWsOffset
//     00:07:31.945167 +0000
//     00:00:00.000 +0000
//     With Extra
//       00:00:00.000 +0300 +03
//   timePeriodWsOffsetAlpha
//     00:07:31.945167 +0000 UTC
//     00:00:00.000 +0000 UTC
//     03:02:00.001 +0300 MSK m=+0.000000001
//   timePeriodWsOffsetColon
//     13:31:51.999 -07:00 MST
//   timePeriodWsAlpha
//     06:20:00.000 UTC

// This really doesn't seem valid, but for some reason when round-tripping a go date
// their is an extra +03 printed out.  seems like go bug to me, but, parsing anyway.
// 00:00:00.000 +0300 +03
// 00:00:00.000 +0300 +0300

// 00:07:31.945167 +0000 UTC
// 00:00:00.000 +0000 UTC
// 03:02:00.001 +0300 MSK m=+0.000000001

// 03:02:00.001 +0300 MSK m=+0.000000001
// eff you golang

// 13:31:51.999 -07:00 MST

// 13:31:51.999 -07:00 MST

// continue

// timeZ
//   15:04:05.99Z
// With a time-zone at end after Z
// 2006-01-02T15:04:05.999999999Z07:00
// 2006-01-02T15:04:05Z07:00
// RFC3339     = "2006-01-02T15:04:05Z07:00"
// RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"

// 13:31:51.999 +01:00 CET

// 06:20:00 UTC-05

// 19:55:00+01

// 19:55:00+0100

// 17:57:51 -0700 2009
// 00:12:00 +0000 UTC

// 13:31:51.999 +01:00 CET

// 13:31:51.999 +01:00 CEST

// 17:57:51 -07:00

// 15:04:05+07:00

// 19:55:00.799+0100

// unixy timestamps ish
//  example              ct type
//  1499979655583057426  19 nanoseconds
//  1499979795437000     16 micro-seconds
//  20180722105203       14 yyyyMMddhhmmss
//  1499979795437        13 milliseconds
//  1332151919           10 seconds
//  20140601             8  yyyymmdd
//  2014                 4  yyyy

// 19
// nano-seconds

// 16
// micro-seconds

// 14
// yyyyMMddhhmmss

// 13

//10

// 171113 14:14:20

// 2006-01

// 2006-01-02
// 2006-1-02
// 2006-1-2
// 2006-01-2

///  2020-07-20+00:00

// 2013-Feb-03
// 2013-Feb-3

// 2013-04-01

// 13-Feb-03   ambiguous
// 28-Feb-03   ambiguous
// 29-Jun-2016

// We now also know that part1 was the day

// We have no idea if this is
// yy-mon-dd   OR  dd-mon-yy
//
// We are going to ASSUME (bad, bad) that it is dd-mon-yy  which is a horible assumption

// We now also know that part1 was the day

// 2014.05

// 03.31.1981
// 3.31.2014
// 3.2.1981
// 3.2.81
// 08.21.71
// 2018.09.30

// 2 Jan 2018
// 2 Jan 18
// 2 Jan 2018 23:59
// 02 Jan 2018 23:59
// 12 Feb 2006, 19:17

// 18 January 2018
// 8 January 2018

// parse("2 January 2006", datestr, loc)

// oct 1, 1970

// May 8, 2009 5:57:51 PM
// Jun 7, 2005, 05:57:51

// 3/1/2014
// 10/13/2014
// 01/02/2006

// 03/Jun/2014

// 2014/10/13

// 3:1:2014
// 10:13:2014
// 01:02:2006
// 2014:10:13

// dateDigitChineseYear
//   2014年04月08日

// Monday, 02 Jan 2006 15:04:05 -0700
// Monday, 02 Jan 2006 15:04:05 +0100
// Monday, 02-Jan-06 15:04:05 MST

// Mon, 02-Jan-06 15:04:05 MST
// Mon, 02 Jan 2006 15:04:05 MST

type parser struct {
	loc                        *time.Location
	preferMonthFirst           bool
	retryAmbiguousDateWithSwap bool
	ambiguousMD                bool
	stateDate                  dateState
	stateTime                  timeState
	format                     []byte
	datestr                    string
	fullMonth                  string
	skip                       int
	extra                      int
	part1Len                   int
	yeari                      int
	yearlen                    int
	moi                        int
	molen                      int
	dayi                       int
	daylen                     int
	houri                      int
	hourlen                    int
	mini                       int
	minlen                     int
	seci                       int
	seclen                     int
	msi                        int
	mslen                      int
	offseti                    int
	offsetlen                  int
	tzi                        int
	tzlen                      int
	t                          *time.Time
}

// ParserOption defines a function signature implemented by options
// Options defined like this accept the parser and operate on the data within
type ParserOption func(*parser) error

// PreferMonthFirst is an option that allows preferMonthFirst to be changed from its default
func PreferMonthFirst(preferMonthFirst bool) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

// RetryAmbiguousDateWithSwap is an option that allows retryAmbiguousDateWithSwap to be changed from its default
func RetryAmbiguousDateWithSwap(retryAmbiguousDateWithSwap bool) ParserOption {
	_ = "STUB: not implemented"
	return *new(ParserOption)
}

func newParser(dateStr string, loc *time.Location, opts ...ParserOption) *parser {
	_ = "STUB: not implemented"
	return nil
}

// allow the options to mutate the parser fields from their defaults

func (p *parser) nextIs(i int, b byte) bool { _ = "STUB: not implemented"; return false }

func (p *parser) set(start int, val string) { _ = "STUB: not implemented"; return }

func (p *parser) setMonth() { _ = "STUB: not implemented"; return }

func (p *parser) setDay() { _ = "STUB: not implemented"; return }

func (p *parser) setYear() { _ = "STUB: not implemented"; return }

func (p *parser) coalesceDate(end int) { _ = "STUB: not implemented"; return }

func (p *parser) ts() string { _ = "STUB: not implemented"; return "" }

func (p *parser) ds() string { _ = "STUB: not implemented"; return "" }

func (p *parser) coalesceTime(end int) {
	_ = "STUB: not implemented"
	// 03:04:05
	// 15:04:05
	// 3:04:05
	// 3:4:5
	// 15:04:05.00
	return
}

func (p *parser) setFullMonth(month string) { _ = "STUB: not implemented"; return }

func (p *parser) trimExtra() { _ = "STUB: not implemented"; return }

// func (p *parser) remove(i, length int) {
// 	if len(p.format) > i+length {
// 		//append(a[:i], a[j:]...)
// 		p.format = append(p.format[0:i], p.format[i+length:]...)
// 	}
// 	if len(p.datestr) > i+length {
// 		//append(a[:i], a[j:]...)
// 		p.datestr = fmt.Sprintf("%s%s", p.datestr[0:i], p.datestr[i+length:])
// 	}
// }

func (p *parser) parse() (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }

// gou.Debugf("parse layout=%q input=%q   \ntx, err := time.Parse(%q, %q)", string(p.format), p.datestr, string(p.format), p.datestr)

//gou.Debugf("parse layout=%q input=%q   \ntx, err := time.ParseInLocation(%q, %q, %v)", string(p.format), p.datestr, string(p.format), p.datestr, p.loc)

func isDay(alpha string) bool { _ = "STUB: not implemented"; return false }

func isMonthFull(alpha string) bool { _ = "STUB: not implemented"; return false }
