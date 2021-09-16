package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// mình tìm được một số ví dụ trên mạng là sẽ in trực tiếp cây thư mục duyệt đến thư mục đó
// mình làm theo cách khác là sẽ build ra 1 cây bằng linkedList sau đó query và in ra
// trong đó mỗi thực thể là 1 node
// nếu node đó là folder thì sẽ có children là 1 linkedList
// tham khảo cách dựng linkedList ở đây: https://github.com/DavidMoranchel/go-data-structures/blob/master/linked_list.go

type node struct {
	name     string
	size     int64
	next     *node
	children *linkedList
}

type linkedList struct {
	head *node
}

func (l linkedList) Display(level int) {
	for l.head != nil {
		fmt.Printf("%s| %s\n", strings.Repeat("   ", level), l.head.name)
		if l.head.children != nil {
			l.head.children.Display(level + 1)
		}
		l.head = l.head.next
	}
}

func (l linkedList) Count() (countDir, countFile int) {
	for l.head != nil {
		if l.head.children != nil {
			_countDir, _countFile := l.head.children.Count()
			countDir += _countDir + 1
			countFile += _countFile
		} else {
			countFile += 1
		}
		l.head = l.head.next
	}
	return
}

func (l *linkedList) Append(n *node) {
	if l.head == nil {
		l.head = n
	} else {
		cNode := l.head
		for cNode.next != nil {
			cNode = cNode.next
		}
		cNode.next = n
	}
}

func getDir(root *linkedList, path string) {
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		// mình sẽ check có symlink ko, vì nếu có sẽ có thể lặp vô tận :(
		if (entry.Mode() & os.ModeSymlink) == os.ModeSymlink {
			n := &node{}
			fullPath, err := os.Readlink(filepath.Join(path, entry.Name()))
			if err != nil {
				fullPath = ""
			}
			n.name = fmt.Sprintf("%s -> %s", entry.Name(), fullPath)
			n.size = entry.Size()
			root.Append(n)
		} else if entry.IsDir() {
			n := &node{}
			cList := &linkedList{}
			n.name = entry.Name()
			n.children = cList
			fullPath := filepath.Join(path, entry.Name())
			if err != nil {
				log.Fatal()
			}
			root.Append(n)
			getDir(cList, fullPath)
		} else {
			n := &node{}
			n.name = entry.Name()
			n.size = entry.Size()
			root.Append(n)
		}
	}
}

func main() {
	path := "/Users/toanvo/Documents/git/learning-go/techmaster"
	root := &linkedList{}
	getDir(root, path)
	root.Display(0)
	dirs, files := root.Count()
	fmt.Printf("%d directories, %d files", dirs, files)
}
