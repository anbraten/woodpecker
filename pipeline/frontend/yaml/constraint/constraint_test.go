package constraint

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/woodpecker-ci/woodpecker/pipeline/frontend"
)

func TestConstraint(t *testing.T) {
	testdata := []struct {
		conf string
		with string
		want bool
	}{
		// string value
		{
			conf: "master",
			with: "develop",
			want: false,
		},
		{
			conf: "master",
			with: "master",
			want: true,
		},
		{
			conf: "feature/*",
			with: "feature/foo",
			want: true,
		},
		// slice value
		{
			conf: "[ master, feature/* ]",
			with: "develop",
			want: false,
		},
		{
			conf: "[ master, feature/* ]",
			with: "master",
			want: true,
		},
		{
			conf: "[ master, feature/* ]",
			with: "feature/foo",
			want: true,
		},
		// includes block
		{
			conf: "include: master",
			with: "develop",
			want: false,
		},
		{
			conf: "include: master",
			with: "master",
			want: true,
		},
		{
			conf: "include: feature/*",
			with: "master",
			want: false,
		},
		{
			conf: "include: feature/*",
			with: "feature/foo",
			want: true,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "develop",
			want: false,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "master",
			want: true,
		},
		{
			conf: "include: [ master, feature/* ]",
			with: "feature/foo",
			want: true,
		},
		// excludes block
		{
			conf: "exclude: master",
			with: "develop",
			want: true,
		},
		{
			conf: "exclude: master",
			with: "master",
			want: false,
		},
		{
			conf: "exclude: feature/*",
			with: "master",
			want: true,
		},
		{
			conf: "exclude: feature/*",
			with: "feature/foo",
			want: false,
		},
		{
			conf: "exclude: [ master, develop ]",
			with: "master",
			want: false,
		},
		{
			conf: "exclude: [ feature/*, bar ]",
			with: "master",
			want: true,
		},
		{
			conf: "exclude: [ feature/*, bar ]",
			with: "feature/foo",
			want: false,
		},
		// include and exclude blocks
		{
			conf: "{ include: [ master, feature/* ], exclude: [ develop ] }",
			with: "master",
			want: true,
		},
		{
			conf: "{ include: [ master, feature/* ], exclude: [ feature/bar ] }",
			with: "feature/bar",
			want: false,
		},
		{
			conf: "{ include: [ master, feature/* ], exclude: [ master, develop ] }",
			with: "master",
			want: false,
		},
		// empty blocks
		{
			conf: "",
			with: "master",
			want: true,
		},
	}
	for _, test := range testdata {
		c := parseConstraint(t, test.conf)
		got, want := c.Match(test.with), test.want
		if got != want {
			t.Errorf("Expect %q matches %q is %v", test.with, test.conf, want)
		}
	}
}

func TestConstraintList(t *testing.T) {
	testdata := []struct {
		conf    string
		with    []string
		message string
		want    bool
	}{
		{
			conf: "",
			with: []string{"CHANGELOG.md", "README.md"},
			want: true,
		},
		{
			conf: "CHANGELOG.md",
			with: []string{"CHANGELOG.md", "README.md"},
			want: true,
		},
		{
			conf: "'*.md'",
			with: []string{"CHANGELOG.md", "README.md"},
			want: true,
		},
		{
			conf: "['*.md']",
			with: []string{"CHANGELOG.md", "README.md"},
			want: true,
		},
		{
			conf: "'docs/*'",
			with: []string{"docs/README.md"},
			want: true,
		},
		{
			conf: "'docs/*'",
			with: []string{"docs/sub/README.md"},
			want: false,
		},
		{
			conf: "'docs/**'",
			with: []string{"docs/README.md", "docs/sub/README.md", "docs/sub-sub/README.md"},
			want: true,
		},
		{
			conf: "'docs/**'",
			with: []string{"README.md"},
			want: false,
		},
		{
			conf: "{ include: [ README.md ] }",
			with: []string{"CHANGELOG.md"},
			want: false,
		},
		{
			conf: "{ exclude: [ README.md ] }",
			with: []string{"design.md"},
			want: true,
		},
		// include and exclude blocks
		{
			conf: "{ include: [ '*.md', '*.ini' ], exclude: [ CHANGELOG.md ] }",
			with: []string{"README.md"},
			want: true,
		},
		{
			conf: "{ include: [ '*.md' ], exclude: [ CHANGELOG.md ] }",
			with: []string{"CHANGELOG.md"},
			want: false,
		},
		{
			conf: "{ include: [ '*.md' ], exclude: [ CHANGELOG.md ] }",
			with: []string{"README.md", "CHANGELOG.md"},
			want: false,
		},
		// commit message ignore matches
		{
			conf:    "{ include: [ README.md ], ignore_message: '[ALL]' }",
			with:    []string{"CHANGELOG.md"},
			message: "Build them [ALL]",
			want:    true,
		},
		{
			conf:    "{ exclude: [ '*.php' ], ignore_message: '[ALL]' }",
			with:    []string{"myfile.php"},
			message: "Build them [ALL]",
			want:    true,
		},
		{
			conf:    "{ ignore_message: '[ALL]' }",
			with:    []string{},
			message: "Build them [ALL]",
			want:    true,
		},
		// empty commit
		{
			conf: "{ include: [ README.md ] }",
			with: []string{},
			want: true,
		},
	}
	for _, test := range testdata {
		c := parseConstraintPath(t, test.conf)
		got, want := c.Match(test.with, test.message), test.want
		if got != want {
			t.Errorf("Expect %q matches %q should be %v got %v", test.with, test.conf, want, got)
		}
	}
}

func TestConstraintMap(t *testing.T) {
	testdata := []struct {
		conf string
		with map[string]string
		want bool
	}{
		{
			conf: "GOLANG: 1.7",
			with: map[string]string{"GOLANG": "1.7"},
			want: true,
		},
		{
			conf: "GOLANG: tip",
			with: map[string]string{"GOLANG": "1.7"},
			want: false,
		},
		{
			conf: "{ GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.1", "MYSQL": "5.6"},
			want: true,
		},
		{
			conf: "{ GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.0"},
			want: false,
		},
		{
			conf: "{ GOLANG: 1.7, REDIS: 3.* }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.0"},
			want: true,
		},
		{
			conf: "{ GOLANG: 1.7, BRANCH: release/**/test }",
			with: map[string]string{"GOLANG": "1.7", "BRANCH": "release/v1.12.1//test"},
			want: true,
		},
		{
			conf: "{ GOLANG: 1.7, BRANCH: release/**/test }",
			with: map[string]string{"GOLANG": "1.7", "BRANCH": "release/v1.12.1/qest"},
			want: false,
		},
		// include syntax
		{
			conf: "include: { GOLANG: 1.7 }",
			with: map[string]string{"GOLANG": "1.7"},
			want: true,
		},
		{
			conf: "include: { GOLANG: tip }",
			with: map[string]string{"GOLANG": "1.7"},
			want: false,
		},
		{
			conf: "include: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.1", "MYSQL": "5.6"},
			want: true,
		},
		{
			conf: "include: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.0"},
			want: false,
		},
		// exclude syntax
		{
			conf: "exclude: { GOLANG: 1.7 }",
			with: map[string]string{"GOLANG": "1.7"},
			want: false,
		},
		{
			conf: "exclude: { GOLANG: tip }",
			with: map[string]string{"GOLANG": "1.7"},
			want: true,
		},
		{
			conf: "exclude: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.1", "MYSQL": "5.6"},
			want: false,
		},
		{
			conf: "exclude: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.0"},
			want: true,
		},
		// exclude AND include values
		{
			conf: "{ include: { GOLANG: 1.7 }, exclude: { GOLANG: 1.7 } }",
			with: map[string]string{"GOLANG": "1.7"},
			want: false,
		},
		// blanks
		{
			conf: "",
			with: map[string]string{"GOLANG": "1.7", "REDIS": "3.0"},
			want: true,
		},
		{
			conf: "GOLANG: 1.7",
			with: map[string]string{},
			want: false,
		},
		{
			conf: "{ GOLANG: 1.7, REDIS: 3.0 }",
			with: map[string]string{},
			want: false,
		},
		{
			conf: "include: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{},
			want: false,
		},
		{
			conf: "exclude: { GOLANG: 1.7, REDIS: 3.1 }",
			with: map[string]string{},
			want: true,
		},
	}
	for _, test := range testdata {
		c := parseConstraintMap(t, test.conf)
		assert.Equal(t, test.want, c.Match(test.with), "config: '%s', with: '%s'", test.conf, test.with)
	}
}

func TestConstraints(t *testing.T) {
	testdata := []struct {
		desc string
		conf string
		with frontend.Metadata
		want bool
	}{
		{
			desc: "no constraints, must match on default events",
			conf: "",
			with: frontend.Metadata{
				Curr: frontend.Build{
					Event: frontend.EventPush,
				},
			},
			want: true,
		},
		{
			desc: "global branch filter",
			conf: "{ branch: develop }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush, Commit: frontend.Commit{Branch: "master"}}},
			want: false,
		},
		{
			desc: "global branch filter",
			conf: "{ branch: master }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush, Commit: frontend.Commit{Branch: "master"}}},
			want: true,
		},
		{
			desc: "repo constraint",
			conf: "{ repo: owner/* }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Repo: frontend.Repo{Name: "owner/repo"}},
			want: true,
		},
		{
			desc: "repo constraint",
			conf: "{ repo: octocat/* }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Repo: frontend.Repo{Name: "owner/repo"}},
			want: false,
		},
		{
			desc: "ref constraint",
			conf: "{ ref: refs/tags/* }",
			with: frontend.Metadata{Curr: frontend.Build{Commit: frontend.Commit{Ref: "refs/tags/v1.0.0"}, Event: frontend.EventPush}},
			want: true,
		},
		{
			desc: "ref constraint",
			conf: "{ ref: refs/tags/* }",
			with: frontend.Metadata{Curr: frontend.Build{Commit: frontend.Commit{Ref: "refs/heads/master"}, Event: frontend.EventPush}},
			want: false,
		},
		{
			desc: "platform constraint",
			conf: "{ platform: linux/amd64 }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Sys: frontend.System{Platform: "linux/amd64"}},
			want: true,
		},
		{
			desc: "platform constraint",
			conf: "{ repo: linux/amd64 }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Sys: frontend.System{Platform: "windows/amd64"}},
			want: false,
		},
		{
			desc: "instance constraint",
			conf: "{ instance: agent.tld }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Sys: frontend.System{Host: "agent.tld"}},
			want: true,
		},
		{
			desc: "instance constraint",
			conf: "{ instance: agent.tld }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Sys: frontend.System{Host: "beta.agent.tld"}},
			want: false,
		},
		{
			desc: "filter cron by default constraint",
			conf: "{}",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventCron}},
			want: false,
		},
		{
			desc: "filter cron by matching name",
			conf: "{ event: cron, cron: job1 }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventCron, Cron: "job1"}},
			want: true,
		},
		{
			desc: "filter cron by name",
			conf: "{ event: cron, cron: job2 }",
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventCron, Cron: "job1"}},
			want: false,
		},
		{
			desc: "no constraints, event gets filtered by default event filter",
			conf: "",
			with: frontend.Metadata{
				Curr: frontend.Build{Event: "non-default"},
			},
			want: false,
		},
		{
			desc: "filter if by event",
			conf: `{ if: 'CI_BUILD_EVENT == "push"' }`,
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}},
			want: true,
		},
		{
			desc: "filter if by event and repo",
			conf: `{ if: 'CI_BUILD_EVENT == "push" && CI_REPO == "owner/repo"' }`,
			with: frontend.Metadata{Curr: frontend.Build{Event: frontend.EventPush}, Repo: frontend.Repo{Name: "owner/repo"}},
			want: true,
		},
	}

	for _, test := range testdata {
		t.Run(test.desc, func(t *testing.T) {
			c := parseConstraints(t, test.conf)
			got, err := c.Match(test.with, false)
			if err != nil {
				t.Errorf("Match returned error: %v", err)
			}
			if got != test.want {
				t.Errorf("Expect %+v matches %q is %v", test.with, test.conf, test.want)
			}
		})
	}
}

func parseConstraints(t *testing.T, s string) *When {
	c := &When{}
	assert.NoError(t, yaml.Unmarshal([]byte(s), c))
	return c
}

func parseConstraint(t *testing.T, s string) *List {
	c := &List{}
	assert.NoError(t, yaml.Unmarshal([]byte(s), c))
	return c
}

func parseConstraintMap(t *testing.T, s string) *Map {
	c := &Map{}
	assert.NoError(t, yaml.Unmarshal([]byte(s), c))
	return c
}

func parseConstraintPath(t *testing.T, s string) *Path {
	c := &Path{}
	assert.NoError(t, yaml.Unmarshal([]byte(s), c))
	return c
}
