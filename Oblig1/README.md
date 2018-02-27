Obligatorisk oppgave 1

# 1   

Binær|Hexa|Deci
   ---|---|---
1101|	0xD|	13
110111101010|	0xDEA|	3562
1010111100110100|	0xAF34|	44852
1111111111111111|	0xFFFF|	65535
00010001011110001010|	0x1178A|	71562

## 1a)

### Binær til hex og motsatt

Ved å bruke denne tabellen, blir prosessen å konvertere binært til hexa og motsatt.

Decimal|Hexadecimal|Binær
---|---|---
0|0|0
1|1|1
2|2|10
3|3|11
4|4|100
5|5|101
6|6|110
7|7|111
8|8|1000
9|9|1001
10|A|1010
11|B|1100
12|C|1100
13|D|1101
14|E|1110
15|F|1111 

Som eksempel å konvertere 1111001101011010 deler man det opp i 4
og konverter. 1111=F    0011=3 	0101=5 	 1010=A  
1111001101011010=F35A

Å gå fra hexa til binært kan man følge den samme prosessen bare motsatt 

Som eksempel
0xDEA=	 D=1101	E=1110	A=1010	=1101111010101

### Binær til Decimal

Å gå fra binært til decimal må man utføre litt mer matteregning. 
Først må man gange hvert av sifrene med 2 og opphøyd i plassen tallet står på. 
Som eksempel bruker vi tallet 110010110.  
(1\*2^8)+(1\*2^7)+(0\*2^6)+(0\*2^5)+(1\*2^4)+(0\*2^3)+(1\*2^2)+(1\*2^1)+(0\*2^0)=406

### Decimal til Binær

For å gå fra fra decimal til binær kan man sette opp denne typen tabell
Som eksempel bruker vi tallet 156. 

n |128|64|32|16|8|4|2|1
--|--|--|--|--|--|--|--|--|
156-128=28|1
28<64||0
28<32|||0
28-16=12||||1
12-8=4|||||1
4-4=0||||||1
0>2|||||||0
0>1||||||||0

156 = 10011100

## 1b)

### Hex til decimal
Å gå fra hex til decimal ganger man med 16 potens  
Vi bruker det samme eksempelet som tidligere 0xDEA.  
D(13)E(14)A(10)=(13\*16^2)+(14\*16^1)+(10\*^16^0)  
13\*256+14\*16+10\*1  
3328+224+10=35620  

### Decimal til hex
Når man skal konvertere fra decimal til hex må man utføre en delingsprosess med 16 eller kan man gå via binær og  
ende i hex.  
Med eksempel bruker vi 100decimal  
Da deler man på 2 og skriver 1 hvis det ikke kan deles på 2.  
2:100=50    0  
2:50=25     0  
2:25=12     1  
2:6=3       0  
2:3=1       0  
2:1=0		1  
0			1	  
Deretter snur man sekvensen slik 0010011 = 01100100  
Så kan man separere binarystingen  
0110=6  
0100=4  
100deci=0x64  



# 2

## 2a) og 2b)

Se [sorting.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/algorithms/sorting.go) og [sorting_test.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/algorithms/sorting_test.go)

## 2c)

Benchmark resultater:  
<img src="https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/cmd_benchmark.png">
<img src="https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/benchmark_graph.png">

# 3

Loop: [loop.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/loop/loop.go)  
<img src="https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/loop_cpuandmemory.png">

# 4

## 4a)

Hverken GoLand eller Windows cmd vil printe karakterer for 0x80 til 0x9F (det blir firkanter, spørsmålstegn eller lignende). Resten av karakterene i extended ASCII blir printet ut.

Se [ascii.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/ascii/ascii.go)

## 4b)

Se [ascii.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/ascii/ascii.go)

## 4c)

Se [iso_test.go](https://github.com/TobiasAlbert123/midgets/blob/master/Obligatorisk1/ascii/iso_test.go)
