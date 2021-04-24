# sampleCoding
General practice code

!Logic behind the standalone program!

fetch the complete data from the long string file and store it

fetch each word from the dictionar file and try to compare if the word directly exists in the long string using strings.Conatins function from strings package

if not available check the length of the string fetched from the dictionary file and start checking for match of the first letter and last letter by length of the fecthed word,
if matching, then go for calling checkscrambleExistence function.

check scramble uses map and updates the characters available in the both passed strings
for example , if "SAAMPLE" is the word then we create a map by occurance of each letter in the word
map1[s] = 1
map1[a] = 2 // there are 2 a's in the word
map1[m] = 1
map1[p] = 1
map1[l] = 1
map1[e] = 1

same way build map for the second word passed by checking the start letter and end letter and length

compare maps using reflect.DeepEqual function.
if matches, it returns true with which we would increment the counter.



