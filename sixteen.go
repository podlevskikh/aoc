// You can edit this code!
// Click here and start typing.
package main

// Для решения этой задачи подойдет встроенный пакет
// `strconv`.
import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var hexBinMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	"A": "1010",
	"B": "1011",
	"C": "1100",
	"D": "1101",
	"E": "1110",
	"F": "1111",
}

type Block struct {
	Version     string
	TypeID      string
	Value       int
	L           string
	Length      int
	Count       int
	SubPackages []SubPackage
	SubBlocks   []Block
}

type SubPackage struct {
	StartIndex int
	SubPack    string
}

func main() {
	starts := strings.Split(inputHex, ``)
	inBin := ""
	for _, h := range starts {
		inBin += hexBinMap[h]
	}
	fmt.Println(inBin)
	in := strings.Split(inBin, "")

	blocks := getBlocks(in)
	/*res := 0
	for _, b := range blocks {
		res += bToInt(b.Version)
	}*/
	//fmt.Println(res)
	fmt.Println(blocks)
}

func getBlocks(in []string) Block {
	b, rest := findBlocks(in)
	fmt.Println(rest)

	/*for _, b := range blocks {
		if b.TypeID != "100" {
			for _, subPackage := range b.SubPackages {
				inSub := in[subPackage.StartIndex:]
				subBlocks := getBlocks(inSub)
				blocks = append(blocks, subBlocks...)
			}
		}
	}*/
	return b
}

func findBlocks(in []string) (Block, []string) {
	v := strings.Join(in[0:3], "")
	t := strings.Join(in[3:6], "")

	if t == "100" {
		packs := []string{}
		for i := 0; i < len(in); i++ {
			pack := strings.Join(in[6+i*5:6+i*5+5], "")
			packs = append(packs, pack)
			if strings.Split(pack, "")[0] == "0" {
				res := ""
				for _, pack := range packs {
					res += strings.Join(strings.Split(pack, "")[1:], "")
				}
				resInt := bToInt(res)
				return Block{
					Version: v,
					TypeID:  t,
					Value:   resInt,
					L: strings.Join(in[6:], ""),
				}, in[6+i*5+5:]
			}
		}
	} else {
		blocks := []Block{}
		rest := []string{}

		l := in[6]
		if l == "1" {
			countPack := bToInt(strings.Join(in[7:18], ""))
			rest = in[18:]
			for j := 0; j < countPack; j++ {
				bs, r := findBlocks(rest)
				blocks = append(blocks, bs)
				rest = r
			}
		} else {
			lengthPack := bToInt(strings.Join(in[7:22], ""))
			subIn := in[22 : 22+lengthPack]
			b, sRest := findBlocks(subIn)
			blocks = append(blocks, b)
			for len(subIn) != len(sRest) && len(sRest) != 0 {
				subIn = sRest
				b, sRest = findBlocks(subIn)
				blocks = append(blocks, b)
			}
			rest = in[22+lengthPack:]
		}
		val := 0
		if t == "000" {
			for _, b := range blocks {
				val += b.Value
			}
		} else if t == "001" {
			val = 1
			for _, b := range blocks {
				val *= b.Value
			}
		} else if t == "010" {
			val = 0
			for _, b := range blocks {
				if val == 0 || val > b.Value {
					val = b.Value
				}
			}
		} else if t == "011" {
			val = 0
			for _, b := range blocks {
				if val == 0 || val < b.Value {
					val = b.Value
				}
			}
		} else if t == "101" {
			val = 0
			if blocks[0].Value > blocks[1].Value {
				val = 1
			}
		} else if t == "110" {
			val = 0
			if blocks[0].Value < blocks[1].Value {
				val = 1
			}
		} else if t == "111" {
			val = 0
			if blocks[0].Value == blocks[1].Value {
				val = 1
			}
		}
		//fmt.Println(val)
		return Block{
			Version: v,
			TypeID:  t,
			Value:   val,
			SubBlocks: blocks,
		}, rest
	}
	return Block{}, []string{}

}

func bToInt(b string) int {
	res := 0
	for i, s := range strings.Split(b, "") {
		sInt, _ := strconv.Atoi(s)
		res += sInt * int(math.Pow(2, float64(len(strings.Split(b, ""))-i-1)))
	}
	return res
}

//const inputHex = `D2FE28`
//const inputHex = `38006F45291200`
//const inputHex = `EE00D40C823060`
//const inputHex = `8A004A801A8002F478`
//const inputHex = `620080001611562C8802118E34`
///const inputHex = `C0015000016115A2E0802F182340`
//const inputHex = `A0016C880162017C3686B18A3D4780`
//const inputHex = `C200B40A82`
//const inputHex = `04005AC33890`
//const inputHex = `880086C3E88112`
//const inputHex = `CE00C43D881120`
//const inputHex = `D8005AC2A8F0`
//const inputHex = `F600BC2D8F`
//const inputHex = `9C005AC2F8F0`
//const inputHex = `9C0141080250320F1802104A08`

const inputHex = `020D78804D397973DB5B934D9280CC9F43080286957D9F60923592619D3230047C0109763976295356007365B37539ADE687F333EA8469200B666F5DC84E80232FC2C91B8490041332EB4006C4759775933530052C0119FAA7CB6ED57B9BBFBDC153004B0024299B490E537AFE3DA069EC507800370980F96F924A4F1E0495F691259198031C95AEF587B85B254F49C27AA2640082490F4B0F9802B2CFDA0094D5FB5D626E32B16D300565398DC6AFF600A080371BA12C1900042A37C398490F67BDDB131802928F5A009080351DA1FC441006A3C46C82020084FC1BE07CEA298029A008CCF08E5ED4689FD73BAA4510C009981C20056E2E4FAACA36000A10600D45A8750CC8010989716A299002171E634439200B47001009C749C7591BD7D0431002A4A73029866200F1277D7D8570043123A976AD72FFBD9CC80501A00AE677F5A43D8DB54D5FDECB7C8DEB0C77F8683005FC0109FCE7C89252E72693370545007A29C5B832E017CFF3E6B262126E7298FA1CC4A072E0054F5FBECC06671FE7D2C802359B56A0040245924585400F40313580B9B10031C00A500354009100300081D50028C00C1002C005BA300204008200FB50033F70028001FE60053A7E93957E1D09940209B7195A56BCC75AE7F18D46E273882402CCD006A600084C1D8ED0E8401D8A90BE12CCF2F4C4ADA602013BC401B8C11360880021B1361E4511007609C7B8CA8002DC32200F3AC01698EE2FF8A2C95B42F2DBAEB48A401BC5802737F8460C537F8460CF3D953100625C5A7D766E9CB7A39D8820082F29A9C9C244D6529C589F8C693EA5CD0218043382126492AD732924022CE006AE200DC248471D00010986D17A3547F200CA340149EDC4F67B71399BAEF2A64024B78028200FC778311CC40188AF0DA194CF743CC014E4D5A5AFBB4A4F30C9AC435004E662BB3EF0`
