package als

import (
	"bytes"
	json "github.com/funkygao/go-simplejson"
	"strconv"
	"strings"
	"testing"
)

const (
	jsonLineForTest = `us,1389326456474,{"uri":"\/?fb_source=notification&request_ids=629862167081523%2C231759870340420%2C597190080352387%2C640624999328961%2C235464713291862%2C753053901389297%2C790469374302126%2C192819610918125%2C1409213372656992%2C1395677210684824%2C219547141565670%2C445351695593355%2C353291448144469%2C374894915987858%2C1405041129742942%2C1386152901642951%2C1444273795788958%2C268848269934670&ref=notif&app_request_type=user_to_user&notif_t=app_request","_log_info":{"uid":10304512,"snsid":"100006490632784","level":39,"gender":"male","ab":{"pay":"a","quest":"a"},"payment_cash":13,"script_id":9524283412,"serial":1,"uri":"\/","host":"172.31.7.194","ip":"81.65.52.251","callee":"POST+\/+44eae87","sid":null,"elapsed":0.014667987823486}}`
)

func BenchmarkByteStringConvert(b *testing.B) {
	b.ReportAllocs()
	x := []byte(jsonLineForTest)
	for i := 0; i < b.N; i++ {
		_ = string(x)
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	b.ReportAllocs()
	x := []byte(jsonLineForTest)
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		buf.Write(x)
		buf.String()
	}
}

func BenchmarkTrimSpace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.TrimSpace(jsonLineForTest)
	}
}

func BenchmarkParseUint(b *testing.B) {
	// parseAlsLine will use this func
	t := "1389326456474"
	for i := 0; i < b.N; i++ {
		strconv.ParseUint(t, 10, 64)
	}
}

func BenchmarkSplitN(b *testing.B) {
	// parseAlsLine will use this func
	for i := 0; i < b.N; i++ {
		strings.SplitN(jsonLineForTest, ",", 3)
	}
}

func BenchmarkParseAlsLine(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parseAlsLine(jsonLineForTest)
	}
	b.SetBytes(int64(len([]byte(jsonLineForTest))))
}

func BenchmarkLogfileExt(b *testing.B) {
	l := NewAlsLogfile()
	for i := 0; i < b.N; i++ {
		l.SetPath("/var//funplus/logs/fp_rstory/history/session_mm_20131208230103_1")
		l.Ext()
	}
}

func BenchmarkCamelCaseName(b *testing.B) {
	l := NewAlsLogfile()
	for i := 0; i < b.N; i++ {
		l.SetPath("/var//funplus/logs/fp_rstory/history/session_mm_20131208230103_1")
		l.CamelCaseName()
	}
}

func BenchmarkJsonizeByString(b *testing.B) {
	msg := NewAlsMessage()
	_, _, msg.Payload, _ = parseAlsLine(jsonLineForTest)
	for i := 0; i < b.N; i++ {
		msg.jsonize()
		msg.decoded = false
	}
}

func BenchmarkJsonizeByBytes(b *testing.B) {
	_, _, payload, _ := parseAlsLine(jsonLineForTest)
	bp := []byte(payload)
	for i := 0; i < b.N; i++ {
		json.NewJson(bp)
	}
}

func BenchmarkAlsMessageMarshalPayload(b *testing.B) {
	msg := NewAlsMessage()
	if err := msg.FromLine(jsonLineForTest); err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		msg.MarshalPayload()
	}
	b.SetBytes(int64(len([]byte(jsonLineForTest))))
}

func BenchmarkAlsMessageFromEmptyJson(b *testing.B) {
	msg := NewAlsMessage()
	for i := 0; i < b.N; i++ {
		msg.FromEmptyJson()
	}
}

func BenchmarkAlsMessageFromLine(b *testing.B) {
	b.ReportAllocs()
	msg := NewAlsMessage()
	for i := 0; i < b.N; i++ {
		msg.FromLine(jsonLineForTest)
	}
	b.SetBytes(int64(len([]byte(jsonLineForTest))))
}

func BenchmarkAlsMessageFieldValueShallow(b *testing.B) {
	b.ReportAllocs()
	msg := NewAlsMessage()
	msg.FromLine(jsonLineForTest)
	for i := 0; i < b.N; i++ {
		msg.FieldValue("uri", KEY_TYPE_STRING)
	}
}

func BenchmarkAlsMessageFieldValueDeep(b *testing.B) {
	b.ReportAllocs()
	msg := NewAlsMessage()
	msg.FromLine(jsonLineForTest)
	for i := 0; i < b.N; i++ {
		msg.FieldValue("_log_info.uid", KEY_TYPE_INT)
	}
}

func BenchmarkAlsMessageSetField(b *testing.B) {
	b.ReportAllocs()
	msg := NewAlsMessage()
	msg.FromEmptyJson()
	for i := 0; i < b.N; i++ {
		msg.SetField("foo", "bar")
	}
}

func BenchmarkMessageClone(b *testing.B) {
	msg := NewAlsMessage()
	if err := msg.FromLine(jsonLineForTest); err != nil {
		panic(err)
	}

	for i := 0; i < b.N; i++ {
		msg.Clone()
	}
	b.SetBytes(int64(len([]byte(jsonLineForTest))))
}
