package main

import (
	"fmt"
	"github.com/shiningrush/avatarbuilder"
	"image/color"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

func main() {
	f, _ := os.Create("test.txt")
	f.Write([]byte("haha"))
	name := "I "
	//name = "に ほんご／にっぽんご"
	//name = "조 선말"
	names := strings.Split(name, " ")
	fmt.Println(len(names))
	runes := make([]rune, 0)
	cnt := 0
	for i := len(names) - 1; cnt < 2 && i >= 0; i-- {
		cnt++
		runes = append(runes, []rune(names[i])[0])
	}
	runes[0] = []rune(names[len(names)-1])[0]
	runes[1] = []rune(names[0])[0]

	//r, _ := utf8.DecodeRuneInString(name)
	//fmt.Println(fmt.Sprintf("https://twemoji.maxcdn.com/36x36/%x.png", r))
	//  SourceHanSansSC.ttf Mirza-Medium.ttf
	ab := avatarbuilder.NewAvatarBuilder("/Users/xiongchuanhong/Mirza-Medium.ttf", &FontMedium{})
	ab.SetFrontgroundColor(color.White)
	ab.SetBackgroundColor(getRandomRGBA())
	ab.SetFontSize(100)

	ab.SetAvatarSize(200, 200)
	if err := ab.GenerateImageAndSave(string(runes), "./outCn.png"); err != nil {
		fmt.Println(err)
		return
	}
}

type ArOneFontMedium struct{}

func (a *ArOneFontMedium) CalculateCenterLocation(s string, ab *avatarbuilder.AvatarBuilder) (x int, y int) {
	x = 50
	y = ab.H / 2
	return
}

type FontMedium struct{}

func (f *FontMedium) CalculateCenterLocation(s string, ab *avatarbuilder.AvatarBuilder) (x int, y int) {
	cr := regexp.MustCompile("[\u4e00-\u9FA5,\u3040-\u31FF]{1}")
	er := regexp.MustCompile("[a-zA-Z]{1}")
	nr := regexp.MustCompile("[0-9]{1}")
	//ar := regexp.MustCompile("[؀-ۿ]{1}")
	ar := regexp.MustCompile("[\u0600-\u06FF]{1}")

	cCount := len(cr.FindAllStringSubmatch(s, -1))
	eCount := len(er.FindAllStringSubmatch(s, -1))
	nCount := len(nr.FindAllStringSubmatch(s, -1))
	arCount := len(ar.FindAllStringSubmatch(s, -1))

	x = ab.W/2 - (cCount*ab.GetFontWidth()+eCount*ab.GetFontWidth()*3/5+nCount*ab.GetFontWidth()*3/5+arCount*ab.GetFontWidth()*3/5)/2
	y = ab.H/2 + ab.GetFontWidth()*4/11
	return
}

func getRandomRGBA() color.RGBA {
	rand.Seed(time.Now().Unix())
	return color.RGBA{
		R: uint8(100 + rand.Intn(100)),
		G: uint8(100 + rand.Intn(100)),
		B: uint8(100 + rand.Intn(100)),
		A: 255,
	}
}
