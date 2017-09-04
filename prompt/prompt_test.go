package prompt

import (
	"bytes"
	"testing"

	"github.com/antham/strumt"

	"github.com/stretchr/testify/assert"
)

func TestPrompt(t *testing.T) {
	var stdout bytes.Buffer

	type test struct {
		userInput string
		scenario  []struct {
			inputs []string
			err    error
		}
		expected map[string]string
	}

	tests := []test{
		// Mandatory parameters only
		{
			"HEAD\nHEAD~2\n/home/project\nq\n",
			[]struct {
				inputs []string
				err    error
			}{
				{[]string{"HEAD"}, nil},
				{[]string{"HEAD~2"}, nil},
				{[]string{"/home/project"}, nil},
				{[]string{"q"}, nil},
			},
			map[string]string{
				"CHYLE_GIT_REFERENCE_FROM":  "HEAD",
				"CHYLE_GIT_REFERENCE_TO":    "HEAD~2",
				"CHYLE_GIT_REPOSITORY_PATH": "/home/project",
			},
		},
		// Matchers
		{
			"HEAD\nHEAD~2\n/home/project\n1\n1\nregular\n2\ntest.*\n3\njohn.*\n4\nsam.*\nm\nq\n",
			[]struct {
				inputs []string
				err    error
			}{
				{[]string{"HEAD"}, nil},
				{[]string{"HEAD~2"}, nil},
				{[]string{"/home/project"}, nil},
				{[]string{"1"}, nil},
				{[]string{"1"}, nil},
				{[]string{"regular"}, nil},
				{[]string{"2"}, nil},
				{[]string{"test.*"}, nil},
				{[]string{"3"}, nil},
				{[]string{"john.*"}, nil},
				{[]string{"4"}, nil},
				{[]string{"sam.*"}, nil},
				{[]string{"m"}, nil},
				{[]string{"q"}, nil},
			},
			map[string]string{
				"CHYLE_GIT_REFERENCE_FROM":  "HEAD",
				"CHYLE_GIT_REFERENCE_TO":    "HEAD~2",
				"CHYLE_GIT_REPOSITORY_PATH": "/home/project",
				"CHYLE_MATCHERS_MESSAGE":    "test.*",
				"CHYLE_MATCHERS_TYPE":       "regular",
				"CHYLE_MATCHERS_COMMITTER":  "john.*",
				"CHYLE_MATCHERS_AUTHOR":     "sam.*",
			},
		},
		// Matchers
		{
			"HEAD\nHEAD~2\n/home/project\n2\nid\nidParsed\n#\\d+\nq\n",
			[]struct {
				inputs []string
				err    error
			}{
				{[]string{"HEAD"}, nil},
				{[]string{"HEAD~2"}, nil},
				{[]string{"/home/project"}, nil},
				{[]string{"2"}, nil},
				{[]string{"id"}, nil},
				{[]string{"idParsed"}, nil},
				{[]string{"#\\d+"}, nil},
				{[]string{"q"}, nil},
			},
			map[string]string{
				"CHYLE_GIT_REFERENCE_FROM":   "HEAD",
				"CHYLE_GIT_REFERENCE_TO":     "HEAD~2",
				"CHYLE_GIT_REPOSITORY_PATH":  "/home/project",
				"CHYLE_EXTRACTORS_0_DESTKEY": "idParsed",
				"CHYLE_EXTRACTORS_0_ORIGKEY": "id",
				"CHYLE_EXTRACTORS_0_REG":     "#\\d+",
			},
		},
		// Decorators
		{
			"HEAD\nHEAD~2\n/home/project\n3\n1\nmessage\nid\n#\\d+\nhttp://test.com\n=@eTN#d0t:x4TgKE|XJ531!H<n0rJH\nobjectId\nfields.id\n1\ndate\nfields.date\nm\n3\n2\nmessage\nid\n#\\d+\nhttp://api.jira.com\nuser\npassword\nobjectId\nfields.id\n1\ndate\nfields.date\nm\n3\n3\nmessage\nid\n#\\d+\nd41d8cd98f00b204e9800998ecf8427e\nuser\nobjectId\nfields.id\n1\ndate\nfields.date\nm\n3\n4\necho\nmessage\nid\n4\necho\nmessage\nfield\nm\n3\n5\nTEST\ntest\n5\nfoo\nbar\nq\n",
			[]struct {
				inputs []string
				err    error
			}{
				{[]string{"HEAD"}, nil},
				{[]string{"HEAD~2"}, nil},
				{[]string{"/home/project"}, nil},
				{[]string{"3"}, nil},
				{[]string{"1"}, nil},
				{[]string{"message"}, nil},
				{[]string{"id"}, nil},
				{[]string{"#\\d+"}, nil},
				{[]string{"http://test.com"}, nil},
				{[]string{"=@eTN#d0t:x4TgKE|XJ531!H<n0rJH"}, nil},
				{[]string{"objectId"}, nil},
				{[]string{"fields.id"}, nil},
				{[]string{"1"}, nil},
				{[]string{"date"}, nil},
				{[]string{"fields.date"}, nil},
				{[]string{"m"}, nil},
				{[]string{"3"}, nil},
				{[]string{"2"}, nil},
				{[]string{"message"}, nil},
				{[]string{"id"}, nil},
				{[]string{"#\\d+"}, nil},
				{[]string{"http://api.jira.com"}, nil},
				{[]string{"user"}, nil},
				{[]string{"password"}, nil},
				{[]string{"objectId"}, nil},
				{[]string{"fields.id"}, nil},
				{[]string{"1"}, nil},
				{[]string{"date"}, nil},
				{[]string{"fields.date"}, nil},
				{[]string{"m"}, nil},
				{[]string{"3"}, nil},
				{[]string{"3"}, nil},
				{[]string{"message"}, nil},
				{[]string{"id"}, nil},
				{[]string{"#\\d+"}, nil},
				{[]string{"d41d8cd98f00b204e9800998ecf8427e"}, nil},
				{[]string{"user"}, nil},
				{[]string{"objectId"}, nil},
				{[]string{"fields.id"}, nil},
				{[]string{"1"}, nil},
				{[]string{"date"}, nil},
				{[]string{"fields.date"}, nil},
				{[]string{"m"}, nil},
				{[]string{"3"}, nil},
				{[]string{"4"}, nil},
				{[]string{"echo"}, nil},
				{[]string{"message"}, nil},
				{[]string{"id"}, nil},
				{[]string{"4"}, nil},
				{[]string{"echo"}, nil},
				{[]string{"message"}, nil},
				{[]string{"field"}, nil},
				{[]string{"m"}, nil},
				{[]string{"3"}, nil},
				{[]string{"5"}, nil},
				{[]string{"TEST"}, nil},
				{[]string{"test"}, nil},
				{[]string{"5"}, nil},
				{[]string{"foo"}, nil},
				{[]string{"bar"}, nil},
				{[]string{"q"}, nil},
			},
			map[string]string{
				"CHYLE_GIT_REFERENCE_FROM":                            "HEAD",
				"CHYLE_GIT_REFERENCE_TO":                              "HEAD~2",
				"CHYLE_GIT_REPOSITORY_PATH":                           "/home/project",
				"CHYLE_DECORATORS_CUSTOMAPIID_CREDENTIALS_TOKEN":      "=@eTN#d0t:x4TgKE|XJ531!H<n0rJH",
				"CHYLE_DECORATORS_CUSTOMAPIID_ENDPOINT_URL":           "http://test.com",
				"CHYLE_EXTRACTORS_CUSTOMAPIID_DESTKEY":                "id",
				"CHYLE_EXTRACTORS_CUSTOMAPIID_ORIGKEY":                "message",
				"CHYLE_EXTRACTORS_CUSTOMAPIID_REG":                    "#\\d+",
				"CHYLE_DECORATORS_CUSTOMAPIID_KEYS_0_DESTKEY":         "objectId",
				"CHYLE_DECORATORS_CUSTOMAPIID_KEYS_0_FIELD":           "fields.id",
				"CHYLE_DECORATORS_CUSTOMAPIID_KEYS_1_DESTKEY":         "date",
				"CHYLE_DECORATORS_CUSTOMAPIID_KEYS_1_FIELD":           "fields.date",
				"CHYLE_EXTRACTORS_JIRAISSUEID_ORIGKEY":                "message",
				"CHYLE_EXTRACTORS_JIRAISSUEID_DESTKEY":                "id",
				"CHYLE_EXTRACTORS_JIRAISSUEID_REG":                    "#\\d+",
				"CHYLE_DECORATORS_JIRAISSUE_ENDPOINT_URL":             "http://api.jira.com",
				"CHYLE_DECORATORS_JIRAISSUE_CREDENTIALS_USERNAME":     "user",
				"CHYLE_DECORATORS_JIRAISSUE_CREDENTIALS_PASSWORD":     "password",
				"CHYLE_DECORATORS_JIRAISSUE_KEYS_0_DESTKEY":           "objectId",
				"CHYLE_DECORATORS_JIRAISSUE_KEYS_0_FIELD":             "fields.id",
				"CHYLE_DECORATORS_JIRAISSUE_KEYS_1_DESTKEY":           "date",
				"CHYLE_DECORATORS_JIRAISSUE_KEYS_1_FIELD":             "fields.date",
				"CHYLE_EXTRACTORS_GITHUBISSUEID_ORIGKEY":              "message",
				"CHYLE_EXTRACTORS_GITHUBISSUEID_DESTKEY":              "id",
				"CHYLE_EXTRACTORS_GITHUBISSUEID_REG":                  "#\\d+",
				"CHYLE_DECORATORS_GITHUBISSUE_CREDENTIALS_OAUTHTOKEN": "d41d8cd98f00b204e9800998ecf8427e",
				"CHYLE_DECORATORS_GITHUBISSUE_CREDENTIALS_OWNER":      "user",
				"CHYLE_DECORATORS_GITHUBISSUE_KEYS_0_DESTKEY":         "objectId",
				"CHYLE_DECORATORS_GITHUBISSUE_KEYS_0_FIELD":           "fields.id",
				"CHYLE_DECORATORS_GITHUBISSUE_KEYS_1_DESTKEY":         "date",
				"CHYLE_DECORATORS_GITHUBISSUE_KEYS_1_FIELD":           "fields.date",
				"CHYLE_DECORATORS_SHELL_0_COMMAND":                    "echo",
				"CHYLE_DECORATORS_SHELL_0_ORIGKEY":                    "message",
				"CHYLE_DECORATORS_SHELL_0_DESTKEY":                    "id",
				"CHYLE_DECORATORS_SHELL_1_COMMAND":                    "echo",
				"CHYLE_DECORATORS_SHELL_1_ORIGKEY":                    "message",
				"CHYLE_DECORATORS_SHELL_1_DESTKEY":                    "field",
				"CHYLE_DECORATORS_ENV_0_VARNAME":                      "TEST",
				"CHYLE_DECORATORS_ENV_0_DESTKEY":                      "test",
				"CHYLE_DECORATORS_ENV_1_VARNAME":                      "foo",
				"CHYLE_DECORATORS_ENV_1_DESTKEY":                      "bar",
			},
		},

		// Senders
		{
			"HEAD\nHEAD~2\n/home/project\n4\n1\njson\n2\nd41d8cd98f00b204e9800998ecf8427e\nuser\nfalse\nRelease 1\nfalse\nv1.0.0\nmaster\n{{.}}\nfalse\nrepository\n3\nd41d8cd98f00b204e9800998ecf8427e\nhttp://test.com\nq\n",
			[]struct {
				inputs []string
				err    error
			}{
				{[]string{"HEAD"}, nil},
				{[]string{"HEAD~2"}, nil},
				{[]string{"/home/project"}, nil},
				{[]string{"4"}, nil},
				{[]string{"1"}, nil},
				{[]string{"json"}, nil},
				{[]string{"2"}, nil},
				{[]string{"d41d8cd98f00b204e9800998ecf8427e"}, nil},
				{[]string{"user"}, nil},
				{[]string{"false"}, nil},
				{[]string{"Release 1"}, nil},
				{[]string{"false"}, nil},
				{[]string{"v1.0.0"}, nil},
				{[]string{"master"}, nil},
				{[]string{"{{.}}"}, nil},
				{[]string{"false"}, nil},
				{[]string{"repository"}, nil},
				{[]string{"3"}, nil},
				{[]string{"d41d8cd98f00b204e9800998ecf8427e"}, nil},
				{[]string{"http://test.com"}, nil},
				{[]string{"q"}, nil},
			},
			map[string]string{
				"CHYLE_GIT_REFERENCE_FROM":                            "HEAD",
				"CHYLE_GIT_REFERENCE_TO":                              "HEAD~2",
				"CHYLE_GIT_REPOSITORY_PATH":                           "/home/project",
				"CHYLE_SENDERS_STDOUT_FORMAT":                         "json",
				"CHYLE_SENDERS_GITHUBRELEASE_CREDENTIALS_OAUTHTOKEN":  "d41d8cd98f00b204e9800998ecf8427e",
				"CHYLE_SENDERS_GITHUBRELEASE_CREDENTIALS_OWNER":       "user",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_DRAFT":           "false",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_NAME":            "Release 1",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_PRERELEASE":      "false",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_TAGNAME":         "v1.0.0",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_TARGETCOMMITISH": "master",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_TEMPLATE":        "{{.}}",
				"CHYLE_SENDERS_GITHUBRELEASE_RELEASE_UPDATE":          "false",
				"CHYLE_SENDERS_GITHUBRELEASE_REPOSITORY_NAME":         "repository",
				"CHYLE_SENDERS_CUSTOMAPI_CREDENTIALS_TOKEN":           "d41d8cd98f00b204e9800998ecf8427e",
				"CHYLE_SENDERS_CUSTOMAPI_ENDPOINT_URL":                "http://test.com",
			},
		},
	}

	for _, test := range tests {
		buf := test.userInput

		p := Prompts{}
		p.prompts = strumt.NewPromptsFromReaderAndWriter(bytes.NewBufferString(buf), &stdout)

		envs, err := p.Run()

		assert.NoError(t, err)
		assert.Equal(t, test.expected, map[string]string(*envs))

		for i, s := range test.scenario {
			if i+1 > len(p.prompts.Scenario()) {
				t.Fatal("Scenario doesn't match expected one")
			}

			assert.Equal(t, s.inputs, p.prompts.Scenario()[i].Inputs())
			assert.Equal(t, s.err, p.prompts.Scenario()[i].Error())
		}
	}
}
