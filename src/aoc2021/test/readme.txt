Remember:

- 'go test -bench=. ./src/aoc2021/test/day07_test.go -benchtime=10s'

- Due to issues with the test importing methods from the main-package:
-- Need to temporarily make the day##.go a non-main package (using same name as directory seems required too).
--- 'package aoc2021' at top of the file
- And due to issues from having "multiple packages" in the same directory (if you rename one day's package):
-- Need to move the file to its own directory OR copy it directly into this "test" directory (and make it 'package test' or whatever directory is named)

Can still run it as a main package (if you rename the package to 'package main' again),
just need to change the command to something like 'go run ./src/aoc2021/test/day07.go'

There is some issue with working-directory too, so copy ./cookie_session here.


Remember to return it to the aoc2021 folder afterwards, as a main package.
Also remember to delete the session-token file.