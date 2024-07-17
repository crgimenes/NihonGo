package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"
)

type Word struct {
	Japanese   string
	Meaning    string
	Romaji     string
	Difficulty int
}

var words = []Word{
	{"おっと", "my husband", "otto", 1},
	{"七", "7, seven, seventh", "nana", 1},
	{"あたまがいい", "smart", "atamagaii", 1},
	{"むすこ", "my son", "musuko", 1},
	{"むすめ", "my daughter", "musume", 1},
	{"六", "six", "roku", 1},
	{"つま", "my wife", "tsuma", 1},
	{"だいがくせい", "college student", "daigakusei", 1},
	{"ゆうめい", "famous", "yuumei", 1},
	{"かんごし", "nurse", "kangoshi", 1},
	{"おとうと", "my younger brother", "otouto", 1},
	{"うるさい", "noisy", "urusai", 1},
	{"五", "5, five", "go", 1},
	{"あに", "my older brother", "ani", 1},
	{"かいしゃいん", "office worker", "kaishain", 1},
	{"あね", "my older sister", "ane", 1},
	{"四", "four", "yon", 1},
	{"いもうと", "my younger sister", "imouto", 1},
	{"さい", "years old", "sai", 1},
	{"いそがしい", "busy", "isogashii", 1},
	{"エンジニア", "engineer", "enjiniya", 1},
	{"父", "dad", "chichi", 1},
	{"かぞく", "family", "kazoku", 1},
	{"母", "mom, mother, my mom", "haha", 1},
	{"まち", "town, wait, city", "machi", 1},
	{"とても", "very", "totemo", 1},
	{"きれい", "clean", "kirei", 1},
	{"しずか", "quiet", "shizuka", 1},
	{"にぎやか", "lively", "nigiyaka", 1},
	{"とし", "years, cities, city", "toshi", 1},
	{"日本", "Japan, Japanese", "nihon", 1},
	{"おおさか", "Osaka", "oosaka", 1},
	{"京都", "Kyoto", "kyouto", 1},
	{"ニューヨーク", "New York", "nyuuyooku", 1},
	{"も", "even, any, also", "mo", 1},
	{"私", "I", "watashi", 1},
	{"ようこそ", "welcome", "youkoso", 1},
	{"おはようございます", "good morning", "ohayougozaimasu", 1},
	{"かばん", "bag", "kaban", 1},
	{"田中", "Tanaka", "tanaka", 1},
	{"パスポート", "passport", "pasupooto", 1},
	{"の", "for, one, of", "no", 1},
	{"スマホ", "smartphone", "sumaho", 1},
	{"ちず", "map", "chizu", 1},
	{"きっぷ", "ticket", "kippu", 1},
	{"ちかてつ", "subway", "chikatetsu", 1},
	{"くうこう", "airport", "kuukou", 1},
	{"すみません", "excuse me", "sumimasen", 1},
	{"いま", "now, current", "ima", 1},
	{"一", "one, an, a", "ichi", 1},
	{"二", "two, 2", "ni", 1},
	{"三", "3, three, third", "san", 1},
	{"じ", "o'clock", "ji", 1},
	{"に", "to, in, as (a)", "ni", 1},
	{"いち", "one, an", "ichi", 1},
	{"おもしろい", "interesting, fun, funny", "omoshiroi", 1},
	{"その", "the, that, its", "sono", 1},
	{"かわいい", "cute, pretty", "kawaii", 1},
	{"この", "this, these", "kono", 1},
	{"くつ", "shoes, shoe", "kutsu", 1},
	{"しろい", "white", "shiroi", 1},
	{"コート", "coat", "kooto", 1},
	{"あかい", "red", "akai", 1},
	{"ぼうし", "hat, hats", "boushi", 1},
	{"そこ", "there", "soko", 1},
	{"だいがく", "university, college", "daigaku", 1},
	{"えき", "station, train station, stations", "eki", 1},
	{"デパート", "department store", "depaato", 1},
	{"ここ", "here", "koko", 1},
	{"バスてい", "bus stop", "basutei", 1},
	{"ホテル", "hotel", "hoteru", 1},
	{"コンビニ", "convenience store", "konbini", 1},
	{"どこ", "where", "doko", 1},
	{"ブラジルじん", "Brazilian", "burajirujin", 1},
	{"イギリスじん", "British", "igirisu jin", 1},
	{"ブラジル", "Brazil", "burajiru", 1},
	{"イギリス", "Britain, the United Kingdom, British", "igirisu", 1},
	{"にほんじん", "Japanese", "nihonjin", 1},
	{"カナダ", "Canada", "kanada", 1},
	{"アメリカじん", "American", "amerika jin", 1},
	{"カナダじん", "Canadian", "kanada jin", 1},
	{"アメリカ", "America, the USA, the US", "amerika", 1},
	{"ちいさい", "small, little", "chiisai", 1},
	{"にほん", "Japan", "nihon", 1},
	{"おおきい", "big, large", "ookii", 1},
	{"か", "or, is it?, that", "ka", 1},
	{"いいえ", "no", "iie", 1},
	{"はい", "yes", "hai", 1},
	{"ケーキ", "cake", "keeki", 1},
	{"ピザ", "pizza", "piza", 1},
	{"それ", "that (one), it, that", "sore", 1},
	{"これ", "this (one), these, this", "kore", 1},
	{"ラーメン", "ramen", "raamen", 1},
	{"おいしい", "good, delicious, tasty", "oishii", 1},
	{"は", "is, with, regarding", "wa", 1},
	{"カレー", "curry", "karee", 1},
	{"じゃあね", "bye", "jaane", 1},
	{"さん", "Mr., Miss, Mrs.", "san", 1},
	{"こんにちは", "hello", "konnichiwa", 1},
	{"がくせい", "student, students", "gakusei", 1},
	{"ひと", "person, people", "hito", 1},
	{"かっこいい", "cool, good-looking, stylish", "kakkoii", 1},
	{"べんごし", "lawyer, lawyers", "bengoshi", 1},
	{"やさしい", "kind, easy, nice", "yasashii", 1},
	{"せんせい", "teacher, Professor, teachers", "sensei", 1},
	{"いしゃ", "doctor, doctors", "isha", 1},
	{"と", "and, that, door", "to", 1},
	{"みず", "water", "mizu", 1},
	{"ごはん", "rice, meal, meals", "gohan", 1},
	{"すし", "sushi", "sushi", 1},
	{"ください", "please", "kudasai", 1},
	{"おちゃ", "green teas, green tea, tea", "ocha", 1},
}

func say(word string) {
	cmd := exec.Command("say", "-v", "Kyoko (Enhanced)", word)
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		word := words[rand.Intn(len(words))]

		say(word.Japanese)
		fmt.Printf("Japanese word: %s %q - Meaning: %s\n",
			word.Japanese,
			word.Romaji,
			word.Meaning)

		fmt.Print("Type the word in Japanese: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == word.Japanese {
			fmt.Println("Success! Correct word.")
			say(word.Japanese)
			continue
		}

		fmt.Printf("Failure! The correct word was: %s %s\n", word.Japanese, word.Romaji)
		say(word.Japanese)
	}
}
