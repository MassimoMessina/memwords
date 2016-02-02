package main

import (
		"flag"
		"fmt"
		"errors"
		"log"
 		"bufio"
 		"strconv"
 		"strings"
  		"os"
)

//Global Variable definition
var mapNumbers map[int]string
var aLines []string
var aConsonanti []string
var aDiscard []string

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  	}
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    aLines = append(aLines, scanner.Text())
 	}
  return aLines, scanner.Err()
}

// It will scan the words matching the rules for the specified number
func FindWords (number int) ([5000]string, int) {

	inxWords := 0
	var words [5000]string

	//log.Printf("memwords - FindWords number=%d\n", number)
	 
	val := mapNumbers[number]
	first, _ := strconv.Atoi (string(val[0]))
	second, _ := strconv.Atoi (string(val[1]))

	//init the discard box
	copy (aDiscard, aConsonanti)
	aDiscard[first]=" "
	aDiscard[second]=" "
	//fmt.Println(aDiscard)

	bWord_is_good:=false
	bFirst_is_good:=false

	var lastChar string
	//log.Printf("memwords first=%d second=%d, consonant=%s,%s\n", first, second, aConsonanti[first] ,aConsonanti[second])

	 
	//For all the words in the dictionary
	scan_aLines: for wi, word := range aLines {

	 	wordUpper := strings.ToUpper(word)
	 	//log.Print(wordUpper)
	 	bWord_is_good=false
	    bFirst_is_good=false

	 	// For all the characters in the word
	 	for _, c := range wordUpper {

	 		st := string(c)

	 		if st == string(aConsonanti[first]) {
	 			if bFirst_is_good {
	 				if lastChar==st {
	 					continue // it means that it is a double consonant
	 				} else {
	 					continue scan_aLines  // if you have another occurrence of first then discard the word
	 				    }
	 				
	 			}
	 			bFirst_is_good=true
	 			lastChar=st
	 			continue
	 		}

	 		if st == string(aConsonanti[second] ) {
	 			if  bWord_is_good {
	 				if lastChar==st {
	 					continue // it means that it is a double consonant
	 				} else {
	 					continue scan_aLines  // if you have another occurrence of first then discard the word
	 				    }
	 			}
	 			if bFirst_is_good {
	 				bWord_is_good=true
	 				lastChar=st
	 				continue
	 			}
	 			
	 			continue scan_aLines
	 		}

	 		lastChar=st
	 		
	 		for _ ,d := range aDiscard {
	 			//log.Printf("st=%s d=%s ; ", st, string(d))
	 			if (st==string(d)) {
	 				//log.Printf("scan interrotto per %s", wordUpper)
	 				continue scan_aLines
	 			}
	 		}

	 	} // end scan wordUpper

	 	if bWord_is_good {
	 		//log.Printf("memwords trovata %s\n", wordUpper)
	 		words[inxWords]=aLines[wi]
	 		inxWords++
	 	}
	 		

	}  // End scan aLines
	 
	 return words, inxWords

}


// Creates the map containing the association between Numbers 
// and its string representation with leading zeros
func CreateNumbersMap () (error) {

	 mapNumbers = make (map [int] string)

	 for i := 0 ; i < 100 ; i++ {

	 	n := fmt.Sprintf("%02d",i)
	 	mapNumbers[i]=n
	 	//log.Printf("%d (%s)\n", i, mapNumbers[i] )
	 } 
	 
	 log.Printf("memwords - Created mapNumbers with %d elements\n", len(mapNumbers))
	 return  errors.New("Completed SUccessfully")

}

// Initialize the arrays with constant data
func init() {
	
    log.Printf("MEMWORDS- (C) Massimo Messina \n")

	aConsonanti = append (aConsonanti, "Z" , "L", "N", "M", "R" , "F", "B", "T", "G", "P", "C", "D", "H", "Q", "S", "V", "K", "Y", "W", "X", "J")
	aDiscard = make ([]string, len(aConsonanti))
	fmt.Println(aConsonanti)

}

func main() {

	var srvItalian string

	// Initialize arrays with constants
	//InitializeArrays()
	

	// command line flags
    flag.StringVar(&srvItalian, "italian", "Italian.txt", "File name containing Italian Dictionary")
	
	flag.Parse()

	log.Printf("Server starting with italian:%s\n", srvItalian)

	aLines, err := readLines(srvItalian)
  	
  	if err != nil {
    	log.Fatalf("readLines: %s", err)
  	}
  	
  	//for i, line := range aLines {
    //	fmt.Println(i, line)
  	//}

  	log.Printf("memwords - read %d lines\n", len(aLines))

  	err = CreateNumbersMap()

    
  	// Look for all numbers between 0 and 99
  	for iii:=0; iii< 100; iii++ {

		    words, numWords := FindWords(iii)
		  	
		    //fmt.Println("words=",words)
		  	fmt.Printf("%s = ",mapNumbers[iii])

		  	for  i:=0; i< numWords; i++ {
		    	fmt.Printf("%s,", words[i])
		  	}
		  	fmt.Println(" ")  	
	}


  	log.Printf("memwords - Program terminated without errors\n")
}

