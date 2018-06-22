package gonwr_test

import (
	"strings"
	"testing"

	"github.com/ThomasK81/gonwr"
)

type testgroup struct {
	input  []string
	output []string
	score  int
}

var test = testgroup{
	input:  []string{"parīkṣitāni pramāṇāni prameyam idānīṃ parīkṣyate tac cātmādīty ātmā vivicyate kiṃ dehendriyamanobuddhisaṃghātamātra ātmā āhosvit tadvyatirikta iti kutaḥ saṃśayaḥ vyapadeśasyobhayathā siddheḥ kriyākaraṇayoḥ kartrā saṃbandhasyābhidhānaṃ vyapadeśaḥ sa dvividhaḥ avayavena samudāyasya mūlair vṛkṣas tiṣṭhati stambhaiḥ prāsādo dhriyata iti anyenānyasya vyapadeśaḥ paraśunā vṛścati pradīpena paśyati asti cāyaṃ vyapadeśaḥ cakṣuṣā paśyati manasā vijānāti buddhyā vicārayati śarīreṇa sukhaduḥkham anubhavatīti tatra nāvadhāryate kim avayavena samudāyasya dehādisaṃghātasya athānyenāsya tadvyatiriktasyeti anyenāyam anyasya vyapadeśaḥ kasmāt darśanasparśanābhyām ekārthagrahaṇāt darśanena kaścid artho gṛhītaḥ sparśanenāpi so rtho gṛhyate yam aham adrākṣaṃ cakṣuṣā taṃ sparśanenāpi spṛśāmīti yaṃ cāspārkṣaṃ sparśanena taṃ cakṣuṣā paśyāmīti ekaviṣayau cemau pratyayāv ekakarttṛkau pratisandhīyete na ca saṃghātakartṛkau nendriyeṇaikakartṛkau tad yo sau cakṣuṣā tvagindriyeṇa caikārthasya saṃgṛhītā bhinnanimittāv ananyakartṛkau pratyayau samānaviṣayau pratisandadhāti so rthāntarabhūta ātmā kathaṃ punar nendriyeṇaikakartṛkau indriyaṃ khalu svaṃ svaṃ viṣayagrahaṇam ananyakartṛkaṃ pratisandhātum arhati nendriyāntarasya viṣayāntaragrahaṇam iti kathaṃ na saṃghātakartṛkau ekaḥ khalv ayaṃ bhinnanimittau svātmakartṛkau pratisaṃhitau vedayate na saṃghātaḥ kasmāt anivṛttaṃ hi saṃghāte pratyekaṃ viṣayāntaragrahaṇasyāpratisandhānam indriyāṃtareṇeveti ", "parīkṣitāni pramāṇāni | prameyam idānīṃ parīkṣyate | tac cātmādīty ātmā vivicyate kiṃ dehendriyamanobuddhivedanāsaṅghātamātram ātmāho svit tato vyatirikta iti | kutaḥ saṃśayaḥ | vyapadeśasyobhayathā siddheḥ saṃśayaḥ | kriyākaraṇayoḥ kartrā sambandhasyābhidhānaṃ vyapadeśaḥ | sa dvividhaḥ | avayavena samudāyasya mūlair vṛkṣas tiṣṭhati stambhaiḥ prāsādo dhriyata iti | anyena cānyasya vyapadeśaḥ paraśunā vṛścati pradīpena paśyatīti | asti cāyaṃ vyapadeśaś cakṣuṣā paśyati manasā vijānāti buddhyā vicārayati śarīreṇa sukhaduḥkham anubhavatīti | tatra nāvadhāryate kim avayavena samudāyasya dehādisaṅghātasya vyapadeśaḥ | athānyenānyasya tadvyatiriktasya veti | anyenānyasya vyapadeśaḥ | kasmāt | darśanasparśanābhyām ekārthagrahaṇāt | darśanena kaścid artho gṛhītaḥ sparśanenāpi so 'rtho gṛhyate yam aham adrākṣaṃ cakṣuṣā taṃ sparśanenāpi spṛśāmīti yaṃ cāspārkṣaṃ sparśanena taṃ cakṣuṣā paśyāmīti | ekaviṣayau dvāv imau pratyayāv ekakartṛkau pratisandhīyete | na ca saṅghātakartṛkau nendriyeṇaikakartṛkau | tad yo 'sau cakṣuṣā tvagindriyeṇa caikārthasya grahītā bhinnanimittāv ananyakartṛkau pratyayau samānaviṣayau pratisandadhāti so 'rthāntarabhūta ātmeti | kathaṃ punar nendriyeṇaikakartṛkau | indriyaṃ khalu svaṃ svaṃ viṣayagrahaṇam ananyakartṛkaṃ pratisandhātum arhati nendriyāntarasya viṣayāntaragrahaṇam iti | kathaṃ na saṅghātakartṛkau | ekaḥ khalv ayaṃ bhinnanimittau svātmakartṛkau pratyayau pratisaṃhitau vedayate na saṅghātaḥ | kasmāt | anivṛttaṃ hi saṅghāte pratyekaṃ viṣayāntaragrahaṇasyāpratisandhānam indriyāntareṇeveti | "},
	output: []string{"parīkṣitāni pramāṇāni ##prameyam idānīṃ parīkṣyate ##tac cātmādīty ātmā vivicyate kiṃ dehendriyamanobudd#hi######saṃghātamātra# ātmā āho#svit t##advyatirikta iti ##kutaḥ saṃśayaḥ ##vyapadeśasyobhayathā sidd#heḥ######### ##kriyākaraṇayoḥ kartrā saṃbandhasyābhidhānaṃ vyapadeśaḥ ##sa dvividhaḥ ##avayavena samudāyasya mūlair vṛkṣas tiṣṭhati stambhaiḥ prāsādo dhriyata iti ##anyen###ānyasya vyapadeśaḥ paraśunā vṛścati pradīpena paśyat##i ##asti cāyaṃ vyapadeśaḥ cakṣuṣā paśyati manasā vijānāti budd#hyā vicārayati śarīreṇa sukhaduḥkham anubhavatīti ##tatra nāvadhāryate kim avayavena samudāyasya dehādisaṃghātasy####a####### ##athānyen###āsya tadvyatiriktasy###eti ##anyenā#yam anyasya vyapadeśaḥ ##kasmāt ##darśanasparśanābhyām ekārthagrahaṇāt ##darśanena kaścid artho gṛhītaḥ sparśanenāpi so #rtho gṛhyate yam aham adrākṣaṃ cakṣuṣā taṃ sparśanenāpi spṛśāmīti yaṃ cāspārkṣaṃ sparśanena taṃ cakṣuṣā paśyāmīti ##ekaviṣayau #####cemau pratyayāv ekakarttṛkau pratisandhīyete ##na ca saṃghātakartṛkau nendriyeṇaikakartṛkau ##tad yo #sau cakṣuṣā tvagindriyeṇa caikārthasya saṃg#ṛhītā bhinn#animitt#āv ananyakartṛkau pratyayau samānaviṣayau pratisandadhāti so #rthāntarabhūta ātm##ā ##kathaṃ punar nendriyeṇaikakartṛkau ##indriyaṃ khalu svaṃ svaṃ viṣayagrahaṇam ananyakartṛkaṃ pratisandhātum arhati nendriyāntarasya viṣayāntaragrahaṇam iti ##kathaṃ na saṃghātakartṛkau ##ekaḥ khalv ayaṃ bhinn#animitt#au svātmakartṛka##########u pratisaṃhitau vedayate na saṃghātaḥ ##kasmāt ##anivṛtt#aṃ hi saṃghāte pratyekaṃ viṣayāntaragrahaṇasyāpratisandhānam indriyāṃtareṇeveti ##", "parīkṣitāni pramāṇāni | prameyam idānīṃ parīkṣyate | tac cātmādīty ātmā vivicyate kiṃ dehendriyamanobud#dhivedanāsaṅghātamātram ātmā##ho svit tato vyatirikta iti | kutaḥ saṃśayaḥ | vyapadeśasyobhayathā sid#dheḥ saṃśayaḥ | kriyākaraṇayoḥ kartrā sambandhasyābhidhānaṃ vyapadeśaḥ | sa dvividhaḥ | avayavena samudāyasya mūlair vṛkṣas tiṣṭhati stambhaiḥ prāsādo dhriyata iti | anyena cānyasya vyapadeśaḥ paraśunā vṛścati pradīpena paśyatīti | asti cāyaṃ vyapadeśaś cakṣuṣā paśyati manasā vijānāti bud#dhyā vicārayati śarīreṇa sukhaduḥkham anubhavatīti | tatra nāvadhāryate kim avayavena samudāyasya dehādisaṅghātasya vyapadeśaḥ | athānyenānyasya tadvyatiriktasya veti | anyenānya######sya vyapadeśaḥ | kasmāt | darśanasparśanābhyām ekārthagrahaṇāt | darśanena kaścid artho gṛhītaḥ sparśanenāpi so 'rtho gṛhyate yam aham adrākṣaṃ cakṣuṣā taṃ sparśanenāpi spṛśāmīti yaṃ cāspārkṣaṃ sparśanena taṃ cakṣuṣā paśyāmīti | ekaviṣayau dvāv #imau pratyayāv ekakart#ṛkau pratisandhīyete | na ca saṅghātakartṛkau nendriyeṇaikakartṛkau | tad yo 'sau cakṣuṣā tvagindriyeṇa caikārthasya### grahītā bhin#nanimit#tāv ananyakartṛkau pratyayau samānaviṣayau pratisandadhāti so 'rthāntarabhūta ātmeti | kathaṃ punar nendriyeṇaikakartṛkau | indriyaṃ khalu svaṃ svaṃ viṣayagrahaṇam ananyakartṛkaṃ pratisandhātum arhati nendriyāntarasya viṣayāntaragrahaṇam iti | kathaṃ na saṅghātakartṛkau | ekaḥ khalv ayaṃ bhin#nanimit#tau svātmakartṛkau pratyayau pratisaṃhitau vedayate na saṅghātaḥ | kasmāt | anivṛt#taṃ hi saṅghāte pratyekaṃ viṣayāntaragrahaṇasyāpratisandhānam indriyāntareṇeveti | "},
	score:  1041433,
}

func TestAlign(t *testing.T) {
	string1 := test.input[0]
	string2 := test.input[1]
	word1, word2, score := gonwr.Align([]rune(string1), []rune(string2), rune('#'), 1, -1, -1)
	if string(word1) != test.output[0] {
		t.Error(
			"expected", test.output[0],
			"got", word1,
		)
	}
	if string(word2) != test.output[1] {
		t.Error(
			"expected", test.output[1],
			"got", word2,
		)
	}
	if score != test.score {
		t.Error(
			"expected", test.score,
			"got", score,
		)
	}
	lentest := len(strings.Split(word1, "")) == len(strings.Split(word2, ""))
	if !lentest {
		t.Error(
			"expected", true,
			"got", false,
		)
	}
}
