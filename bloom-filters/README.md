## bloom filter
A bloom filter is basically a highly space efficent, probablistics data structure which is designed to
identify that and element exist already or not very quickly.
This can be used to check the million of username in a nanosecond.
# working
Its working mechanism is that you have along set of array which was initalized a zero value by default
we pass the data into `n` number of hash function with salted value we will make the every index of array to switch `1`
and index we will get from the hash function.
when we have to check whether that username is already taken or not so we will use the again pass the user name to `n` number of same hash function
and get the index and check to the array.
1. if all the value in array at index is 1 then we can say that it might have already taken the username by someone else.
2. if any of the value have the 0 then it is gurantee that its has not taken by the any user.

# formula
1. The probability of getting the false positive is 
    p = (1-e^(-kn/m))^k
    where k = number of the hash function
    n = number of inserted element
    m = number of bit in the array

# calculate (written in code)
k = 3
n = 1
m = 100000
p = (1-e^(3*1/100000))^3
p = 0.000000000000027
1 in 37 Trillion will occur