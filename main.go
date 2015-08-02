package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

var levels = map[string]int{
	"major": 0,
	"minor": 1,
	"patch": 2,
	"x":     0,
	"y":     1,
	"z":     2,
}

func main() {
	level := flag.String("l", "patch", `Version part to increase - "major", "minor" or "patch"`)
	flag.Parse()

	sign := getGitConfigBool("autotag.sign")

	curVer := closestVersion()
	if curVer == "" {
		fmt.Println("unknown current version")
		os.Exit(1)
	}

	newVer := bumpVersion(curVer, levels[*level])
	args := []string{"tag", "-a", "-m", newVer}

	if sign {
		args = append(args, "-s")
	}
	args = append(args, newVer)

	fmt.Println(newVer)
	git(args...)
}

func git(args ...string) {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func getGitConfig(args ...string) string {
	args = append([]string{"config", "--get"}, args...)
	cmd := exec.Command("git", args...)
	bs, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(bs))
}

func getGitConfigBool(args ...string) bool {
	args = append([]string{"--bool"}, args...)
	return getGitConfig(args...) == "true"
}

func closestVersion() string {
	cmd := exec.Command("git", "describe", "--abbrev=0")
	bs, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(bs))
}

func bumpVersion(ver string, part int) string {
	prefix, parts := versionParts(ver)
	parts[part]++
	for i := part + 1; i < len(parts); i++ {
		parts[i] = 0
	}
	return versionString(prefix, parts)
}

func versionString(prefix string, parts []int) string {
	ver := fmt.Sprintf("%s%d", prefix, parts[0])
	for _, part := range parts[1:] {
		ver = fmt.Sprintf("%s.%d", ver, part)
	}
	return ver
}

// versionParts matches a px.y.z version, for non-digit values of p and digits
// x, y, and z.
func versionParts(s string) (prefix string, parts []int) {
	exp := regexp.MustCompile(`^([^\d]+)(\d+)\.(\d+)\.(\d+)$`)
	match := exp.FindStringSubmatch(s)
	if len(match) > 1 {
		prefix = match[1]
		parts = make([]int, len(match)-2)
		for i := range parts {
			parts[i], _ = strconv.Atoi(match[i+2])
		}
	}
	return
}
