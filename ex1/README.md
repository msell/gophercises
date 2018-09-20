
** Refactoring suggestions from initial submission **
1) rand.Seed usually goes in `func init()`
2) don't name things by their type (i.e. ptr to indicate a pointer or str to indicate a string)
3) if you're using a channel, instead of loading the entire file into memory and then looping over it with a `for range` in a goroutine I would have `getQuizData` receive a `chan problem` and for each loop it would send the problem over the channel.
4) `time.NewTimer(time.Duration(*timeoutPtr) * time.Second)` can be simplified as a channel with just `<- time.After((*timeoutPtr)*time.Second)` in the `select` (edited)
Durations are also usually represented as a string when parsing. So the command line flag should be something like `5s`, `20s`, `1m`, etc