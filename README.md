# pinfcked

According to Copilot: *The overall purpose of the program is to take a 4-digit PIN, find adjacent digits for each digit, and print all possible combinations of these adjacent digits.*

The scenario is when you effed up the key safe's new PIN and you need a list of potential PINs to brute-force open the safe again.

## Example

You believe to have coded the new PIN as `0000`, but it's just not working...

In general, there are 3\*3\*3\*3 = 81-1 combinations (-1 because of `0000` which is the one not working).

So, potential PIN candidates to try are:

`1000`
`0100`
`0010`
`0001`
`1100`
...
`9900`
`9990`
`9999`

`pinfcked` generates you this list.

## Usage

```
go get
go test
go build
./pinfcked 0000
```

### Debug

`./pinfcked -debug 1234`