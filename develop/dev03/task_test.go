package main

import (
	"dev03/mansort"
	"os"
	"path/filepath"
	"testing"
)

func TestSort(t *testing.T) {
	dat, err := os.ReadFile(filepath.Join(".", "assets", "test.txt"))
	if err != nil {
		t.Fatal(err)
	}

	ts := TestSorts{
		data: dat,
	}

	t.Log("Run sorts test")
	t.Run("DEFAULT SORT WITHOUT KEYS", ts.TestDefaultSortWithoutKeys)
	t.Run("DEFAULT SORT FLAG K=0", ts.TestDefaultSortFlagK0)
	t.Run("DEFAULT SORT FLAG K=6", ts.TestDefaultSortFlagK6)
	t.Run("DEFAULT SORT FLAG K=6 U=TRUE", ts.TestDefaultSortFlagK6U)
	t.Run("DEFAULT SORT FLAG K=6 U=TRUE R=TRUE", ts.TestDefaultSortFlagK6UR)
	t.Run("NUM SORT WITHOUT FLAGS", ts.TestNumSortWithoutKeys)
	t.Run("NUM SORT FLAG K=6", ts.TestNumSortFlagK6)
	t.Run("NUM SORT FLAG K=6 U=TRUE", ts.TestNumSortFlagK6U)
	t.Run("NUM SORT FLAG K=6 U=TRUE R=TRUE", ts.TestNumSortFlagK6UR)
}

type TestSorts struct {
	data []byte
}

func (ts *TestSorts) TestDefaultSortWithoutKeys(t *testing.T) {
	want := `1 Oct 9 12:00:00 cindy@example.com deferred 5
1 Oct 9 12:00:00 cindy@example.com deferred 5
10 Oct 10 13:30:00 billy@example.com deferred 8
10 Oct 10 14:30:00 billy@example.com bounced 10
2 Oct 10 12:00:00 andy@example.com deferred 2
2 Oct 10 13:00:00 andy@example.com deferred 6
2 Oct 9 13:00:00 cindy@example.com deferred 4
3 Oct 9 14:00:00 cindy@example.com bounced 3
5 Oct 10 12:30:00 billy@example.com deferred 1
9 Oct 10 14:00:00 andy@example.com bounced 9
A
a
b
b
c`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestDefaultSortFlagK0(t *testing.T) {
	want := `1 Oct 9 12:00:00 cindy@example.com deferred 5
1 Oct 9 12:00:00 cindy@example.com deferred 5
10 Oct 10 13:30:00 billy@example.com deferred 8
10 Oct 10 14:30:00 billy@example.com bounced 10
2 Oct 10 12:00:00 andy@example.com deferred 2
2 Oct 10 13:00:00 andy@example.com deferred 6
2 Oct 9 13:00:00 cindy@example.com deferred 4
3 Oct 9 14:00:00 cindy@example.com bounced 3
5 Oct 10 12:30:00 billy@example.com deferred 1
9 Oct 10 14:00:00 andy@example.com bounced 9
A
a
b
b
c`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		K: 0,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestDefaultSortFlagK6(t *testing.T) {
	want := `A
c
b
a
b
5 Oct 10 12:30:00 billy@example.com deferred 1
10 Oct 10 14:30:00 billy@example.com bounced 10
2 Oct 10 12:00:00 andy@example.com deferred 2
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 9 13:00:00 cindy@example.com deferred 4
1 Oct 9 12:00:00 cindy@example.com deferred 5
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 10 13:00:00 andy@example.com deferred 6
10 Oct 10 13:30:00 billy@example.com deferred 8
9 Oct 10 14:00:00 andy@example.com bounced 9`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		K: 6,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestDefaultSortFlagK6U(t *testing.T) {
	want := `A
c
b
a
5 Oct 10 12:30:00 billy@example.com deferred 1
10 Oct 10 14:30:00 billy@example.com bounced 10
2 Oct 10 12:00:00 andy@example.com deferred 2
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 9 13:00:00 cindy@example.com deferred 4
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 10 13:00:00 andy@example.com deferred 6
10 Oct 10 13:30:00 billy@example.com deferred 8
9 Oct 10 14:00:00 andy@example.com bounced 9`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		K: 6,
		U: true,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestDefaultSortFlagK6UR(t *testing.T) {
	want := `9 Oct 10 14:00:00 andy@example.com bounced 9
10 Oct 10 13:30:00 billy@example.com deferred 8
2 Oct 10 13:00:00 andy@example.com deferred 6
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 9 13:00:00 cindy@example.com deferred 4
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 10 12:00:00 andy@example.com deferred 2
10 Oct 10 14:30:00 billy@example.com bounced 10
5 Oct 10 12:30:00 billy@example.com deferred 1
a
b
c
A`
	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		K: 6,
		U: true,
		R: true,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestNumSortWithoutKeys(t *testing.T) {
	want := `A
c
b
a
b
10 Oct 10 13:30:00 billy@example.com deferred 8
10 Oct 10 14:30:00 billy@example.com bounced 10
9 Oct 10 14:00:00 andy@example.com bounced 9
5 Oct 10 12:30:00 billy@example.com deferred 1
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 9 13:00:00 cindy@example.com deferred 4
2 Oct 10 12:00:00 andy@example.com deferred 2
2 Oct 10 13:00:00 andy@example.com deferred 6
1 Oct 9 12:00:00 cindy@example.com deferred 5
1 Oct 9 12:00:00 cindy@example.com deferred 5`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		N: true,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestNumSortFlagK6(t *testing.T) {
	want := `A
c
b
a
b
10 Oct 10 14:30:00 billy@example.com bounced 10
9 Oct 10 14:00:00 andy@example.com bounced 9
10 Oct 10 13:30:00 billy@example.com deferred 8
2 Oct 10 13:00:00 andy@example.com deferred 6
1 Oct 9 12:00:00 cindy@example.com deferred 5
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 9 13:00:00 cindy@example.com deferred 4
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 10 12:00:00 andy@example.com deferred 2
5 Oct 10 12:30:00 billy@example.com deferred 1`
	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		N: true,
		K: 6,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestNumSortFlagK6U(t *testing.T) {
	want := `A
c
b
a
10 Oct 10 14:30:00 billy@example.com bounced 10
9 Oct 10 14:00:00 andy@example.com bounced 9
10 Oct 10 13:30:00 billy@example.com deferred 8
2 Oct 10 13:00:00 andy@example.com deferred 6
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 9 13:00:00 cindy@example.com deferred 4
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 10 12:00:00 andy@example.com deferred 2
5 Oct 10 12:30:00 billy@example.com deferred 1`
	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		N: true,
		K: 6,
		U: true,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}

func (ts *TestSorts) TestNumSortFlagK6UR(t *testing.T) {
	want := `5 Oct 10 12:30:00 billy@example.com deferred 1
2 Oct 10 12:00:00 andy@example.com deferred 2
3 Oct 9 14:00:00 cindy@example.com bounced 3
2 Oct 9 13:00:00 cindy@example.com deferred 4
1 Oct 9 12:00:00 cindy@example.com deferred 5
2 Oct 10 13:00:00 andy@example.com deferred 6
10 Oct 10 13:30:00 billy@example.com deferred 8
9 Oct 10 14:00:00 andy@example.com bounced 9
10 Oct 10 14:30:00 billy@example.com bounced 10
a
b
c
A`

	sort := mansort.NewSort(ts.data, mansort.SortFlags{
		N: true,
		K: 6,
		U: true,
		R: true,
	})
	sort.Sort()
	got := sort.Lines()
	if want != got {
		t.Fatalf("\nwant: \n%s\n\ngot: \n%s", want, got)
	}
}
