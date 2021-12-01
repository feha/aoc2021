# aoc2021
Advent of Code 2021 - Time to learn Go
Will contain my solutions for aoc2021, so avoid reading the files in `.src/aoc2021/` unless you want spoilers.
Also has 2020-day1-part1 in `./src/aoc2020` for demo/example purposes, as it's less relevant than the current year (personally used it to set up and test the environment, without having to break open this years advent-calendar).

# Setup

Copy your session-cookie into .cookie_session or it wont work.
You can retrieve this string by:
- Logging into aoc and opening dev-console in your browser
- Selecting the "Application" tab
- Expanding "Cookies" in the left sidepanel ("Cookies" is found under header "Storage")
- Selecting the aoc url
- Then finally copy the Value from the "session"-row (will be a long lowercase hex-string)

# aoc.py and day-template

This project comes with a handy (but simple) utility I have made called aoc-py.
This can be used in the command-line to (more info in `python oac-py -h`):
1. Create a new file for the day by copying/'instantiating' the day.template file, renaming, and replacing placeholder values. (`python aoc.py [--year 2021] new [1-25]`)
    * 'year' - by default the current year. Mainly changes the folder used, and the `%{year}` placeholder when instantiating the template file.
    * [1-25] ('day') - by default finding the highest existing day/file (in ./src/aocYYYY/), incremented by 1. If no file exists its considered `0`, giving `0+1`. It overwrites the entire file if an existing day/file is specified. So beware losing data!
    * Normally used as `python aoc.py new`
    * The created file will be renamed using the format `./src/aoc{year}/day{day:02d}.go`
2. Test (run) a day (file) (`python aoc.py [--year 2021] test [1-25]`)
    * 'year' - by default the current year. Changes the folder used.
    * 'day' - by default uses the highest existing day/file (in ./src/aocYYYY/).
    * Normally used as `python aoc.py test`
    * It does not autmatically submit your answer to aoc, that has to be done manually. But generally the template file does include assertions for hte test-inputs, and prints the answer to the real input.

The day.template file can be edited to change the initial state of new days (files).
Generally should contain an easy way to add assertions for the tests aoc gives in each part (test-input with example-output).
There are various useful placeholder variables that can be used:
* `%{starttime}` - This is replaced with the timestamp (HH-MM-SS) of when the file was created. I usually put this in a comment to help me keep track of when I started a day for tracking-purposes (aoc's built-in timing is bugged, and assumes everyone starts at same time).
* `%{p1Done}` & `%{p2Done}` - UNIMPLEMENTED. Meant to be able to be replaced retroactively when a part is completed. Possibly with some delta's too, to effortlessly know how long you took. For now, you would have to edit those values manually.
* `%{year}` & `%{day}` - Self-explanatory. Replaced with the selected year (`YYYY`) and day (`DD`/`day:02d`) respectively.

# Installation

So to summarize, 'installing' the project could be done as such:
```console
?:/???> cd desired-project-directory

?:/desired-project-directory> git clone githubs-project-url .

# Get your session-cookie-token and write it into './cookie_session'

?:/desired-project-directory> python aoc.py new

# Edit the created file to complete part1

?:/desired-project-directory> python aoc.py --year 2020 test
argv: aoc.py
argv: --year
argv: 2020
argv: test
args  Namespace(day='', dir='./src/', func=<function test at 0x015F8F18>, year='2020')
Testing aoc: year=2020 - day=latest ...
go run ./src/aoc2020/day01.go
part1 minitest success: true!
part1: 63616

part2 minitest success: true!
part2:

# Edit the created file to complete part2

?:/desired-project-directory> python aoc.py -y 2020 test
argv: aoc.py
argv: --year
argv: 2020
argv: test
args  Namespace(day='', dir='./src/', func=<function test at 0x015F8F18>, year='2020')
Testing aoc: year=2020 - day=latest ...
go run ./src/aoc2020/day01.go
part1 minitest success: true!
part1: 63616

part2 minitest success: true!
part2: 67877784

?:/desired-project-directory> 
```
Note that I specified year 2020 in the block above. This is because I include a demo for the first day of the prior year in the project. Which allows you to test the environment without opening the actual/current year. Also serves to show how the project can contain/be used for multiple years at once.