
#!/usr/bin/env python
import sys

# Kudos to https://github.com/taddeus/advent-of-code/blob/master/2020/19_regex.py
def solve(rules, messages):
    def expand(value):
        if not value.isdigit(): return value
        s = "(?:" + "".join(map(expand, rules[value].split())) + ")"
        print(value + ": " + s)
        return s

    print("final: " + expand("0"))
    expand("0")
    return 555


raw_rules, messages = sys.stdin.read().split("\n\n")
messages = messages.splitlines()
rules = dict(
    raw_rule.replace('"', "").split(": ", 1)
    for raw_rule in raw_rules.splitlines()
)

solve(rules, messages)
rules["8"] = "42 +"  # repeat pattern
rules["11"] = "(?P<R> 42 (?&R)? 31 )"  # recursive pattern
#print(solve(rules, messages))