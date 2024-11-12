package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type configurations struct {
	maxDate       time.Time
	duration      time.Duration
	workDir       string
	fileException string
	isFolders     bool
}

func main() {
	cfg := newConfiguration()

	wd, err := os.ReadDir(cfg.workDir)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range wd {

		f := filepath.Join(cfg.workDir, file.Name())
		stat, err := os.Stat(f)
		if err != nil {
			continue
		}
		created := stat.ModTime()
		if cfg.maxDate.Sub(created) > time.Minute {
			if file.IsDir() && cfg.isFolders {
				err := os.RemoveAll(filepath.Join(cfg.workDir, file.Name()))
				if err != nil {
					log.Println("Удаление папки ", filepath.Join(cfg.workDir, file.Name()), " выдало ошибку: ", err)
				}
			} else if strings.HasSuffix(file.Name(), "."+cfg.fileException) {
				err := os.Remove(filepath.Join(cfg.workDir, file.Name()))
				if err != nil {
					log.Println("Удаление файла ", filepath.Join(cfg.workDir, file.Name()), " выдало ошибку: ", err)
				}
			}
		}

	}

}

func newConfiguration() configurations {
	var conf configurations
	var defauntDate time.Time
	var maxdate string
	var durationYear, durationMonth, durationWeek, durationDays, durationHours int
	getwd, err := os.Getwd()
	if err != nil {
		return configurations{}
	}

	flag.BoolVar(&conf.isFolders, "folder", false, "Работать с папками?")
	flag.StringVar(&maxdate, "date", "", "Дата создания после которой не будет ничего не будет удаляться")
	flag.StringVar(&conf.workDir, "workdir", getwd, "Рабочая директория")
	flag.StringVar(&conf.fileException, "fileException", "", "Расширения файлов")
	flag.IntVar(&durationYear, "durationYear", 0, "Возраст файла. Максимальное Количество лет.")
	flag.IntVar(&durationMonth, "durationMonth", 0, "Возраст файла. Максимальное Количество месяцев.")
	flag.IntVar(&durationWeek, "durationWeek", 0, "Возраст файла. Максимальное Количество недель.")
	flag.IntVar(&durationDays, "durationDays", 0, "Возраст файла. Максимальное Количество дней.")
	flag.IntVar(&durationHours, "durationHours", 0, "Возраст файла. Максимальное Количество часов.")

	flag.Parse()
	if maxdate != "" {
		conf.maxDate = stringDateAdapter(maxdate)
		if conf.maxDate == defauntDate {
			fmt.Println("Ошибка: Неправильная дата")
			os.Exit(1)
		}
		if !conf.isFolders && conf.fileException == "" {
			fmt.Println("Ошибка: Пустое поле расширения файла и не указан флаг для удаления папок")
			os.Exit(1)
		}
	} else {
		conf.maxDate = time.Now()
		conf.maxDate = conf.maxDate.Add(time.Duration(-durationYear) * time.Hour * 24 * 365)
		conf.maxDate = conf.maxDate.Add(time.Duration(-durationMonth) * time.Hour * 24 * 30)
		conf.maxDate = conf.maxDate.Add(time.Duration(-durationWeek) * time.Hour * 24 * 7)
		conf.maxDate = conf.maxDate.Add(time.Duration(-durationDays) * time.Hour * 24)
		conf.maxDate = conf.maxDate.Add(time.Duration(-durationHours) * time.Hour)
	}

	return conf
}

func stringDateAdapter(str string) time.Time {
	var separator string
	var err error
	var out time.Time
	timeStr := "23:59:58"
	out, err = time.Parse("2006-01-02 15:04:05", str)
	if err == nil {
		return out
	}
	for _, word := range []string{".", ",", "/", "-", ""} {
		if strings.Contains(str, word) {
			separator = word
			break
		}
	}
	strArr := strings.Split(str, separator)
	for i, s := range strArr {
		strArr[i] = strings.TrimSpace(strings.ToLower(s))
	}
	if len(strArr) == 3 && utf8.RuneCountInString(strArr[2]) >= 3 {
		separator = ""
		if strings.Contains(strArr[2], " ") {
			separator = " "
		} else if strings.Contains(strArr[2], "t") {
			separator = "t"
		}
		tempArr := strings.Split(strArr[2], separator)
		if len(tempArr) == 2 {
			timeStr = tempArr[1]
			strArr[2] = tempArr[0]
		}

	}
	if len(strArr) > 1 && len(strArr[1]) != 2 {
		symbolList := strings.Split(strArr[1], "")
		if len(symbolList) > 3 {
			mounthCode, OK := DictionaryMounth[strings.Join(symbolList[3:], "")]
			if OK {
				mounthCode++
				if mounthCode > 9 {
					strArr[1] = fmt.Sprintf("%c", mounthCode)
				} else {
					strArr[1] = fmt.Sprintf("%c", mounthCode)
				}
			} else {
				return time.Time{}
			}
		}
	}
	isInt := func(str string) bool {
		in, err := strconv.Atoi(str)
		return in < 13 || err == nil
	}
	if len(strArr) > 2 && utf8.RuneCountInString(strArr[1]) == 1 {
		strArr[1] = fmt.Sprintf("0%s", strArr[1])
	}
	if len(strArr) < 3 || str == "" {
		out = time.Now()

	} else if utf8.RuneCountInString(strArr[0]) == 4 && utf8.RuneCountInString(strArr[1]) == 2 && utf8.RuneCountInString(strArr[2]) == 2 {
		if !isInt(strArr[1]) {
			return out
		}
		out, _ = time.Parse("2006-01-02T15:04:05", fmt.Sprintf("%s-%s-%sT%s", strArr[0], strArr[1], strArr[2], timeStr))

	} else if utf8.RuneCountInString(strArr[0]) == 2 && utf8.RuneCountInString(strArr[1]) == 2 && utf8.RuneCountInString(strArr[2]) == 4 {
		if !isInt(strArr[1]) {
			return out
		}
		out, _ = time.Parse("02-01-2006T15:04:05", fmt.Sprintf("%s-%s-%sT%s", strArr[0], strArr[1], strArr[2], timeStr))

	}

	return out
}
