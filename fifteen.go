// You can edit this code!
// Click here and start typing.
package main

// Для решения этой задачи подойдет встроенный пакет
// `strconv`.
import (
	"fmt"
	"strconv"
	"strings"
)

type SumRiskMap struct {
	Map map[int]map[int]int
	I int
	J int
}

func NewSumRiskMap(start, i, j int) *SumRiskMap {
	m := make(map[int]map[int]int)
	for k := 0; k < i; k++ {
		ml := make(map[int]int)
		for l := 0; l < j; l++ {
			ml[l] = 0
		}
		m[k] = ml
	}
	m[0][0] = start
	return &SumRiskMap{Map: m, I: i, J: j}
}

func (s *SumRiskMap) process(risks [][]int) bool {
	changed := false
	for i := 0; i < s.I; i++ {
		for j := 0; j < s.J; j++ {
			min := s.getMinWay(i, j, risks[i][j])
			if min > 0 && (min < s.Map[i][j] || s.Map[i][j] == 0) {
				s.Map[i][j] = min
				changed = true
			}
		}
	}
	return changed
}

func (s *SumRiskMap) getMinWay(i, j, v int) int {
	min := -1
	if c, ok := s.Map[i-1]; ok {
		if c[j] != 0 {
			min = c[j] + v
		}
	}
	if c, ok := s.Map[i+1]; ok {
		if c[j] != 0 {
			if min == -1 || min > c[j] + v {
				min = c[j] + v
			}
		}
	}
	if c, ok := s.Map[i][j-1]; ok {
		if c != 0 {
			if min == -1 || min > c + v {
				min = c + v
			}
		}
	}
	if c, ok := s.Map[i][j+1]; ok {
		if c != 0 {
			if min == -1 || min > c + v {
				min = c + v
			}
		}
	}
	return min
}

func (s *SumRiskMap) print() {
	for i := 0; i < s.I; i++ {
		l := ""
		for j := 0; j < s.J; j++ {
			l += " " + strconv.Itoa(s.Map[i][j])
		}
		fmt.Println(l)
	}
}

func main() {
	starts := strings.Split(input_map, `
`)
	in := [][]int{}
	for i := 0; i < 5; i++ {
		for _, start := range starts {
			in_line := []int{}
			for j := 0; j < 5; j++ {
				for _, s := range strings.Split(start, "") {
					in_c, _ := strconv.Atoi(s)
					c := in_c + i + j
					if c >= 10 {
						c = c - 9
					}
					if c >= 10 {
						c = c - 9
					}
					in_line = append(in_line, c)
				}
			}
			in = append(in, in_line)
		}
	}
	fmt.Println(in)

	m := NewSumRiskMap(in[0][0], len(in), len(in[0]))
	m.print()
	fmt.Println("")
	changed := m.process(in)
	m.print()
	fmt.Println("")
	for changed {
		changed = m.process(in)
		m.print()
		fmt.Println("")
	}

}

/*const input_map = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`*/

const input_map = `5698271399793414121494916897678273334792458352199237144928619653191794144917556817941885991921787155
5661218626919981949139858312691785612259857529334128895157352159114297128511813456894839991981828919
1829946539172173769987529428979213588624889988719998181167931129997731961946189278688942251875931179
1999119539431115938818955166792951461581819916221516821681692884358864981966117818421936138119178448
8248487982679211958329811353187936515794549929583184191911666793879132359192179796818752613119298188
7391613611139499971186919159397755657961352334771192112294557688911149919935819422199281134255797767
8124729396114568961156298172414395681961627813193215526122172712983935817479529728239799416288287177
9649211889945273498525484578866135611815817899936588391937238711562851327977738262898917714539618992
1141741151396594515311116985514997126571483473844951949329748922961996548324126664446185813879832188
6514441248746823748338931782958583948711212951974131921212693782738511196996128897191782852921157717
9391886829151987712527174199971418888695457266233516772351113144394198243117872737133772919969819519
1198331423874598397283651925996811591794589456846981969193281925566799829729889189889897828398875973
3993119331859999411778999198997168314599391886293921999193929916292896686191667979558511312556596894
4885648184914319438678987185948358172121994124174189148399374584265796556987819754991475787318732495
7161834654516953633561499967999394938799899319729924222599794975138255185728879581989683917238688499
5984849949933778929823916964338381451782979977879923213397398539961494556984787412996655398495998949
5715125898984926281417488864215481241818325728147899985999237792243119997919523716364219919384984139
9859999616929285298541192766179418656994667513839176241589982149119397731438826191347741844116719249
8175137179841286592241832831664588945248987741981849615626887738291971771877815459199919596297762924
1195455998712612867139126669399111711242319181378472389957511612195887919281465911167399661193991259
6683294289715197916912936717981879539595921479689377149531198421412217683129512395995551899999253636
4384549778929184387719457231348656645896889888925317219969479184937844937194733193727596993257773958
3276812228776279165124713927385779131439398642869891821824228691957821333121694837396724995969152871
9919942131189697791118151719399685129324814276422561998279371918941781591744772893513849374159719689
9381491523919433281119518148692169889399181847183427693123932741321897894329241437714386791132191185
9491214195479284697132149999835916255594317799411165638929858419125934487114527249742829613155512959
6311661999995998283745198816594176551398793726177951929118916489667398222393341857119652996537191216
6982112288787347411291733589694945315895949612819676734199121729449261277878511178985354939268994351
3674562711198699238983994755221985299338191681822892997159181992896128148896975971957144448941814152
9899194386288999913747123919122811877922937176682999664964885597922734185747279621687621731836928862
2116816742597199191219978826825994722615889928911124939448369538185829869273435191136937156632583851
4349941568369745619775621261571293998779732985286699593958874613332787391776856914459351828283129499
2139545889871837989617135399996761155985919219131927394712919919165898258917812865693299999372143937
2797918178177159361299151932139531949921215449467992544898132943871367946184259699282519985747111894
1486233998921528551532861544482468798798911836599829829996424887757819971589876285997126989165281171
8861164846365591334792621191572483286587764653747651579926829499721832813168648324199795414996798928
9291853568543999152479124588672419544474748794822216282929683566381835292497918418175413139195238875
4459193479395871171969518189835778911861995829613879912466644113246996925946217134981986998535612761
3917415977315473896693284392244697742378915475693115524188981984797471519797297767819983299114674163
9479997975888641554837851979187615719199134617129683999813865847829191959434496893219481276491982115
1192917311668797426356261998346738779982194786382977926498879633266998527258133639952619989139917922
7719288735263969519584635229174517995733918532265569199269899656492443157891869119696416131629593639
2184119869185523966813259979736987537578526292398798983845392291835172615194758136222947294271887893
8729953944813191318117475998726288371635861354847924941923488971912998769489173795899411912899358291
4865718997911481411828882919717899614955999831128828316982172988928635967691754477941769257918149225
2993917151292413828544861759731699793999794939134673192415919816989946192187733812832551933836349353
2392657723121896915129994993779197965282385759494924426239478311328999729999373913937519813997858629
1231999494818876979724169419419112419866229157682954926121898668165939174414423891657921992944249919
9918532449129634933779591992923399274618345391686218726361524731393799762548689548319471496829728796
9998799217976696793819516835698899484118485923118178917448293998592525822132531922781281129488216397
5192761493599676289836184964748349843714842349297638881284681117695677295987225931879393958925236789
7268883329987412183226822596174125816931195839329615165616819294836862793251385934133841918139897811
6965568945376938749193143429716268629281259918717113392767177211856962116259137767717711199819531318
9946637998921258586362681893277981845668197589625911332226897948938743394789819211693113189244919914
4871851551891469283341781775791494521158119524555886558545149797364636497391519989229278417113895118
1345983432979827816668993821658815763434864899249549719186967999197848321762371279955891998781349927
6314411929798522969667512333572214219998191666968187141597233789569459669912829965921239279371936316
3992169277898367859361361786747537912321799661448888219875991399581567758199888738334192143862413181
2222496914541859183177213262919191272991899639234997678238191986198749293843119619698119988293199525
5915371189328353219855382685759969775542849935211982181425259529898186916649392484997392838711191117
8169587973671717318581118999998425389847631991852491196188179989154396999942138113786691868241134889
5344685492577218115877995132456423826951924968919925139319849473936971671795487241479296442941929811
1251739861149952116921832956986111819961136698717996819282184146866727525158711738228839674622782386
9847314199798528792466291586899159199338477391961835887792183113688154215493472212798899991864295742
5257272784491799827451173416779878548132798785993631488174276918175531348892491918523719192299467898
7992979991181179831169821468119411577879187992796818917682661441819267843277919392835983139489719471
4128797488218887928681673379728116122642549998949328966228125433858567196198751462818891984938354911
3475959599312382511869439189922686358873981983568449243997159614192913999914236993319949291439999198
1517571791296243999998371134995497132685919198842128735596841665424591185199389913668775951218689423
5654967112585715992259299352411729124559287912759919919348673481869219527276935916149899615111454338
7959995196194629598269761799738458135587969367688876114561929169848572742199927976195348111813159694
8341237826999851918267518141919121565937312914613679117869654472139112992413357268683764935191892197
9191144725134452491124119549826718699739659572818352415387359897217137915122751995997915493159884984
6114992423997191951469189895299879231999121815923893319995434539493649119619119321913937699983662725
8691733293139727547586431961128887545859836374323612455831994799481883935133399817678718165893368111
9561258499233959148319937411137559674714812297689646923979591185916186723579892829686891551677119835
3817583836164917989158856146991679117129286763822991921599939389559162988815144775134949221893987156
8869393821317962938731194192326717912615937839319191487729718951863811937788448721968368554499469998
9779931174618997891716996479823339528914111348193299619911134519591869598192982377992849895786861628
9228849524211171112561193761591642939783616849941114962751144449512592397581717151998199129916399569
3681981471397593418983945941198912291467516143121382229831485899639873221948653399284748979865113829
4511319899177843659998889999335128559884168939718989492987161554392999453863694462822117736248363914
6889791329888928324783251893783438993799912282811711691692819391686821652429416117277191688919199791
1813466918799319576952728154162197258151141791197279928841612511822193355919372971629299829231292443
6929796618267421196847896679863622321769172269232621692961879886142891298588221949715178615764931656
2181488467213697845251891984157371193949716425395214118928493219545175866269193351849171839139818194
9145886424872389992784162159124115362691198296549871231961989118911915871253826391146992645811797217
7261813715196955953539695411269291743513585462916998619485955985198911978299891114856577679915119891
2395375323918811172993132575669768211797188142297317469339437887956823821249292252996118812429293966
3493698117398613593295939911464991949189591628859269618891373198822876449981519279397383938445938961
7188399498811299474499795927858692591488691212476189948484913298291739588851722167749543818816918814
6639711919811875921558499577359957913939617544713997397137148176493159178897792388987517214382594658
9868637769618335566138613187813295523435485163161618919397998236266915715199319761642818478856959385
1181969758279137971179348866926365624188612783894276496899793241258119122436174649949517195172514793
6282561199121299888181471164167212292274611428169718849631799674421998399218169533259817922649993996
7892189127116936145295482897726952912496729194898291625689311833439971918991813338319988542899569332
1984849148728139499199899219892824339819797636941197152362487518771542766624971973432971739847999929
1389955743198821981998929159786983988371511612124611537771554447966118799114165127857297139873268292
7114948568589988719191915583826591117862589116475856191177131841299229935172958819245872994369999426
4275554996585879411196552781973718511282969951492149869419418957345666452981882879663949917972311592`
