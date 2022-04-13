package applylist

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jrcichra/kube-applier/sysutil"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	rawList           []string
	repoPath          string
	blacklistPath     string
	whitelistPath     string
	fs                sysutil.FileSystemInterface
	expectedApplyList []string
	expectedBlacklist []string
	expectedErr       error
}

// TestPurgeComments verifies the comment processing and purging from the
// whitelist and blacklist specifications.
func TestPurgeComments(t *testing.T) {

	var testData = []struct {
		rawList        []string
		expectedReturn []string
	}{
		// No comment
		{
			[]string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
			[]string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
		},
		// First line is commented
		{
			[]string{"#/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
			[]string{"/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
		},
		// Last line is commented
		{
			[]string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "# /repo/c.json"},
			[]string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c"},
		},
		// Empty line
		{
			[]string{"/repo/a/b.json", "", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
			[]string{"/repo/a/b.json", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
		},
		// Comment line only containing the comment character.
		{
			[]string{"/repo/a/b.json", "#", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
			[]string{"/repo/a/b.json", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"},
		},
		// Empty file
		{
			[]string{},
			[]string{},
		},
		// File with only comment lines.
		{
			[]string{"# some comment "},
			[]string{},
		},
	}

	assert := assert.New(t)
	mockCtrl := gomock.NewController(t)
	fs := sysutil.NewMockFileSystemInterface(mockCtrl)
	f := &Factory{"", "", "", fs}
	for _, td := range testData {

		rv := f.purgeCommentsFromList(td.rawList)
		assert.Equal(rv, td.expectedReturn)
	}
}

func TestFactoryCreate(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	fs := sysutil.NewMockFileSystemInterface(mockCtrl)

	// ReadLines error -> return nil lists and error
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return(nil, fmt.Errorf("error")),
	)
	tc := testCase{[]string{}, "/repo", "/blacklist", "/whitelist", fs, nil, nil, fmt.Errorf("error")}
	createAndAssert(t, tc)

	// ReadLines error -> return nil lists and error
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return(nil, fmt.Errorf("error")),
	)
	tc = testCase{[]string{}, "/repo", "/blacklist", "/whitelist", fs, nil, nil, fmt.Errorf("error")}
	createAndAssert(t, tc)

	// Single .json file, empty whitelist, empty blacklist -> file in applyList
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	tc = testCase{[]string{"/repo/a.json"}, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a.json"}, []string{}, nil}
	createAndAssert(t, tc)

	// Single .yaml file, empty blacklist empty whitelist -> file in applyList
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	tc = testCase{[]string{"/repo/a.yaml"}, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a.yaml"}, []string{}, nil}
	createAndAssert(t, tc)

	// Single non-.json & non-.yaml file, empty blacklist empty whitelist
	// -> file not in applyList
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	tc = testCase{[]string{"/repo/a"}, "/repo", "/blacklist", "/whitelist", fs, []string{}, []string{}, nil}
	createAndAssert(t, tc)

	// Multiple files (mixed extensions), empty blacklist, emptry whitelist
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	rawList := []string{"/repo/a.json", "/repo/b.jpg", "/repo/a/b.yaml", "/repo/a/b"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a.json", "/repo/a/b.yaml"}, []string{}, nil}
	createAndAssert(t, tc)

	// Multiple files (mixed extensions), blacklist, empty whitelist
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{"b.json", "b/c.json"}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	rawList = []string{"/repo/a.json", "/repo/b.json", "/repo/a/b/c.yaml", "/repo/a/b", "/repo/b/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a.json", "/repo/a/b/c.yaml"}, []string{"/repo/b.json", "/repo/b/c.json"}, nil}
	createAndAssert(t, tc)

	// File in blacklist but not in repo
	// (Ends up on returned blacklist anyway)
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{"a/b/c.yaml", "f.json"}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{}, nil),
	)
	rawList = []string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a/b.json", "/repo/c.json"}, []string{"/repo/a/b/c.yaml", "/repo/f.json"}, nil}
	createAndAssert(t, tc)

	// Empty blacklist, valid whitelist all whitelist is in the repo
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{"a/b/c.yaml", "c.json"}, nil),
	)
	rawList = []string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a/b/c.yaml", "/repo/c.json"}, []string{}, nil}
	createAndAssert(t, tc)

	// Empty blacklist, valid whitelist some whitelist is not included in repo
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{"a/b/c.yaml", "c.json", "someRandomFile.yaml"}, nil),
	)
	rawList = []string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/a/b/c.yaml", "/repo/c.json"}, []string{}, nil}
	createAndAssert(t, tc)

	// Both whitelist and blacklist contain the same file
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{"a/b/c.yaml"}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{"a/b/c.yaml", "c.json"}, nil),
	)
	rawList = []string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/c.json"}, []string{"/repo/a/b/c.yaml"}, nil}
	createAndAssert(t, tc)

	// Both whitelist and blacklist contain the same file and other comments.
	gomock.InOrder(
		fs.EXPECT().ReadLines("/blacklist").Times(1).Return([]string{"a/b/c.yaml", "#   c.json"}, nil),
		fs.EXPECT().ReadLines("/whitelist").Times(1).Return([]string{"a/b/c.yaml", "c.json", "#   a/b/c.yaml"}, nil),
	)
	rawList = []string{"/repo/a/b.json", "/repo/b/c", "/repo/a/b/c.yaml", "/repo/a/b/c", "/repo/c.json"}
	tc = testCase{rawList, "/repo", "/blacklist", "/whitelist", fs, []string{"/repo/c.json"}, []string{"/repo/a/b/c.yaml"}, nil}
	createAndAssert(t, tc)
}

func createAndAssert(t *testing.T, tc testCase) {
	assert := assert.New(t)
	f := &Factory{tc.repoPath, tc.blacklistPath, tc.whitelistPath, tc.fs}
	applyList, blacklist, _, err := f.Create(tc.rawList)
	assert.Equal(tc.expectedApplyList, applyList)
	assert.Equal(tc.expectedBlacklist, blacklist)
	assert.Equal(tc.expectedErr, err)
}
