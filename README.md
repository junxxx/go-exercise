#go-exercise

## Exercise
1.1 Modify the `echo` program to also print os.Args[0], the name of the command that invoked it.

1.2 Modify the `echo` program to print the index and value of each of its arguments, one per line.

1.3 Experiment to measure the difference in running time between our potentially inefficient versions and the one that uses `strings.Join`

1.4 Modify `dup2` to print the names of all files in which each duplicated line occurs

1.7 The function call `io.Copy(dst, src)` reads from src and writes to dst. Use it instead of `ioutil.ReadAll` to copy the response body to `os.Stdout` without requiring a buffer large enough to hold the entire stream. Be sure to check the error result of `io.Copy`
