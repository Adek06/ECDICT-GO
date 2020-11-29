// Code generated by entc, DO NOT EDIT.

package ecdict

const (
	// Label holds the string label denoting the ecdict type in the database.
	Label = "ecdict"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldWord holds the string denoting the word field in the database.
	FieldWord = "word"
	// FieldSw holds the string denoting the sw field in the database.
	FieldSw = "sw"
	// FieldPhonetic holds the string denoting the phonetic field in the database.
	FieldPhonetic = "phonetic"
	// FieldDefinition holds the string denoting the definition field in the database.
	FieldDefinition = "definition"
	// FieldTranslation holds the string denoting the translation field in the database.
	FieldTranslation = "translation"
	// FieldPos holds the string denoting the pos field in the database.
	FieldPos = "pos"
	// FieldCollins holds the string denoting the collins field in the database.
	FieldCollins = "collins"
	// FieldOxford holds the string denoting the oxford field in the database.
	FieldOxford = "oxford"
	// FieldTag holds the string denoting the tag field in the database.
	FieldTag = "tag"
	// FieldBnc holds the string denoting the bnc field in the database.
	FieldBnc = "bnc"
	// FieldFrq holds the string denoting the frq field in the database.
	FieldFrq = "frq"
	// FieldExchange holds the string denoting the exchange field in the database.
	FieldExchange = "exchange"
	// FieldDetail holds the string denoting the detail field in the database.
	FieldDetail = "detail"
	// FieldAudio holds the string denoting the audio field in the database.
	FieldAudio = "audio"

	// Table holds the table name of the ecdict in the database.
	Table = "ecdicts"
)

// Columns holds all SQL columns for ecdict fields.
var Columns = []string{
	FieldID,
	FieldWord,
	FieldSw,
	FieldPhonetic,
	FieldDefinition,
	FieldTranslation,
	FieldPos,
	FieldCollins,
	FieldOxford,
	FieldTag,
	FieldBnc,
	FieldFrq,
	FieldExchange,
	FieldDetail,
	FieldAudio,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
