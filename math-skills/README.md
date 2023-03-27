# Math-Skills

This program read data from a file and print the result of each statistic mentioned above. In other words, this program read the data present in the path passed as argument.
The data in the file will be presented as the following example:

``` console
189
113
121
114
145
110
...
```



## How to Use
- Ensure you have Golang installed on your computer.
- Clone this repository ==> [math-skills](learn.zone01dakar.sn/git/bcoulibal/math-skills.git) .
- Open a terminal or command prompt in the directory where the program is located.
- Run the program by typing ```go run .``` followed by the the path of the file you want to analyze.
   For example,
   ```sh
   go run . datafile.txt
   ```
## Concepts and Formulas
### Variance
The variance measures how spread out a set of data is from its average value. It is calculated as follows:
```
Variance = Σ(xi - mean)^2 / N
```
where xi is each data point, mean is the average value of the data set, and N is the number of data points.

### Standard Deviation
The standard deviation is the square root of the variance. It measures how much the data deviates from its mean value. It is calculated as follows:

```
Standard Deviation = sqrt(Variance)
```
### Average
The average, also known as the mean, is the sum of all the data points divided by the number of data points. It is calculated as follows:
```
Average = Σxi / N
```
### Median
The median is the middle value in a sorted list of numbers. If there is an even number of values, the median is the average of the two middle values.
