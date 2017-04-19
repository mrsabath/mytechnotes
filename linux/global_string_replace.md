# Global string replace:

```
BEFORE GLOBAL CHANGE
Diana-Arroyos-MacBook-Pro:temp darroyo$ find . -name "*.go" | xargs grep who
./0.go:who is your daddy?
./1/1.go:who's your daddy?
./2/2.go:who is your daddy?

GLOBAL CHANGE
Diana-Arroyos-MacBook-Pro:temp darroyo$ find . -name "*.go" | xargs sed "s/who is/who's/g"
who's your daddy?
who's your daddy?
who's your daddy?

AFTER GLOBAL CHANGE
Diana-Arroyos-MacBook-Pro:temp darroyo$ find . -name "*.go" | xargs grep who
./0.go:who is your daddy?
./1/1.go:who's your daddy?
./2/2.go:who is your daddy?
```
