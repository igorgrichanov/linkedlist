package task

import (
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func GenerateSliceOfCommits(n int) [][]Commit {
	result := make([][]Commit, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, GenerateCommit(1000))
	}
	return result
}

func GenerateSliceOfLists(n int) []*DoubleLinkedList {
	result := make([]*DoubleLinkedList, 0, n)
	for i := 0; i < n; i++ {
		listToAppend := GenerateData(1000)
		r := rand.Intn(1000)
		for j := 0; j < r; j++ {
			listToAppend.Next()
		}
		result = append(result, listToAppend)
	}
	return result
}

func TestQuickSort(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now()
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time1,
	}
	commit2 := &Commit{
		Message: "We need to program the bluetooth ADP protocol!",
		UUID:    "6956d3f4-875b-11ed-8150-acde48001122",
		Date:    time2,
	}
	node1 := &Node{
		data: commit1,
	}
	node2 := &Node{
		data: commit2,
		prev: node1,
		next: nil,
	}
	node1.next = node2

	type args struct {
		data []Commit
	}
	tests := []struct {
		name string
		args args
		want *DoubleLinkedList
	}{
		{
			name: "testing QuickSort algorithm",
			args: args{
				data: []Commit{*commit1, *commit2},
			},
			want: &DoubleLinkedList{
				head: node1,
				tail: node2,
				curr: node1,
				len:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.data); got.head.data.Message != tt.want.head.data.Message {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Append(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	type args struct {
		commit *Commit
	}
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		args   args
	}{
		{
			name:   "appending to an empty list",
			fields: &DoubleLinkedList{},
			args: args{
				commit: commit1,
			},
		},
		{
			name:   "appending to a non-empty list",
			fields: testLinkedList,
			args: args{
				commit: commit1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			d.Append(tt.args.commit)
			t.Logf("len: %d\n", d.len)
			if d.tail.data != commit1 {
				t.Errorf("failed to append a commit to the list")
			}
		})
	}
}

func TestDoubleLinkedList_LoadData(t *testing.T) {
	type fields struct {
		head *Node
		tail *Node
		curr *Node
		len  int
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "load data from test.json",
			fields:  fields{},
			args:    args{path: "./test.json"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.LoadData(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("LoadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Len(t *testing.T) {
	testLinkedList := GenerateData(10)
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   int
	}{
		{
			name:   "testing len method with an empty linked list",
			fields: &DoubleLinkedList{},
			want:   0,
		},
		{
			name:   "testing len method with a non-empty linked list",
			fields: testLinkedList,
			want:   10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Len(); got != tt.want {
				t.Errorf("Len() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Current(t *testing.T) {
	testLinkedList := GenerateData(10)
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *Node
	}{
		{
			name:   "testing current method with an empty linked list",
			fields: &DoubleLinkedList{},
			want:   nil,
		},
		{
			name:   "testing current method with a non-empty linked list",
			fields: testLinkedList,
			want:   testLinkedList.head.next.next,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if i == 1 {
				d.Next()
				d.Next()
				if got := d.Current(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Current2() = %v, want %v", got, tt.want)
				}
			} else {
				if got := d.Current(); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Current() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}

func TestDoubleLinkedList_Next(t *testing.T) {
	testLinkedList := GenerateData(10)
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *Node
	}{
		{
			name:   "testing next method with a non-empty list",
			fields: testLinkedList,
			want:   testLinkedList.head.next,
		},
		{
			name:   "testing next method with an empty list",
			fields: &DoubleLinkedList{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Next(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Next() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Prev(t *testing.T) {
	testLinkedList := GenerateData(10)
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *Node
	}{
		{
			name:   "testing prev method with a non-empty list",
			fields: testLinkedList,
			want:   nil,
		},
		{
			name:   "testing prev method with an empty list",
			fields: &DoubleLinkedList{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Prev(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Prev() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Insert(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	type args struct {
		n int
		c *Commit
	}
	tests := []struct {
		name    string
		fields  *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{
			name:   "inserting to the end of the list",
			fields: testLinkedList,
			args: args{
				n: 9,
				c: commit,
			},
			wantErr: false,
		},
		{
			name:   "inserting to the beginning of the list",
			fields: testLinkedList,
			args: args{
				n: 0,
				c: &Commit{
					Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
					UUID:    "6956db06-875b-11ed-8150-acde48001122",
					Date:    time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:   "inserting to an empty list and creating a new list if n=0",
			fields: &DoubleLinkedList{},
			args: args{
				n: 0,
				c: &Commit{
					Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
					UUID:    "6956db06-875b-11ed-8150-acde48001122",
					Date:    time.Now(),
				},
			},
			wantErr: false,
		},
		{
			name:   "inserting to the list when n > length",
			fields: testLinkedList,
			args: args{
				n: 10,
				c: &Commit{
					Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
					UUID:    "6956db06-875b-11ed-8150-acde48001122",
					Date:    time.Now(),
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Insert(tt.args.n, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("len: %d\n", d.len)
		})
	}
}

func TestDoubleLinkedList_Delete(t *testing.T) {
	testLinkedList := GenerateData(10)
	type args struct {
		n int
	}
	tests := []struct {
		name    string
		fields  *DoubleLinkedList
		args    args
		wantErr bool
	}{
		{
			name:   "remove the first element from a non-empty list",
			fields: testLinkedList,
			args: args{
				n: 0,
			},
			wantErr: false,
		},
		{
			name:   "remove the last element from a non-empty list",
			fields: testLinkedList,
			args: args{
				n: 9,
			},
			wantErr: false,
		},
		{
			name:   "remove an element from an empty list",
			fields: &DoubleLinkedList{},
			args: args{
				n: 0,
			},
			wantErr: true,
		},
		{
			name:   "remove an element with and index that is more than the length of the list",
			fields: testLinkedList,
			args: args{
				n: 10,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.Delete(tt.args.n); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Logf("length: %d\n", d.len)
		})
	}
}

func TestDoubleLinkedList_DeleteCurrent(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	node1 := &Node{
		data: commit1,
	}
	tests := []struct {
		name    string
		fields  *DoubleLinkedList
		wantErr bool
	}{
		{
			name:    "remove current element from a non-empty list",
			fields:  testLinkedList,
			wantErr: false,
		},
		{
			name:    "remove current element from an empty list",
			fields:  &DoubleLinkedList{},
			wantErr: true,
		},
		{
			name: "remove an element from a non-empty list with len=1",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if err := d.DeleteCurrent(); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCurrent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDoubleLinkedList_Index(t *testing.T) {
	testLinkedList := GenerateData(10)
	tests := []struct {
		name    string
		fields  *DoubleLinkedList
		want    int
		wantErr bool
	}{
		{
			name:    "return an index of a current element of an empty list",
			fields:  &DoubleLinkedList{},
			want:    0,
			wantErr: true,
		},
		{
			name:    "return an index of a current element of a non-empty list",
			fields:  testLinkedList,
			want:    0,
			wantErr: false,
		},
		{
			name:    "return an index of a current element if next() method has been invoked",
			fields:  testLinkedList,
			want:    1,
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if i == 2 {
				d.Next()
			}
			got, err := d.Index()
			if (err != nil) != tt.wantErr {
				t.Errorf("Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Index() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Pop(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	node1 := &Node{
		data: commit1,
	}
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *Node
	}{
		{
			name:   "pop an element from an empty list",
			fields: &DoubleLinkedList{},
			want:   &Node{},
		},
		{
			name: "pop an element from a list with length=1",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			want: node1,
		},
		{
			name:   "pop an element from a non-empty list",
			fields: testLinkedList,
			want:   testLinkedList.tail,
		},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Pop(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pop() = %v, want %v", got, tt.want)
			}
			if i == 1 && d.len != 0 {
				t.Errorf("lenght of a list must be 0, have: %d\n", d.len)
			}
			if i == 2 && d.len != 9 {
				t.Errorf("lenght of a list must be 9, have: %d\n", d.len)
			}
		})
	}
}

func TestDoubleLinkedList_Shift(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	node1 := &Node{
		data: commit1,
	}
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *Node
	}{
		{
			name:   "shift an empty list",
			fields: &DoubleLinkedList{},
			want:   &Node{},
		},
		{
			name: "shift a list with length=1",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			want: node1,
		},
		{
			name:   "shift a non-empty list",
			fields: testLinkedList,
			want:   testLinkedList.tail,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Shift(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Shift() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_SearchUUID(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	node1 := &Node{
		data: commit1,
	}
	type args struct {
		uuID string
	}
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		args   args
		want   *Node
	}{
		{
			name:   "search an element by UUID in an empty list",
			fields: &DoubleLinkedList{},
			args: args{
				uuID: "does not matter",
			},
			want: &Node{},
		},
		{
			name: "search an element by UUID",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			args: args{
				uuID: "6956db06-875b-11ed-8150-acde48001122",
			},
			want: node1,
		},
		{
			name:   "search an element by UUID if UUID does not exist",
			fields: testLinkedList,
			args: args{
				uuID: "I love you, Golang!",
			},
			want: &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.SearchUUID(tt.args.uuID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchUUID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Search(t *testing.T) {
	testLinkedList := GenerateData(10)
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time.Now(),
	}
	node1 := &Node{
		data: commit1,
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields *DoubleLinkedList
		args   args
		want   *Node
	}{
		{
			name:   "search an element by message in an empty list",
			fields: &DoubleLinkedList{},
			args: args{
				message: "does not matter",
			},
			want: &Node{},
		},
		{
			name: "search an element by message",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			args: args{
				message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
			},
			want: node1,
		},
		{
			name:   "search an element by message if message does not exist",
			fields: testLinkedList,
			args: args{
				message: "I love you, Golang!",
			},
			want: &Node{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Search(tt.args.message); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Reverse(t *testing.T) {
	time1 := time.Now()
	time2 := time.Now()
	commit1 := &Commit{
		Message: "Try to calculate the IB bandwidth, maybe it will generate the bluetooth program!",
		UUID:    "6956db06-875b-11ed-8150-acde48001122",
		Date:    time1,
	}
	commit2 := &Commit{
		Message: "We need to program the bluetooth ADP protocol!",
		UUID:    "6956d3f4-875b-11ed-8150-acde48001122",
		Date:    time2,
	}
	node1 := &Node{
		data: commit1,
	}
	node2 := &Node{
		data: commit2,
		prev: node1,
		next: nil,
	}
	node1.next = node2

	node1InReversedList := &Node{
		data: commit2,
	}
	node2InReversedList := &Node{
		data: commit1,
		prev: node1InReversedList,
	}
	node1InReversedList.next = node2InReversedList

	tests := []struct {
		name   string
		fields *DoubleLinkedList
		want   *DoubleLinkedList
	}{
		{
			name:   "trying to receive reversed list from an empty list",
			fields: &DoubleLinkedList{},
			want:   &DoubleLinkedList{},
		},
		{
			name: "trying to receive reversed list from a list with length=1",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
			want: &DoubleLinkedList{
				head: node1,
				tail: node1,
				curr: node1,
				len:  1,
			},
		},
		{
			name: "trying to receive reversed list from a list with length=10",
			fields: &DoubleLinkedList{
				head: node1,
				tail: node2,
				curr: node1,
				len:  2,
			},
			want: &DoubleLinkedList{
				head: node1InReversedList,
				tail: node2InReversedList,
				curr: node1InReversedList,
				len:  2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
				curr: tt.fields.curr,
				len:  tt.fields.len,
			}
			if got := d.Reverse(); !reflect.DeepEqual(got, tt.want) {
				if d.len == 1 && got.head.data == tt.want.head.data {
					return
				} else if got.head.data == tt.want.tail.data && got.len == 2 {
					return
				}
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuickSort(b *testing.B) {
	dataSet := GenerateSliceOfCommits(b.N)
	for i := 0; i < b.N; i++ {
		QuickSort(dataSet[i])
	}
}

func BenchmarkDoubleLinkedList_Append(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	var commitToAppend Commit
	err := gofakeit.Struct(&commitToAppend)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		dataSet[i].Append(&commitToAppend)
	}
}

func BenchmarkDoubleLinkedList_Current(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Current()
	}
}

func BenchmarkDoubleLinkedList_Delete(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		n := rand.Intn(1000)
		err := dataSet[i].Delete(n)
		if err != nil {
			log.Printf("Encountered an error: %v", err)
		}
	}
}

func BenchmarkDoubleLinkedList_DeleteCurrent(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		err := dataSet[i].DeleteCurrent()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkDoubleLinkedList_Index(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		_, err := dataSet[i].Index()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkDoubleLinkedList_Insert(b *testing.B) {
	var commit Commit
	err := gofakeit.Struct(&commit)
	if err != nil {
		log.Fatal(err)
	}
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		set := dataSet[i]
		n := rand.Intn(1000)
		err = set.Insert(n, &commit)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkDoubleLinkedList_Len(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Len()
	}
}

func BenchmarkDoubleLinkedList_LoadData(b *testing.B) {
	dataSet := &DoubleLinkedList{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := dataSet.LoadData("./test.json")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkDoubleLinkedList_Next(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Next()
	}
}

func BenchmarkDoubleLinkedList_Prev(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Prev()
	}
}

func BenchmarkDoubleLinkedList_Pop(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Pop()
	}
}

func BenchmarkDoubleLinkedList_Reverse(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Reverse()
	}
}

func BenchmarkDoubleLinkedList_Search(b *testing.B) {
	err := list.LoadData("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	messageToSearch := "You can't calculate the microchip without indexing the auxiliary HDD capacitor!"
	for i := 0; i < b.N; i++ {
		list.Search(messageToSearch)
	}
}

func BenchmarkDoubleLinkedList_SearchUUID(b *testing.B) {
	err := list.LoadData("./test.json")
	if err != nil {
		log.Fatal(err)
	}
	uuidToSearch := "695860de-875b-11ed-8150-acde48001122"
	for i := 0; i < b.N; i++ {
		list.SearchUUID(uuidToSearch)
	}
}

func BenchmarkDoubleLinkedList_Shift(b *testing.B) {
	dataSet := GenerateSliceOfLists(b.N)
	for i := 0; i < b.N; i++ {
		dataSet[i].Shift()
	}
}
