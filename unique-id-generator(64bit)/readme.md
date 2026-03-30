# how maximum number is calculated 
 ``` (int64(-1)) ^ (int64(-1) << 5) ```
- (-1) in binary represent as all 1s.(111111....111111)
- << its left shiftwise operator means it shit the bit n number to left and adding n zero from the right to left.
- (-1) << 5 going to be 1111111...11100000. here we have pushed the 5 zero from the right.
- ^ its XOR operator means if bits are different then it will be 1 and if bit are same then its will be zero.
- (-1) ^ (1111111....11100000) which results to 0000000...00011111
- so the result is 11111 in binary when we convert it into the decimal then 1 x 2^4 + 1 x 2^3 + 1 x 2^2 + 1 x 2^1 + 1 x 2^0 = 16 + 8 + 4 + 2 + 1 = 31.