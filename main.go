package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var numberWordPair = make(map[int]string)
/*
docodes the message from a .txt file
*/
func decode(file_path string)(string,error){
	//open the file
	file, err:=os.Open(file_path)
	if err != nil{
		return "",err
	}
	defer file.Close()

	//store the number-word pairs in a map
	numberWordPair,err = StoreLineWords(numberWordPair, file)
	if err != nil {
		return "",err
	}
	decoded_message:=PyramidDecoder(numberWordPair)
	return strings.Join(decoded_message," "),nil
}
/*
stores line words into a map[int]string
@params numberWordPair map[int]string
*/
func StoreLineWords(numberWordPair map[int]string,file *os.File)(map[int]string, error){
	scanner:=bufio.NewScanner(file)
	for scanner.Scan() {
		//split line into number and word
		line:=scanner.Text()
		parts:=strings.Fields(line)
		if len(parts)<2{
			continue
		}
		number, err:= strconv.Atoi(parts[0])
		if err != nil {
			return nil, err
		}
		word:= parts[1]
		numberWordPair[number]=word
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numberWordPair, nil
}

/*
pyramid decoding function
@params numberWordPair map[int]string
*/
func PyramidDecoder(numberWordPair map[int]string)[]string{
	var decodedMessage []string
	level := 1
	for i := 1; ;i+=level{
		if word, exists:=numberWordPair[i]; exists{
			decodedMessage = append(decodedMessage,word)
		}else{
			break
		}
		level++
	}
	return decodedMessage
}
func main() {
	file_path := "message.txt"
	message, err :=decode(file_path)
	if err != nil{
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(message)
}
